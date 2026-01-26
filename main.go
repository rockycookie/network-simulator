package main

import (
	"fmt"
	"network-simulator/lib"
	"time"
)

func main() {
	// Specify two hosts
	host1 := lib.Host{
		Name: "Host1",
		Nic:  lib.Nic{ID: "host1-eth0", Mac: "00:11:22:33:44:11"},
	}
	host2 := lib.Host{
		Name: "Host2",
		Nic:  lib.Nic{ID: "host2-eth0", Mac: "00:11:22:33:44:21"},
	}
	host3 := lib.Host{
		Name: "Host3",
		Nic:  lib.Nic{ID: "host3-eth0", Mac: "00:11:22:33:44:31"},
	}

	// Specify a switch with 4 ports
	sw := lib.Switch{
		Name: "MySwitch",
		Nics: []lib.Nic{
			{ID: "sw-eth0", Mac: "00:11:22:33:44:01"},
			{ID: "sw-eth1", Mac: "00:11:22:33:44:02"},
			{ID: "sw-eth2", Mac: "00:11:22:33:44:03"},
			{ID: "sw-eth3", Mac: "00:11:22:33:44:04"},
		},
	}

	// Specify cables
	cable1 := &lib.Cable{}
	cable2 := &lib.Cable{}
	cable3 := &lib.Cable{}

	// Connect hosts to switch via cables
	cable1.Connect(&host1.Nic, &sw.Nics[0])
	cable2.Connect(&host2.Nic, &sw.Nics[1])
	cable3.Connect(&host3.Nic, &sw.Nics[2])

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

	fmt.Printf("\n*******************************\nStarting frame transmission simulation\n*******************************\n")
	// Start hosts to process incoming frames
	host1.Run()
	host2.Run()
	host3.Run()
	sw.Run()
	cable1.Run()
	cable2.Run()
	cable3.Run()

	// Host1 sends a frame to Host2
	frame1 := lib.L2Frame{
		SrcMac:    host1.Nic.Mac,
		DstMac:    host2.Nic.Mac,
		Name:      "Greeting",
		NeedReply: true,
	}
	fmt.Printf("\nHost1 sending frame to Host2\n")
	host1.Nic.SendFrame(frame1)

	// sleep for 3 seconds for processing
	time.Sleep(3 * time.Second)

	fmt.Printf("\nSimulation ended.\n")
	// Stop all goroutines
	host1.Stop()
	host2.Stop()
	host3.Stop()
	sw.Stop()
	cable1.Stop()
	cable2.Stop()
	cable3.Stop()

	// wait a moment for goroutines to finish
	time.Sleep(1 * time.Second)
}
