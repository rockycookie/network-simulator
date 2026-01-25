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
