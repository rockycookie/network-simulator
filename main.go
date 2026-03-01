package main

import (
	"flag"
	"fmt"
	"network-simulator/lib"
	"time"
)

func main() {
	// Define and parse command-line flags
	loggingScope := flag.String("log", "none", "Logging scope: mac, stp, all, or none")
	flag.Parse()

	// Set the logging scope
	lib.SetLoggingScope(*loggingScope)

	fmt.Printf("\n===============================================\n")
	fmt.Printf("   STP LOOP-FREE DEMONSTRATION\n")
	fmt.Printf("===============================================\n\n")

	// Create a triangle topology with 4 switches to demonstrate STP
	// Switch IDs determined by lowest MAC: SW1 < SW2 < SW3 < SW4
	// SW1 should become root bridge

	sw1 := lib.Switch{
		Name: "Switch1",
		Nics: []lib.Nic{
			{ID: "sw1-eth0", Mac: "00:00:00:00:01:01"}, // Lowest MAC - will be root
			{ID: "sw1-eth1", Mac: "00:00:00:00:01:02"},
			{ID: "sw1-eth2", Mac: "00:00:00:00:01:03"},
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
			{ID: "sw4-eth1", Mac: "00:00:00:00:04:02"},
		},
	}

	host1 := lib.Host{
		Name: "Host1",
		Nic: lib.Nic{
			ID:  "host1-eth0",
			Mac: "aa:aa:aa:aa:aa:01",
		},
	}
	host2 := lib.Host{
		Name: "Host2",
		Nic: lib.Nic{
			ID:  "host2-eth0",
			Mac: "aa:aa:aa:aa:aa:02",
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

	// host1 <-> SW1
	cableHost1SW1 := &lib.Cable{}
	cableHost1SW1.Connect(&host1.Nic, &sw1.Nics[2])

	// host2 <-> SW4
	cableHost2SW4 := &lib.Cable{}
	cableHost2SW4.Connect(&host2.Nic, &sw4.Nics[1])

	// Start all devices
	fmt.Printf("Starting all devices...\n")
	sw1.Run()
	sw1.Nics[2].StpInfo.PortFast = true // Enable PortFast on host-facing port
	sw2.Run()
	sw3.Run()
	sw4.Run()
	sw4.Nics[1].StpInfo.PortFast = true // Enable PortFast on host-facing port
	host1.Run()
	host2.Run()
	cableSW1SW2.Run()
	cableSW2SW3.Run()
	cableSW1SW3.Run()
	cableSW3SW4.Run()
	cableHost1SW1.Run()
	cableHost2SW4.Run()

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

	// See how host1 round-trips to host2 with the STP blocking in place
	fmt.Printf("\n===============================================\n")
	fmt.Printf("   STP BLOCKING DEMONSTRATION\n")
	fmt.Printf("===============================================\n\n")

	testFrame := lib.L2Frame{
		SrcMac:    host1.Nic.Mac,
		DstMac:    host2.Nic.Mac,
		Name:      "Greeting Hi",
		NeedReply: true,
	}

	fmt.Printf("Host1 sending frame to Host2...\n")
	host1.SendFrame(testFrame)

	// Wait for a moment to allow frame processing and logging
	time.Sleep(10 * time.Second)

	fmt.Printf("\n===============================================\n")
	fmt.Printf("   SIMULATION COMPLETE\n")
	fmt.Printf("===============================================\n\n")

	// Stop all goroutines
	sw1.Stop()
	sw2.Stop()
	sw3.Stop()
	sw4.Stop()
	host1.Stop()
	host2.Stop()
	cableHost1SW1.Stop()
	cableHost2SW4.Stop()
	cableSW1SW2.Stop()
	cableSW2SW3.Stop()
	cableSW1SW3.Stop()
	cableSW3SW4.Stop()

	// wait for a moment to allow all goroutines to print shutdown messages
	time.Sleep(1 * time.Second)
}

func printSwitchStpInfo(sw *lib.Switch) {
	fmt.Printf("Switch: %s\n", sw.Name)
	fmt.Printf("  Bridge ID: %s\n", sw.StpInfo.ID)
	fmt.Printf("  Root Switch ID: %s\n", sw.StpInfo.RootSwitchId)
	fmt.Printf("  Root Path Cost: %d\n", sw.StpInfo.RootPathCost)
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
		if nic.StpInfo.PortFast {
			portType += " (PortFast)"
		}
		fmt.Printf("    Port %d (%s):\n", i, nic.ID)
		fmt.Printf("      Role: %s\n", portType)
		fmt.Printf("      Root Path Cost: %d\n", nic.StpInfo.RootPathCost)
	}
	fmt.Printf("\n")
}
