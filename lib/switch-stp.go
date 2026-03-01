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
	RootBridgeId        string
	RootPathCost        uint32
	mu                  sync.Mutex
}

type PortStp struct {
	Type         uint8
	LinkCost     uint32
	RootBridgeId string
	RootPathCost uint32
	mu           sync.Mutex
}

const (
	PORT_STP_TYPE_ROOT       = 0
	PORT_STP_TYPE_DESIGNATED = 1
	PORT_STP_TYPE_BLOCKING   = 2

	SW_STP_STATE_DISABLED   = 0
	SW_STP_STATE_BLOCKING   = 1
	SW_STP_STATE_LISTENING  = 2
	SW_STP_STATE_LEARNING   = 3
	SW_STP_STATE_FORWARDING = 4
	SW_STP_STATE_BROKEN     = 5

	// The standard destination MAC address for Bridge Protocol Data Units (BPDUs) in STP/RSTP is the IEEE 802.1D multicast address 01:80:C2:00:00:00
	STP_BPDU_DEST_MAC = "01:80:C2:00:00:00"
)

func (sw *Switch) initStp() {
	sw.StpInfo = SwitchStp{
		ID:                  "",
		HelloTimeSeconds:    2,
		ForwardDelaySeconds: 15,
		State:               SW_STP_STATE_DISABLED,
		RootBridgeId:        "",
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
			RootBridgeId: sw.StpInfo.ID, // initially assume we are the root
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

		// Finished root election
		//TODO: continue more later
		sw.StpInfo.State = SW_STP_STATE_BROKEN // temporary state to indicate STP is done
	}()
}

func (nic *Nic) sendConfigBpdu(helloTimeSeconds uint16) {

	for nic.Switch.StpInfo.State == SW_STP_STATE_LISTENING {
		// Sync Switch local ports STP information first
		nic.Switch.syncPortsSTP()

		bpdu := L2Frame{
			ConfigBpdu: &ConfigBpdu{
				RootBridgeId: nic.Switch.StpInfo.RootBridgeId,
				RootPathCost: nic.Switch.StpInfo.RootPathCost,
			},
			SrcMac: nic.Mac,
			DstMac: STP_BPDU_DEST_MAC,
		}

		// Send BPDU out of each non-blocking port
		if nic.ConnectedCable != nil {
			nic.ConnectedCable.TransmitFrame(nic, bpdu)
		}

		if EnableStpLogging {
			fmt.Printf("[%s][Switch %s] Sending BPDU on NIC %s: RootBridgeId=%s, RootPathCost=%d\n",
				time.Now().UTC().Format(time.RFC3339Nano), nic.Switch.Name, nic.ID,
				nic.Switch.StpInfo.RootBridgeId, nic.Switch.StpInfo.RootPathCost)
		}

		// Sleep for the hello interval
		time.Sleep(time.Duration(helloTimeSeconds) * time.Second)
	}
}

func (sw *Switch) ProcessConfigBpdu(inboundBPDU *ConfigBpdu, inboundNic *Nic) {

	if inboundNic.Switch.StpInfo.State == SW_STP_STATE_LISTENING {

		// Update root bridge info if we receive a better root
		if inboundBPDU.RootBridgeId < inboundNic.StpInfo.RootBridgeId {
			inboundCost := inboundBPDU.RootPathCost + inboundNic.StpInfo.LinkCost
			inboundRootId := inboundBPDU.RootBridgeId

			if EnableStpLogging {
				fmt.Printf("[%s][Switch %s] Received better BPDU on NIC %s: RootBridgeId=%s, RootPathCost=%d (was RootBridgeId=%s, RootPathCost=%d)\n",
					time.Now().UTC().Format(time.RFC3339Nano), sw.Name, inboundNic.ID,
					inboundRootId, inboundCost, inboundNic.StpInfo.RootBridgeId, inboundNic.StpInfo.RootPathCost)
			}

			// write the better root info to the inbound port's STP info
			inboundNic.swLock()
			inboundNic.StpInfo.RootBridgeId = inboundRootId
			inboundNic.StpInfo.RootPathCost = inboundCost
			inboundNic.swUnlock()
		}
	}
}

func (sw *Switch) calculateRootBridgeIdAndCostAndPortIndex() (string, uint32, int) {
	var rootId string = ""
	var cost uint32 = 0
	var portIndex int = -1

	sw.StpInfo.mu.Lock()
	for i := range sw.Nics {
		if rootId == "" || (sw.Nics[i].StpInfo.RootBridgeId != "" && sw.Nics[i].StpInfo.RootBridgeId <= rootId) {

			// If there are multiple paths to the same root, choose the one with the lowest cost
			if sw.Nics[i].StpInfo.RootBridgeId == rootId {
				if sw.Nics[i].StpInfo.RootPathCost < cost {
					cost = sw.Nics[i].StpInfo.RootPathCost
					portIndex = i
					continue
				}
			}

			// update rootId, cost, and portIndex if this port has a better root path
			rootId = sw.Nics[i].StpInfo.RootBridgeId
			cost = sw.Nics[i].StpInfo.RootPathCost
			portIndex = i
		}

	}
	sw.StpInfo.mu.Unlock()

	return rootId, cost, portIndex
}

func (sw *Switch) syncPortsSTP() {
	rootId, cost, _ := sw.calculateRootBridgeIdAndCostAndPortIndex()

	sw.StpInfo.mu.Lock()
	sw.StpInfo.RootBridgeId = rootId
	sw.StpInfo.RootPathCost = cost
	sw.StpInfo.mu.Unlock()
}
