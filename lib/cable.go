package lib

import (
	"fmt"
	"time"
)

// cableFrameEvent is used to pass frames and the sending NIC to the Cable goroutine
type cableFrameEvent struct {
	fromNic *Nic
	frame   L2Frame
}

type Cable struct {
	Nics      [2]*Nic
	FrameChan chan cableFrameEvent
	quitChan  chan struct{}
}

func (c *Cable) Connect(nic1 *Nic, nic2 *Nic) {
	c.Nics[0] = nic1
	c.Nics[1] = nic2
	nic1.ConnectedCable = c
	nic2.ConnectedCable = c

	// Initialize channels if not already
	if c.FrameChan == nil {
		c.FrameChan = make(chan cableFrameEvent, 16)
	}
	if c.quitChan == nil {
		c.quitChan = make(chan struct{})
	}
}

// TransmitFrame enqueues a frame event to the Cable's FrameChan for async delivery
func (c *Cable) TransmitFrame(fromNic *Nic, frame L2Frame) {
	c.FrameChan <- cableFrameEvent{fromNic: fromNic, frame: frame}
}

// Run starts the Cable's goroutine for frame delivery
func (c *Cable) Run() {
	go func() {
		for {
			select {
			case event := <-c.FrameChan:
				c.deliverFrame(event)
			case <-c.quitChan:
				fmt.Printf("Cable between %s and %s shutting down\n", c.Nics[0].ID, c.Nics[1].ID)
				return
			}
		}
	}()
}

// deliverFrame handles the actual delivery of a frame from one NIC to the other with delay and logging
func (c *Cable) deliverFrame(event cableFrameEvent) {
	var toNic *Nic
	if c.Nics[0] == event.fromNic {
		toNic = c.Nics[1]
	} else {
		toNic = c.Nics[0]
	}
	fmt.Printf("Cable transmitting frame: SrcMac=%s, DstMac=%s, from NIC=%s, to NIC=%s with 100ms delay\n", event.frame.SrcMac, event.frame.DstMac, event.fromNic.ID, toNic.ID)
	time.Sleep(100 * time.Millisecond)
	toNic.EnqueueFrame(event.frame)
}

// Stop signals the Cable goroutine to exit
func (c *Cable) Stop() {
	close(c.quitChan)
}
