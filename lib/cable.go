package lib

import "fmt"

type Cable struct {
	Nics [2]*Nic
}

func (c *Cable) Connect(nic1 *Nic, nic2 *Nic) {
	c.Nics[0] = nic1
	c.Nics[1] = nic2
	nic1.ConnectedCable = c
	nic2.ConnectedCable = c
}

func (c *Cable) TransmitFrame(fromNic *Nic, frame L2Frame) {
	var toNic *Nic
	if c.Nics[0] == fromNic {
		toNic = c.Nics[1]
	} else {
		toNic = c.Nics[0]
	}

	fmt.Printf("Cable transmitting frame: SrcMac=%s, DstMac=%s, from NIC=%s, to NIC=%s\n", frame.SrcMac, frame.DstMac, fromNic.ID, toNic.ID)

	// Deliver the frame to the connected NIC's host or switch
	toNic.ReceiveFrame(frame)
}
