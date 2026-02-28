package main

import (
	"fmt"
	"network-simulator/lib"
	"time"
)

func main() {
	fmt.Printf("\n===============================================\n")
	fmt.Printf("   STP ROOT ELECTION DEMONSTRATION\n")
	fmt.Printf("===============================================\n\n")

	// Create a triangle topology with 3 switches to demonstrate STP
	// Switch IDs determined by lowest MAC: SW1 < SW2 < SW3
	// SW1 should become root bridge

	sw1 := lib.Switch{
		Name: "Switch1",
		Nics: []lib.Nic{
			{ID: "sw1-eth0", Mac: "00:00:00:00:01:01"}, // Lowest MAC - will be root
			{ID: "sw1-eth1", Mac: "00:00:00:00:01:02"},
		},
	}

	sw2 := lib.Switch{
		Name: "Switch2",
		Nics: []lib.Nic{
			{ID: "sw2-eth0", Mac: "00:00:00:00:02:01"},
			{ID: "sw2-eth1", Mac: "00:00:00:00:02:02"},
		},
	}

	sw3 := lib.Switch{
		Name: "Switch3",
		Nics: []lib.Nic{
			{ID: "sw3-eth0", Mac: "00:00:00:00:03:01"},
			{ID: "sw3-eth1", Mac: "00:00:00:00:03:02"},
			{ID: "sw3-eth2", Mac: "00:00:00:00:03:03"},
		},
	}

	sw4 := lib.Switch{
		Name: "Switch4",
		Nics: []lib.Nic{
			{ID: "sw4-eth0", Mac: "00:00:00:00:04:01"}, // Highest MAC
		},
	}

	// Create cables for triangle topology
	// SW1 <-> SW2
	cableSW1SW2 := &lib.Cable{}
	cableSW1SW2.Connect(&sw1.Nics[0], &sw2.Nics[0])

	// SW2 <-> SW3
	cableSW2SW3 := &lib.Cable{}
	cableSW2SW3.Connect(&sw2.Nics[1], &sw3.Nics[0])

	// SW3 <-> SW1
	cableSW1SW3 := &lib.Cable{}
	cableSW1SW3.Connect(&sw1.Nics[1], &sw3.Nics[1])

	// SW4 <-> SW3
	cableSW3SW4 := &lib.Cable{}
	cableSW3SW4.Connect(&sw3.Nics[2], &sw4.Nics[0])

	// Start all devices
	fmt.Printf("Starting all devices...\n")
	sw1.Run()
	sw2.Run()
	sw3.Run()
	sw4.Run()
	cableSW1SW2.Run()
	cableSW2SW3.Run()
	cableSW1SW3.Run()
	cableSW3SW4.Run()

	fmt.Printf("\n===============================================\n")
	fmt.Printf("   STARTING STP ROOT ELECTION\n")
	fmt.Printf("===============================================\n\n")

	// Start STP on all switches
	sw1.RunStp()
	sw2.RunStp()
	sw3.RunStp()
	sw4.RunStp()

	// Wait for STP convergence (ForwardDelay + some buffer)
	fmt.Printf("Waiting for STP convergence (16 seconds)...\n\n")
	time.Sleep(16 * time.Second)

	// Display STP results
	fmt.Printf("\n===============================================\n")
	fmt.Printf("   STP CONVERGENCE RESULTS\n")
	fmt.Printf("===============================================\n\n")

	printSwitchStpInfo(&sw1)
	printSwitchStpInfo(&sw2)
	printSwitchStpInfo(&sw3)
	printSwitchStpInfo(&sw4)

	fmt.Printf("\n===============================================\n")
	fmt.Printf("   SIMULATION COMPLETE\n")
	fmt.Printf("===============================================\n\n")

	// Stop all goroutines
	sw1.Stop()
	sw2.Stop()
	sw3.Stop()
	cableSW1SW2.Stop()
	cableSW2SW3.Stop()
	cableSW1SW3.Stop()
	cableSW3SW4.Stop()

	// Wait for cleanup
	time.Sleep(1 * time.Second)
}

func printSwitchStpInfo(sw *lib.Switch) {
	fmt.Printf("Switch: %s\n", sw.Name)
	fmt.Printf("  Bridge ID: %s\n", sw.StpInfo.ID)
	fmt.Printf("  Root Bridge ID: %s\n", sw.StpInfo.RootBridgeId)
	fmt.Printf("  STP State: %d\n", sw.StpInfo.State)
	fmt.Printf("  Ports:\n")
	for i := range sw.Nics {
		nic := &sw.Nics[i]
		portType := "UNKNOWN"
		switch nic.StpInfo.Type {
		case lib.PORT_STP_TYPE_ROOT:
			portType = "ROOT"
		case lib.PORT_STP_TYPE_DESIGNATED:
			portType = "DESIGNATED"
		case lib.PORT_STP_TYPE_BLOCKING:
			portType = "BLOCKING"
		}
		fmt.Printf("    Port %d (%s):\n", i, nic.ID)
		fmt.Printf("      Role: %s\n", portType)
		fmt.Printf("      Root Path Cost: %d\n", nic.StpInfo.RootPathCost)
	}
	fmt.Printf("\n")
}
