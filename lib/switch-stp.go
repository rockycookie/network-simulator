package lib

import (
	"fmt"
	"sync"
	"time"
)

type SwitchStp struct {
	ID                  string // to make life easier, use the first NIC MAC address of the switch
	HelloTimeSeconds    uint16
	ForwardDelaySeconds uint16
	State               uint8
	RootSwitchId        string
	RootPathCost        uint32
	mu                  sync.Mutex
}

const (
	SW_STP_STATE_DISABLED   = 0
	SW_STP_STATE_BLOCKING   = 1
	SW_STP_STATE_LISTENING  = 2
	SW_STP_STATE_LEARNING   = 3
	SW_STP_STATE_FORWARDING = 4
	SW_STP_STATE_BROKEN     = 5
)

func (sw *Switch) initStp() {
	sw.StpInfo = SwitchStp{
		ID:                  "",
		HelloTimeSeconds:    2,
		ForwardDelaySeconds: 15,
		State:               SW_STP_STATE_DISABLED,
		RootSwitchId:        "",
		RootPathCost:        0,
	}

	// Set the switch ID to the lowest MAC address among its NICs
	for i := range sw.Nics {
		if sw.StpInfo.ID == "" || sw.Nics[i].Mac < sw.StpInfo.ID {
			sw.StpInfo.ID = sw.Nics[i].Mac
		}
	}

	// Initialize each port's STP info
	for i := range sw.Nics {
		sw.Nics[i].StpInfo = PortStp{
			Type:         PORT_STP_TYPE_BLOCKING,
			LinkCost:     1,             // default link cost, simplified version
			RootSwitchId: sw.StpInfo.ID, // initially assume we are the root
			RootPathCost: 0,             // no cost to reach ourselves
		}
	}
}

func (sw *Switch) RunStp() {

	// Do not block other work in switch Run()
	go func() {
		// Set all ports to listening state initially, and start sending config BPDUs
		for i := range sw.Nics {
			sw.StpInfo.State = SW_STP_STATE_LISTENING
			//TODO: no-cable connected ports should be in disabled state and not send BPDUs, but for simplicity
			// we will run the goroutines and not sending BPDU within the goroutine if no cable is connected
			// If we want to optimize, we need to support switch handling hot plug/unplug of cables and
			// dynamically start/stop the BPDU goroutines based on cable connection state
			go func(nic *Nic) {
				nic.sendConfigBpdu(sw.StpInfo.HelloTimeSeconds)
			}(&sw.Nics[i])
		}

		// Wait for ForwardDelaySeconds before transitioning ports to forwarding state
		time.Sleep(time.Duration(sw.StpInfo.ForwardDelaySeconds) * time.Second)

		// Finished root election, let's treat listening and learning as the same state for now
		sw.updatePortsStateAfterRootElection()
		sw.StpInfo.State = SW_STP_STATE_FORWARDING
	}()
}

func (nic *Nic) sendConfigBpdu(helloTimeSeconds uint16) {

	for nic.Switch.StpInfo.State == SW_STP_STATE_LISTENING {
		// Sync Switch local ports STP information first
		nic.Switch.syncPortsSTP()

		bpdu := L2Frame{
			ConfigBpdu: &ConfigBpdu{
				RootSwitchId:  nic.Switch.StpInfo.RootSwitchId,
				RootPathCost:  nic.Switch.StpInfo.RootPathCost,
				LocalSwitchId: nic.Switch.StpInfo.ID,
			},
			SrcMac: nic.Mac,
			DstMac: STP_BPDU_DEST_MAC,
		}

		// Send BPDU out of each non-blocking port
		if nic.ConnectedCable != nil {
			nic.ConnectedCable.TransmitFrame(nic, bpdu)
		}

		// Sleep for the hello interval
		time.Sleep(time.Duration(helloTimeSeconds) * time.Second)
	}
}

func (sw *Switch) ProcessConfigBpdu(inboundBPDU *ConfigBpdu, inboundNic *Nic) {

	if inboundNic.Switch.StpInfo.State == SW_STP_STATE_LISTENING {

		// Update root switch info if we receive a better root
		if inboundBPDU.RootSwitchId < inboundNic.StpInfo.RootSwitchId {
			inboundCost := inboundBPDU.RootPathCost + inboundNic.StpInfo.LinkCost
			inboundRootId := inboundBPDU.RootSwitchId

			if EnableStpLogging {
				fmt.Printf("[%s][%s:%s] Received BPDU with lower-Root-ID: RootSwitchId=%s, RootPathCost=%d (was RootBridgeId=%s, RootPathCost=%d)\n",
					time.Now().UTC().Format(time.RFC3339Nano), sw.Name, inboundNic.ID,
					inboundRootId, inboundCost, inboundNic.StpInfo.RootSwitchId, inboundNic.StpInfo.RootPathCost)
			}

			// write the better root info to the inbound port's STP info
			inboundNic.swLock()
			inboundNic.StpInfo.RootSwitchId = inboundRootId
			inboundNic.StpInfo.RootPathCost = inboundCost
			inboundNic.swUnlock()
		} else {

			// otherwise just do logging
			if EnableStpLogging {
				fmt.Printf("[%s][%s:%s] Received BPDU: RootSwitchId=%s, RootPathCost=%d, LocalSwitchId=%s\n",
					time.Now().UTC().Format(time.RFC3339Nano), sw.Name, inboundNic.ID,
					inboundBPDU.RootSwitchId, inboundBPDU.RootPathCost, inboundBPDU.LocalSwitchId)
			}
		}

		// Always update the other switch info
		inboundNic.swLock()
		inboundNic.StpInfo.OtherSwitchId = inboundBPDU.LocalSwitchId
		inboundNic.StpInfo.OtherSwitchRootPathCost = inboundBPDU.RootPathCost
		inboundNic.swUnlock()
	}
}

