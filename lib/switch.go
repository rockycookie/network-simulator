package lib

import (
	"fmt"
	"time"
)

// SwitchFrameEvent is used to pass both a frame and the inbound NIC through the channel
type SwitchFrameEvent struct {
	Frame      L2Frame
	InboundNic *Nic
}

type Switch struct {
	Name            string
	Nics            []Nic
	MacAddressTable map[string]*Nic // MAC to Nic Name mapping
	StpInfo         SwitchStp
	FrameChan       chan SwitchFrameEvent
	quitChan        chan struct{}
}

func (s *Switch) init() {
	s.MacAddressTable = make(map[string]*Nic)
	s.FrameChan = make(chan SwitchFrameEvent, 20)
	s.quitChan = make(chan struct{})

	// Link each Nic back to this switch
	for i := range s.Nics {
		s.Nics[i].Switch = s
	}

	s.initStp()
}

// Run starts the switch's goroutine to process incoming frames
func (s *Switch) Run() {
	s.init()

	// start the NICs' goroutines
	for i := range s.Nics {
		s.Nics[i].Run()
	}

	go func() {
		for {
			select {
			case input := <-s.FrameChan:
				s.ReceiveFrame(input.Frame, input.InboundNic)
			case <-s.quitChan:
				fmt.Printf("[Switch %s] Shutting down goroutine\n", s.Name)
				return
			}
		}
	}()

	fmt.Printf("[Switch %s] is running with %d NICs\n", s.Name, len(s.Nics))
	for i := range s.Nics {
		fmt.Printf("  Nic %d: ID=%s, Mac=%s\n", i, s.Nics[i].ID, s.Nics[i].Mac)
	}
}

func (s *Switch) ReceiveFrame(frame L2Frame, inboundNic *Nic) {
	// fmt.Printf("[%s][Switch %s] Switch received frame: SrcMac=%s, DstMac=%s, on NIC=%s\n", time.Now().UTC().Format(time.RFC3339Nano), s.Name, frame.SrcMac, frame.DstMac, inboundNic.ID)

	// Handle ConfigBpdu frames for STP
	if frame.DstMac == STP_BPDU_DEST_MAC {
		fmt.Printf("[%s][Switch %s] Processing BPDU from %s: RootBridgeId=%s, RootPathCost=%d\n",
			time.Now().UTC().Format(time.RFC3339Nano), s.Name, inboundNic.ID,
			frame.ConfigBpdu.RootBridgeId, frame.ConfigBpdu.RootPathCost)
		s.ProcessConfigBpdu(frame.ConfigBpdu, inboundNic)
		return
	}

	// Only process data frames on ports in FORWARDING or LEARNING state
	if inboundNic.Switch.StpInfo.State != SW_STP_STATE_FORWARDING && inboundNic.Switch.StpInfo.State != SW_STP_STATE_LEARNING {
		fmt.Printf("[%s][Switch %s] is in state %d (not forwarding/learning), dropping frame\n",
			time.Now().UTC().Format(time.RFC3339Nano), s.Name, inboundNic.Switch.StpInfo.State)
		return
	}

	// Learn the source MAC address and associate it with the incoming NIC
	s.MacAddressTable[frame.SrcMac] = inboundNic

	// Handle broadcast frame (destMac all-F)
	if isBroadcastMAC(frame.DstMac) {
		s.broadcastFrame(frame, inboundNic)
		return
	}

	// If destination MAC is unknown, broadcast; otherwise, send to the specific NIC
	outboundNic, found := s.MacAddressTable[frame.DstMac]
	if found {
		s.SendFrame(frame, outboundNic)
	} else {
		fmt.Printf("[%s][Switch %s] Switch does not know destination MAC %s; broadcasting frame\n", time.Now().UTC().Format(time.RFC3339Nano), s.Name, frame.DstMac)
		s.broadcastFrame(frame, inboundNic)
	}
}

// EnqueueFrame sends a frame and inbound NIC to the switch's frame channel for async processing
func (s *Switch) EnqueueFrame(frame L2Frame, inboundNic *Nic) {
	s.FrameChan <- SwitchFrameEvent{Frame: frame, InboundNic: inboundNic}
}

// Stop signals the switch goroutine to exit
func (s *Switch) Stop() {
	close(s.quitChan)
}

// broadcastFrame sends the frame to all Nics except the incoming one
func (s *Switch) broadcastFrame(frame L2Frame, inboundNic *Nic) {
	for i := range s.Nics {
		if s.Nics[i].ID != inboundNic.ID {
			s.SendFrame(frame, &s.Nics[i])
		}
	}
}

func (s *Switch) SendFrame(frame L2Frame, switchNic *Nic) {
	// Allow BPDU frames to be sent regardless of port state
	if frame.ConfigBpdu != nil {
		if switchNic.ConnectedCable != nil {
			fmt.Printf("[%s][Switch %s] Switch sending BPDU via NIC=%s (RootBridgeId=%s, RootPathCost=%d)\n",
				time.Now().UTC().Format(time.RFC3339Nano), s.Name, switchNic.ID,
				frame.ConfigBpdu.RootBridgeId, frame.ConfigBpdu.RootPathCost)
			switchNic.ConnectedCable.TransmitFrame(switchNic, frame)
		}
		return
	}

	// Only forward data frames through ports in FORWARDING state
	if switchNic.Switch.StpInfo.State != SW_STP_STATE_FORWARDING {
		fmt.Printf("[%s][Switch %s] is in state %d (not forwarding), dropping frame\n",
			time.Now().UTC().Format(time.RFC3339Nano), s.Name, switchNic.Switch.StpInfo.State)
		return
	}

	if switchNic.ConnectedCable != nil {
		fmt.Printf("[%s][Switch %s] Switch sending frame: SrcMac=%s, DstMac=%s, via NIC=%s\n", time.Now().UTC().Format(time.RFC3339Nano), s.Name, frame.SrcMac, frame.DstMac, switchNic.ID)
		switchNic.ConnectedCable.TransmitFrame(switchNic, frame)
	}
}

// isBroadcastMAC returns true if the MAC address is all-F (broadcast)
func isBroadcastMAC(mac string) bool {
	return mac == "ff:ff:ff:ff:ff:ff" || mac == "FF:FF:FF:FF:FF:FF"
}
