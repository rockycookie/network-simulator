package lib

import (
	"fmt"
	"time"
)

type Host struct {
	Name string
	Nic  Nic
}

func (h *Host) SendFrame(frame L2Frame) {
	h.Nic.SendFrame(frame)
}

func (h *Host) ReceiveFrame(frame L2Frame) {
	// Print frame reception info with GMT timestamp
	fmt.Printf("[%s][Host %s] Received frame: SrcMac=%s, DstMac=%s\n",
		time.Now().UTC().Format(time.RFC3339), h.Name, frame.SrcMac, frame.DstMac)

	// If the frame needs a reply, send a reply frame back
	if frame.NeedReply {
		replyFrame := L2Frame{
			SrcMac:    h.Nic.Mac,
			DstMac:    frame.SrcMac,
			Name:      "Reply to " + frame.Name,
			NeedReply: false,
		}
		h.SendFrame(replyFrame)
	}
}
