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

func (c *Cable) init() {
	c.FrameChan = make(chan cableFrameEvent, 20)
	c.quitChan = make(chan struct{})
}

func (c *Cable) Connect(nic1 *Nic, nic2 *Nic) {
	c.Nics[0] = nic1
	c.Nics[1] = nic2
	nic1.ConnectedCable = c
	nic2.ConnectedCable = c
}

// TransmitFrame enqueues a frame event to the Cable's FrameChan for async delivery
func (c *Cable) TransmitFrame(fromNic *Nic, frame L2Frame) {
	c.FrameChan <- cableFrameEvent{fromNic: fromNic, frame: frame}
}

// Run starts the Cable's goroutine for frame delivery
func (c *Cable) Run() {
	c.init()

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

	fmt.Printf("Cable between %s <-> %s is connected\n", c.Nics[0].ID, c.Nics[1].ID)
}

// deliverFrame handles the actual delivery of a frame from one NIC to the other with delay and logging
func (c *Cable) deliverFrame(event cableFrameEvent) {
	var toNic *Nic
	if c.Nics[0] == event.fromNic {
		toNic = c.Nics[1]
	} else {
		toNic = c.Nics[0]
	}

	if EnableMacLogging {
		fmt.Printf("[%s][%s] Cable transmitting frame: SrcMac=%s, DstMac=%s with 100ms delay\n", time.Now().UTC().Format(time.RFC3339Nano), c.String(), event.frame.SrcMac, event.frame.DstMac)
	}
	time.Sleep(100 * time.Millisecond)
	toNic.EnqueueFrame(event.frame)
}

// Stop signals the Cable goroutine to exit
func (c *Cable) Stop() {
	close(c.quitChan)
}

func (c *Cable) String() string {
	return fmt.Sprintf("Cable(%s <-> %s)", c.Nics[0].ID, c.Nics[1].ID)
}
