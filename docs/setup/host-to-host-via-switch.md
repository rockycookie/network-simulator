# Host to Host communication via Switch

## Features included
- Switch (array of NICs)
	- when receiving frames on a NIC/port
		- MAC address table remember MAC-port mapping for incoming frames
		- broadcast for broadcast destination MAC, otherwise
			- Unicast the frame to the port if found in MAC address table
			- Broadcast the frame to all ports except the incoming port when destination MAC is unknown
- Host (single NIC)
	- send frames on its NIC
	- receie frames on its NIC, reply when the frame's reply flag is on (simplified L3 logic)
- NIC (Network Interface Card)
	- Send frames (enqueue to Cable)
	- receive frames (enqueue to its owner device, Host or Switch)
- Cable
	- transmit frames (enqueue to desination NIC)

## Execution console output
```
% go run main.go

*******************************
Starting frame transmission simulation
*******************************
[Host Host1] (Nic=host1-eth0, Mac=00:11:22:33:44:11) Started running
[Host Host2] (Nic=host2-eth0, Mac=00:11:22:33:44:21) Started running
[Host Host3] (Nic=host3-eth0, Mac=00:11:22:33:44:31) Started running
[Switch MySwitch] is running with 4 NICs
  Nic 1: ID=sw-eth0, Mac=00:11:22:33:44:01
  Nic 2: ID=sw-eth1, Mac=00:11:22:33:44:02
  Nic 3: ID=sw-eth2, Mac=00:11:22:33:44:03
  Nic 4: ID=sw-eth3, Mac=00:11:22:33:44:04
Cable between host1-eth0 <-> sw-eth0 is connected
Cable between host2-eth0 <-> sw-eth1 is connected
Cable between host3-eth0 <-> sw-eth2 is connected

*******************************
Host1 sending requireReply frame 'Greeting' to Host2
*******************************
Cable transmitting frame: SrcMac=00:11:22:33:44:11, DstMac=00:11:22:33:44:21, from NIC=host1-eth0, to NIC=sw-eth0 with 100ms delay
[2026-01-26T00:21:52.33822Z][Switch MySwitch] Switch received frame: SrcMac=00:11:22:33:44:11, DstMac=00:11:22:33:44:21, on NIC=sw-eth0
[2026-01-26T00:21:52.338515Z][Switch MySwitch] Switch does not know destination MAC 00:11:22:33:44:21; broadcasting frame
[2026-01-26T00:21:52.338526Z][Switch MySwitch] Switch sending frame: SrcMac=00:11:22:33:44:11, DstMac=00:11:22:33:44:21, via NIC=sw-eth1
[2026-01-26T00:21:52.33855Z][Switch MySwitch] Switch sending frame: SrcMac=00:11:22:33:44:11, DstMac=00:11:22:33:44:21, via NIC=sw-eth2
Cable transmitting frame: SrcMac=00:11:22:33:44:11, DstMac=00:11:22:33:44:21, from NIC=sw-eth2, to NIC=host3-eth0 with 100ms delay
Cable transmitting frame: SrcMac=00:11:22:33:44:11, DstMac=00:11:22:33:44:21, from NIC=sw-eth1, to NIC=host2-eth0 with 100ms delay
[2026-01-26T00:21:52.439768Z][Host Host3] Frame (SrcMac=00:11:22:33:44:11, DstMac=00:11:22:33:44:21) not for this host (DstMac=00:11:22:33:44:31); ignoring.
[2026-01-26T00:21:52.439764Z][Host Host2] Frame (SrcMac=00:11:22:33:44:11, DstMac=00:11:22:33:44:21) requests reply; sending reply frame.
Cable transmitting frame: SrcMac=00:11:22:33:44:21, DstMac=00:11:22:33:44:11, from NIC=host2-eth0, to NIC=sw-eth1 with 100ms delay
[2026-01-26T00:21:52.540352Z][Switch MySwitch] Switch received frame: SrcMac=00:11:22:33:44:21, DstMac=00:11:22:33:44:11, on NIC=sw-eth1
[2026-01-26T00:21:52.540482Z][Switch MySwitch] Switch sending frame: SrcMac=00:11:22:33:44:21, DstMac=00:11:22:33:44:11, via NIC=sw-eth0
Cable transmitting frame: SrcMac=00:11:22:33:44:21, DstMac=00:11:22:33:44:11, from NIC=sw-eth0, to NIC=host1-eth0 with 100ms delay
[2026-01-26T00:21:52.640986Z][Host Host1] Received no-reply-required frame (SrcMac=00:11:22:33:44:21, DstMac=00:11:22:33:44:11): Reply to Greeting

*******************************
Simulation ended
*******************************
Cable between host3-eth0 and sw-eth2 shutting down
[Host Host1] Shutting down goroutine
[Host Host2] Shutting down goroutine
[Host Host3] Shutting down goroutine
[Switch MySwitch] Shutting down goroutine
Cable between host1-eth0 and sw-eth0 shutting down
Cable between host2-eth0 and sw-eth1 shutting down
```

## Srouce code setup

<details>
<summary>main.go source code</summary>

```
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
	fmt.Printf("\n*******************************\nHost1 sending requireReply frame '%s' to Host2\n*******************************\n", frame1.Name)
	host1.Nic.SendFrame(frame1)

	// sleep for 3 seconds for processing
	time.Sleep(3 * time.Second)

	fmt.Printf("\n*******************************\nSimulation ended\n*******************************\n")
	// Stop all goroutines
	host1.Stop()
	host2.Stop()
	host3.Stop()
	sw.Stop()
	cable1.Stop()
	cable2.Stop()
	cable3.Stop()

	// wait a moment for goroutines to finish
	time.Sleep(2 * time.Second)
}
```
</details>