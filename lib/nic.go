package lib

import (
	"fmt"
)

type Nic struct {
	ID             string
	Mac            string
	ConnectedCable *Cable
	IP             string
	Switch         *Switch
	Host           *Host
	StpInfo        PortStp
	FrameChan      chan L2Frame
	quitChan       chan struct{}
}

func (n *Nic) init() {
	n.FrameChan = make(chan L2Frame, 10)
	n.quitChan = make(chan struct{})
}

// Run starts the NIC's goroutine to process incoming frames
func (n *Nic) Run() {
	n.init()

	go func() {
		for {
			select {
			case frame := <-n.FrameChan:
				n.handleFrame(frame)
			case <-n.quitChan:
				fmt.Printf("[NIC %s] Shutting down goroutine\n", n.ID)
				return
			}
		}
	}()
}

// EnqueueFrame sends a frame to the NIC's frame channel for async processing
func (n *Nic) EnqueueFrame(frame L2Frame) {
	n.FrameChan <- frame
}

// Stop signals the NIC goroutine to exit
func (n *Nic) Stop() {
	close(n.quitChan)
}

// handleFrame processes a received frame (internal use)
func (n *Nic) handleFrame(frame L2Frame) {
	if n.Switch != nil {
		n.Switch.EnqueueFrame(frame, n)
	} else if n.Host != nil {
		n.Host.EnqueueFrame(frame)
	}
}

func (n *Nic) SendFrame(frame L2Frame) {
	if n.ConnectedCable != nil {
		n.ConnectedCable.TransmitFrame(n, frame)
	} else {
		fmt.Printf("NIC %s is not connected to any cable; cannot send frame", n.ID)
	}
}