func (sw *Switch) calculateRootSwitchIdAndCostAndPortIndex() (string, uint32, int) {
	var rootId string = ""
	var cost uint32 = 0
	var portIndex int = -1

	sw.StpInfo.mu.Lock()
	for i := range sw.Nics {
		if rootId == "" || (sw.Nics[i].StpInfo.RootSwitchId != "" && sw.Nics[i].StpInfo.RootSwitchId <= rootId) {

			// If there are multiple paths to the same root, choose the one with the lowest cost
			if sw.Nics[i].StpInfo.RootSwitchId == rootId && sw.Nics[i].StpInfo.RootPathCost > cost {
				continue
			}

			// update rootId, cost, and portIndex if this port has a better root path
			rootId = sw.Nics[i].StpInfo.RootSwitchId
			cost = sw.Nics[i].StpInfo.RootPathCost
			portIndex = i
		}

	}
	sw.StpInfo.mu.Unlock()

	return rootId, cost, portIndex
}

func (sw *Switch) syncPortsSTP() {
	rootId, cost, _ := sw.calculateRootSwitchIdAndCostAndPortIndex()

	sw.StpInfo.mu.Lock()
	sw.StpInfo.RootSwitchId = rootId
	sw.StpInfo.RootPathCost = cost
	sw.StpInfo.mu.Unlock()
}

func (sw *Switch) updatePortsStateAfterRootElection() {

	if sw.StpInfo.RootSwitchId == sw.StpInfo.ID {
		// All ports are designated ports for STP root
		for i := range sw.Nics {
			sw.Nics[i].StpInfo.Type = PORT_STP_TYPE_DESIGNATED
		}

		if EnableStpLogging {
			fmt.Printf("[%s][%s] ------> I am the root switch, all ports set to DESIGNATED.\n", time.Now().UTC().Format(time.RFC3339Nano), sw.Name)
		}
	} else {
		_, localRootPathCost, rootPortIndex := sw.calculateRootSwitchIdAndCostAndPortIndex()

		for i := range sw.Nics {
			if sw.Nics[i].StpInfo.RootSwitchId == "" {
				continue
			} else if i == rootPortIndex {
				sw.Nics[i].StpInfo.Type = PORT_STP_TYPE_ROOT

				if EnableStpLogging {
					fmt.Printf("[%s][%s:%s] ------> Port set to ROOT (RootSwitchId=%s, RootPathCost=%d)\n",
						time.Now().UTC().Format(time.RFC3339Nano), sw.Name, sw.Nics[i].ID,
						sw.Nics[i].StpInfo.RootSwitchId, sw.Nics[i].StpInfo.RootPathCost)
				}
			} else {
				// decide if designated
				if localRootPathCost < sw.Nics[i].StpInfo.OtherSwitchRootPathCost {
					sw.Nics[i].StpInfo.Type = PORT_STP_TYPE_DESIGNATED
				} else if localRootPathCost == sw.Nics[i].StpInfo.OtherSwitchRootPathCost && sw.StpInfo.ID < sw.Nics[i].StpInfo.OtherSwitchId {
					sw.Nics[i].StpInfo.Type = PORT_STP_TYPE_DESIGNATED
				} else {
					sw.Nics[i].StpInfo.Type = PORT_STP_TYPE_BLOCKING
				}

				if EnableStpLogging {
					fmt.Printf("[%s][%s:%s] ------> Port set to %s (LocalRootPathCost=%d, OtherSwitchRootPathCost=%d, LocalSwitchId=%s, OtherSwitchId=%s)\n",
						time.Now().UTC().Format(time.RFC3339Nano), sw.Name, sw.Nics[i].ID, sw.Nics[i].StpInfo.getPortRole(),
						localRootPathCost, sw.Nics[i].StpInfo.OtherSwitchRootPathCost,
						sw.StpInfo.ID, sw.Nics[i].StpInfo.OtherSwitchId)
				}
			}
		}
	}
}
