package lib

// he configuration BPDU consists of the following fields:
// STP type, root path cost, root bridge identifier, local bridgeidentifier, max age, hello time, and forward delay
type ConfigBpdu struct {
	RootPathCost  uint32 // This is the combined cost for a specific path toward the root switch, negative value indicates unknown path cost
	RootSwitchId  string // This is a combination of the root bridge system MAC address, system ID extension, and system priority of the root bridge.
	LocalSwitchId string // This is a combination of the local bridge system MAC address, system ID extension, and system priority of the local bridge that is sending the BPDU
}
