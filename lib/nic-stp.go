package lib

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
