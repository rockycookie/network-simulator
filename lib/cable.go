package lib

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
	} else if c.Nics[1] == fromNic {
		toNic = c.Nics[0]
	} else {
		// NIC not connected to this cable
		return
	}

	// Deliver the frame to the connected NIC's host or switch
	toNic.ReceiveFrame(frame)
}
