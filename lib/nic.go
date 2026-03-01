package lib

import (
	"fmt"
	"time"
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
		// Drop if STP blocked
		if n.shouldStpDrop(&frame) {
			if EnableStpLogging || EnableMacLogging {
				fmt.Printf("[%s][%s:%s] Received frame to %s but port is in STP BLOCKING state, dropping frame\n",
					time.Now().UTC().Format(time.RFC3339Nano), n.Switch.Name, n.ID, frame.DstMac)
			}
			return
		}

		if EnableMacLogging {
			fmt.Printf("[%s][%s:%s] Received frame: SrcMac=%s, DstMac=%s, Cable=%s\n", time.Now().UTC().Format(time.RFC3339Nano), n.getHostName(), n.ID, frame.SrcMac, frame.DstMac, n.ConnectedCable.String())
		}
		n.Switch.EnqueueFrame(frame, n)
	} else if n.Host != nil {
		n.Host.EnqueueFrame(frame)
	}
}

func (n *Nic) SendFrame(frame L2Frame) {
	if n.ConnectedCable != nil {
		// Drop if STP blocked
		if n.shouldStpDrop(&frame) {
			if EnableStpLogging || EnableMacLogging {
				fmt.Printf("[%s][%s:%s] Port is in STP BLOCKING state, cannot send frame to %s\n",
					time.Now().UTC().Format(time.RFC3339Nano), n.Switch.Name, n.ID, frame.DstMac)
			}
			return
		}

		if EnableMacLogging {
			fmt.Printf("[%s][%s:%s] Sending frame: SrcMac=%s, DstMac=%s, Cable=%s\n", time.Now().UTC().Format(time.RFC3339Nano), n.getHostName(), n.ID, frame.SrcMac, frame.DstMac, n.ConnectedCable.String())
		}
		n.ConnectedCable.TransmitFrame(n, frame)
	} else {
		fmt.Printf("NIC %s is not connected to any cable; cannot send frame", n.ID)
	}
}

func (n *Nic) getHostName() string {
	var nicHostName string = "UnknownHost"
	if n.Host != nil {
		nicHostName = n.Host.Name
	} else if n.Switch != nil {
		nicHostName = n.Switch.Name
	}
	return nicHostName
}
