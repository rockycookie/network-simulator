package lib

type Nic struct {
	ID             string
	Mac            string
	ConnectedCable *Cable
	IP             string
	Switch         *Switch
	Host           *Host
}

func (n *Nic) SendFrame(frame L2Frame) {
	if n.ConnectedCable != nil {
		n.ConnectedCable.TransmitFrame(n, frame)
	}
}

func (n *Nic) ReceiveFrame(frame L2Frame) {
	if n.Switch != nil {
		n.Switch.ReceiveFrame(frame, *n)
	} else if n.Host != nil {
		n.Host.ReceiveFrame(frame)
	}
}
