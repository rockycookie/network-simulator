package main

import (
	"fmt"
	"network-simulator/lib"
)

type Host struct {
	Name string
	Nic  lib.Nic
}

func main() {
	// Specify two hosts
	host1 := Host{
		Name: "Host1",
		Nic:  lib.Nic{ID: "host1-eth0", Mac: "00:11:22:33:44:11"},
	}
	host2 := Host{
		Name: "Host2",
		Nic:  lib.Nic{ID: "host2-eth0", Mac: "00:11:22:33:44:21"},
	}

	// Specify cables
	cable1 := &lib.Cable{}
	cable2 := &lib.Cable{}

	// Specify a switch with 4 ports
	sw := lib.Switch{
		Name: "Switch1",
		Nics: []lib.Nic{
			{ID: "sw-eth0", Mac: "00:11:22:33:44:01"},
			{ID: "sw-eth1", Mac: "00:11:22:33:44:02"},
			{ID: "sw-eth2", Mac: "00:11:22:33:44:03"},
			{ID: "sw-eth3", Mac: "00:11:22:33:44:04"},
		},
	}

	// Connect hosts to switch via cables
	cable1.Connect(&host1.Nic, &sw.Nics[0])
	cable2.Connect(&host2.Nic, &sw.Nics[1])

	fmt.Printf("Switch: %s\n", sw.Name)
	for i, nic := range sw.Nics {
		fmt.Printf("  Nic %d: ID=%s, Mac=%s\n", i+1, nic.ID, nic.Mac)
	}

	fmt.Printf("\nHost connections:\n")
	fmt.Printf("  %s: Nic=%s, Mac=%s\n", host1.Name, host1.Nic.ID, host1.Nic.Mac)
	fmt.Printf("  %s: Nic=%s, Mac=%s\n", host2.Name, host2.Nic.ID, host2.Nic.Mac)

	fmt.Printf("\nCables:\n")
	fmt.Printf("  Cable1: %s <-> %s\n", cable1.Nics[0].ID, cable1.Nics[1].ID)
	fmt.Printf("  Cable2: %s <-> %s\n", cable2.Nics[0].ID, cable2.Nics[1].ID)
}
