package lib

import "sync"

const (
	PORT_STP_TYPE_ROOT       = 0
	PORT_STP_TYPE_DESIGNATED = 1
	PORT_STP_TYPE_BLOCKING   = 2

	// The standard destination MAC address for Bridge Protocol Data Units (BPDUs) in STP/RSTP is the IEEE 802.1D multicast address 01:80:C2:00:00:00
	STP_BPDU_DEST_MAC = "01:80:C2:00:00:00"
)

type PortStp struct {
	Type                    uint8
	LinkCost                uint32
	RootSwitchId            string
	RootPathCost            uint32
	OtherSwitchId           string
	OtherSwitchRootPathCost uint32
	PortFast                bool
	mu                      sync.Mutex
}

func (nic *Nic) swLock() {
	// Lock the switch's STP info first, because someone could be
	// reading all ports' STP info to determine port roles and states
	nic.Switch.StpInfo.mu.Lock()
	nic.StpInfo.mu.Lock()
}

func (nic *Nic) swUnlock() {
	nic.StpInfo.mu.Unlock()
	nic.Switch.StpInfo.mu.Unlock()
}

func (p *PortStp) getPortRole() string {
	switch p.Type {
	case PORT_STP_TYPE_ROOT:
		return "ROOT"
	case PORT_STP_TYPE_DESIGNATED:
		return "DESIGNATED"
	case PORT_STP_TYPE_BLOCKING:
		return "BLOCKING"
	default:
		return "UNKNOWN"
	}
}

func (n *Nic) shouldStpDrop(f *L2Frame) bool {
	return !n.StpInfo.PortFast && f.DstMac != STP_BPDU_DEST_MAC && n.StpInfo.Type == PORT_STP_TYPE_BLOCKING
}
