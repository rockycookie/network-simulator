package lib

import "fmt"

type Switch struct {
	Name            string
	Nics            []Nic
	MacAddressTable map[string]*Nic // MAC to Nic Name mapping
}

func (s *Switch) ReceiveFrame(frame L2Frame, inboundNic Nic) {
	if s.MacAddressTable == nil {
		s.MacAddressTable = make(map[string]*Nic)
	}
	// Learn the source MAC address and associate it with the incoming NIC
	s.MacAddressTable[frame.SrcMac] = &inboundNic

	// Handle broadcast frame (destMac all-F)
	if isBroadcastMAC(frame.DstMac) {
		s.broadcastFrame(frame, inboundNic)
		return
	}

	// If destination MAC is unknown, broadcast; otherwise, send to the specific NIC
	outboundNic, found := s.MacAddressTable[frame.DstMac]
	if found {
		s.SendFrame(frame, *outboundNic)
	} else {
		s.broadcastFrame(frame, inboundNic)
	}
}

// broadcastFrame sends the frame to all Nics except the incoming one
func (s *Switch) broadcastFrame(frame L2Frame, inboundNic Nic) {
	for _, nic := range s.Nics {
		if nic.ID != inboundNic.ID {
			s.SendFrame(frame, nic)
		}
	}
}

func (s *Switch) SendFrame(frame L2Frame, switchNic Nic) {
	// Log the frame sending action
	// Example: Sending frame from SrcMac to DstMac via NIC
	// You may want to adjust the log format as needed
	fmt.Printf("Switch '%s' sending frame: SrcMac=%s, DstMac=%s, via NIC=%s\n", s.Name, frame.SrcMac, frame.DstMac, switchNic.ID)
}

// isBroadcastMAC returns true if the MAC address is all-F (broadcast)
func isBroadcastMAC(mac string) bool {
	return mac == "ff:ff:ff:ff:ff:ff" || mac == "FF:FF:FF:FF:FF:FF"
}
