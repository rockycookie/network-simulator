package lib

import (
	"fmt"
	"time"
)

type Host struct {
	Name      string
	Nic       Nic
	FrameChan chan L2Frame
	quitChan  chan struct{}
}

func (h *Host) init() {
	// Link the Nic back to this host
	h.Nic.Host = h
	h.FrameChan = make(chan L2Frame, 10)
	h.quitChan = make(chan struct{})
}

// Run starts the host's goroutine to process incoming frames
func (h *Host) Run() {
	h.init()

	// Start the NIC's goroutine
	h.Nic.Run()

	go func() {
		for {
			select {
			case frame := <-h.FrameChan:
				h.receiveFrame(frame)
			case <-h.quitChan:
				fmt.Printf("[Host %s] Shutting down goroutine\n", h.Name)
				return
			}
		}
	}()

	fmt.Printf("[Host %s] (Nic=%s, Mac=%s) Started running\n", h.Name, h.Nic.ID, h.Nic.Mac)
}

func (h *Host) SendFrame(frame L2Frame) {
	h.Nic.SendFrame(frame)
}

func (h *Host) receiveFrame(frame L2Frame) {
	if frame.DstMac != h.Nic.Mac {
		// fmt.Printf("[%s][Host %s] Frame (SrcMac=%s, DstMac=%s) not for this host (DstMac=%s); ignoring.\n",
		// 	time.Now().UTC().Format(time.RFC3339Nano), h.Name, frame.SrcMac, frame.DstMac, h.Nic.Mac)
		return
	}

	// If the frame needs a reply, send a reply frame back
	if frame.NeedReply {
		fmt.Printf("[%s][Host %s] Frame (SrcMac=%s, DstMac=%s) requests reply; sending reply frame.\n",
			time.Now().UTC().Format(time.RFC3339Nano), h.Name, frame.SrcMac, frame.DstMac)

		replyFrame := L2Frame{
			SrcMac:    h.Nic.Mac,
			DstMac:    frame.SrcMac,
			Name:      "Reply to " + frame.Name,
			NeedReply: false,
		}
		h.SendFrame(replyFrame)
	} else {
		fmt.Printf("[%s][Host %s] Received no-reply-required frame (SrcMac=%s, DstMac=%s): %s\n",
			time.Now().UTC().Format(time.RFC3339Nano), h.Name, frame.SrcMac, frame.DstMac, frame.Name)
	}
}

// EnqueueFrame sends a frame to the host's frame channel for async processing
func (h *Host) EnqueueFrame(frame L2Frame) {
	h.FrameChan <- frame
}

// Stop signals the host goroutine to exit
func (h *Host) Stop() {
	close(h.quitChan)
}
