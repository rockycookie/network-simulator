# 4 Switches STP Root Election

## Feature included
- Each switch port send out Config BPDU with root ID and root path cost every 2 seconds
- Each switch port receives Config BPDU
    1. select the better values
    2. all ports of the switch next sending out Conig BPDU with the better values

So that after the timeout (15 seconds), all switches know 
- who is the root
- min cost to reach the root
- cost to reach root for each ports of the switch

## Execution console output
Notes:
- eventual roles of all ports stay as `BLOCKING` as the relevant feature is not implemented yet
- switch 3 port to switch 4 has cost of 4 which is a bit counter-intuitive, but acceptable

```bash
% ./network-simulator -log stp

===============================================
   STP ROOT ELECTION DEMONSTRATION
===============================================

Starting all devices...
[Switch Switch1] is running with 2 NICs
  Nic 0: ID=sw1-eth0, Mac=00:00:00:00:01:01
  Nic 1: ID=sw1-eth1, Mac=00:00:00:00:01:02
[Switch Switch2] is running with 2 NICs
  Nic 0: ID=sw2-eth0, Mac=00:00:00:00:02:01
  Nic 1: ID=sw2-eth1, Mac=00:00:00:00:02:02
[Switch Switch3] is running with 3 NICs
  Nic 0: ID=sw3-eth0, Mac=00:00:00:00:03:01
  Nic 1: ID=sw3-eth1, Mac=00:00:00:00:03:02
  Nic 2: ID=sw3-eth2, Mac=00:00:00:00:03:03
[Switch Switch4] is running with 1 NICs
  Nic 0: ID=sw4-eth0, Mac=00:00:00:00:04:01
Cable between sw1-eth0 <-> sw2-eth0 is connected
Cable between sw2-eth1 <-> sw3-eth0 is connected
Cable between sw1-eth1 <-> sw3-eth1 is connected
Cable between sw3-eth2 <-> sw4-eth0 is connected

===============================================
   STARTING STP ROOT ELECTION
===============================================

Waiting for STP convergence (16 seconds)...

[2026-03-01T00:06:59.890991Z][Switch Switch4] Sending BPDU on NIC sw4-eth0: RootBridgeId=00:00:00:00:04:01, RootPathCost=0
[2026-03-01T00:06:59.891064Z][Switch Switch1] Sending BPDU on NIC sw1-eth1: RootBridgeId=00:00:00:00:01:01, RootPathCost=0
[2026-03-01T00:06:59.890997Z][Switch Switch2] Sending BPDU on NIC sw2-eth1: RootBridgeId=00:00:00:00:02:01, RootPathCost=0
[2026-03-01T00:06:59.891073Z][Switch Switch1] Sending BPDU on NIC sw1-eth0: RootBridgeId=00:00:00:00:01:01, RootPathCost=0
[2026-03-01T00:06:59.891103Z][Switch Switch3] Sending BPDU on NIC sw3-eth2: RootBridgeId=00:00:00:00:03:01, RootPathCost=0
[2026-03-01T00:06:59.891084Z][Switch Switch3] Sending BPDU on NIC sw3-eth0: RootBridgeId=00:00:00:00:03:01, RootPathCost=0
[2026-03-01T00:06:59.891094Z][Switch Switch3] Sending BPDU on NIC sw3-eth1: RootBridgeId=00:00:00:00:03:01, RootPathCost=0
[2026-03-01T00:06:59.891079Z][Switch Switch2] Sending BPDU on NIC sw2-eth0: RootBridgeId=00:00:00:00:02:01, RootPathCost=0
[2026-03-01T00:06:59.992179Z][Switch Switch3] Received better BPDU on NIC sw3-eth0: RootBridgeId=00:00:00:00:02:01, RootPathCost=1 (was RootBridgeId=00:00:00:00:03:01, RootPathCost=0)
[2026-03-01T00:06:59.99229Z][Switch Switch3] Received better BPDU on NIC sw3-eth1: RootBridgeId=00:00:00:00:01:01, RootPathCost=1 (was RootBridgeId=00:00:00:00:03:01, RootPathCost=0)
[2026-03-01T00:06:59.992312Z][Switch Switch2] Received better BPDU on NIC sw2-eth0: RootBridgeId=00:00:00:00:01:01, RootPathCost=1 (was RootBridgeId=00:00:00:00:02:01, RootPathCost=0)
[2026-03-01T00:07:00.093264Z][Switch Switch4] Received better BPDU on NIC sw4-eth0: RootBridgeId=00:00:00:00:03:01, RootPathCost=1 (was RootBridgeId=00:00:00:00:04:01, RootPathCost=0)
[2026-03-01T00:07:01.892139Z][Switch Switch2] Sending BPDU on NIC sw2-eth0: RootBridgeId=00:00:00:00:01:01, RootPathCost=1
[2026-03-01T00:07:01.892197Z][Switch Switch3] Sending BPDU on NIC sw3-eth2: RootBridgeId=00:00:00:00:01:01, RootPathCost=1
[2026-03-01T00:07:01.892225Z][Switch Switch1] Sending BPDU on NIC sw1-eth1: RootBridgeId=00:00:00:00:01:01, RootPathCost=0
[2026-03-01T00:07:01.892209Z][Switch Switch3] Sending BPDU on NIC sw3-eth1: RootBridgeId=00:00:00:00:01:01, RootPathCost=1
[2026-03-01T00:07:01.892159Z][Switch Switch3] Sending BPDU on NIC sw3-eth0: RootBridgeId=00:00:00:00:01:01, RootPathCost=1
[2026-03-01T00:07:01.892243Z][Switch Switch4] Sending BPDU on NIC sw4-eth0: RootBridgeId=00:00:00:00:03:01, RootPathCost=1
[2026-03-01T00:07:01.892261Z][Switch Switch1] Sending BPDU on NIC sw1-eth0: RootBridgeId=00:00:00:00:01:01, RootPathCost=0
[2026-03-01T00:07:01.892254Z][Switch Switch2] Sending BPDU on NIC sw2-eth1: RootBridgeId=00:00:00:00:01:01, RootPathCost=1
[2026-03-01T00:07:01.993305Z][Switch Switch4] Received better BPDU on NIC sw4-eth0: RootBridgeId=00:00:00:00:01:01, RootPathCost=2 (was RootBridgeId=00:00:00:00:03:01, RootPathCost=1)
[2026-03-01T00:07:01.993292Z][Switch Switch2] Received better BPDU on NIC sw2-eth1: RootBridgeId=00:00:00:00:01:01, RootPathCost=2 (was RootBridgeId=00:00:00:00:02:01, RootPathCost=0)
[2026-03-01T00:07:02.094306Z][Switch Switch3] Received better BPDU on NIC sw3-eth0: RootBridgeId=00:00:00:00:01:01, RootPathCost=2 (was RootBridgeId=00:00:00:00:02:01, RootPathCost=1)
[2026-03-01T00:07:03.892689Z][Switch Switch4] Sending BPDU on NIC sw4-eth0: RootBridgeId=00:00:00:00:01:01, RootPathCost=2
[2026-03-01T00:07:03.892812Z][Switch Switch3] Sending BPDU on NIC sw3-eth1: RootBridgeId=00:00:00:00:01:01, RootPathCost=1
[2026-03-01T00:07:03.892774Z][Switch Switch3] Sending BPDU on NIC sw3-eth2: RootBridgeId=00:00:00:00:01:01, RootPathCost=1
[2026-03-01T00:07:03.892783Z][Switch Switch1] Sending BPDU on NIC sw1-eth0: RootBridgeId=00:00:00:00:01:01, RootPathCost=0
[2026-03-01T00:07:03.892795Z][Switch Switch1] Sending BPDU on NIC sw1-eth1: RootBridgeId=00:00:00:00:01:01, RootPathCost=0
[2026-03-01T00:07:03.892805Z][Switch Switch2] Sending BPDU on NIC sw2-eth0: RootBridgeId=00:00:00:00:01:01, RootPathCost=2
[2026-03-01T00:07:03.892747Z][Switch Switch2] Sending BPDU on NIC sw2-eth1: RootBridgeId=00:00:00:00:01:01, RootPathCost=2
[2026-03-01T00:07:03.892822Z][Switch Switch3] Sending BPDU on NIC sw3-eth0: RootBridgeId=00:00:00:00:01:01, RootPathCost=1
[2026-03-01T00:07:03.993565Z][Switch Switch3] Received better BPDU on NIC sw3-eth2: RootBridgeId=00:00:00:00:01:01, RootPathCost=3 (was RootBridgeId=00:00:00:00:03:01, RootPathCost=0)
[2026-03-01T00:07:05.89395Z][Switch Switch4] Sending BPDU on NIC sw4-eth0: RootBridgeId=00:00:00:00:01:01, RootPathCost=2
[2026-03-01T00:07:05.894012Z][Switch Switch3] Sending BPDU on NIC sw3-eth0: RootBridgeId=00:00:00:00:01:01, RootPathCost=3
[2026-03-01T00:07:05.894039Z][Switch Switch3] Sending BPDU on NIC sw3-eth1: RootBridgeId=00:00:00:00:01:01, RootPathCost=3
[2026-03-01T00:07:05.894048Z][Switch Switch3] Sending BPDU on NIC sw3-eth2: RootBridgeId=00:00:00:00:01:01, RootPathCost=3
[2026-03-01T00:07:05.894055Z][Switch Switch1] Sending BPDU on NIC sw1-eth0: RootBridgeId=00:00:00:00:01:01, RootPathCost=0
[2026-03-01T00:07:05.894064Z][Switch Switch1] Sending BPDU on NIC sw1-eth1: RootBridgeId=00:00:00:00:01:01, RootPathCost=0
[2026-03-01T00:07:05.894071Z][Switch Switch2] Sending BPDU on NIC sw2-eth0: RootBridgeId=00:00:00:00:01:01, RootPathCost=2
[2026-03-01T00:07:05.894075Z][Switch Switch2] Sending BPDU on NIC sw2-eth1: RootBridgeId=00:00:00:00:01:01, RootPathCost=2
[2026-03-01T00:07:07.894555Z][Switch Switch2] Sending BPDU on NIC sw2-eth0: RootBridgeId=00:00:00:00:01:01, RootPathCost=2
[2026-03-01T00:07:07.894617Z][Switch Switch1] Sending BPDU on NIC sw1-eth0: RootBridgeId=00:00:00:00:01:01, RootPathCost=0
[2026-03-01T00:07:07.894633Z][Switch Switch3] Sending BPDU on NIC sw3-eth2: RootBridgeId=00:00:00:00:01:01, RootPathCost=3
[2026-03-01T00:07:07.894607Z][Switch Switch4] Sending BPDU on NIC sw4-eth0: RootBridgeId=00:00:00:00:01:01, RootPathCost=2
[2026-03-01T00:07:07.894613Z][Switch Switch3] Sending BPDU on NIC sw3-eth0: RootBridgeId=00:00:00:00:01:01, RootPathCost=3
[2026-03-01T00:07:07.894577Z][Switch Switch2] Sending BPDU on NIC sw2-eth1: RootBridgeId=00:00:00:00:01:01, RootPathCost=2
[2026-03-01T00:07:07.894625Z][Switch Switch1] Sending BPDU on NIC sw1-eth1: RootBridgeId=00:00:00:00:01:01, RootPathCost=0
[2026-03-01T00:07:07.894599Z][Switch Switch3] Sending BPDU on NIC sw3-eth1: RootBridgeId=00:00:00:00:01:01, RootPathCost=3
[2026-03-01T00:07:09.895511Z][Switch Switch3] Sending BPDU on NIC sw3-eth1: RootBridgeId=00:00:00:00:01:01, RootPathCost=3
[2026-03-01T00:07:09.895805Z][Switch Switch3] Sending BPDU on NIC sw3-eth0: RootBridgeId=00:00:00:00:01:01, RootPathCost=3
[2026-03-01T00:07:09.895816Z][Switch Switch4] Sending BPDU on NIC sw4-eth0: RootBridgeId=00:00:00:00:01:01, RootPathCost=2
[2026-03-01T00:07:09.895829Z][Switch Switch2] Sending BPDU on NIC sw2-eth0: RootBridgeId=00:00:00:00:01:01, RootPathCost=2
[2026-03-01T00:07:09.895839Z][Switch Switch3] Sending BPDU on NIC sw3-eth2: RootBridgeId=00:00:00:00:01:01, RootPathCost=3
[2026-03-01T00:07:09.895843Z][Switch Switch2] Sending BPDU on NIC sw2-eth1: RootBridgeId=00:00:00:00:01:01, RootPathCost=2
[2026-03-01T00:07:09.89585Z][Switch Switch1] Sending BPDU on NIC sw1-eth1: RootBridgeId=00:00:00:00:01:01, RootPathCost=0
[2026-03-01T00:07:09.895859Z][Switch Switch1] Sending BPDU on NIC sw1-eth0: RootBridgeId=00:00:00:00:01:01, RootPathCost=0
[2026-03-01T00:07:11.897035Z][Switch Switch1] Sending BPDU on NIC sw1-eth0: RootBridgeId=00:00:00:00:01:01, RootPathCost=0
[2026-03-01T00:07:11.897202Z][Switch Switch3] Sending BPDU on NIC sw3-eth1: RootBridgeId=00:00:00:00:01:01, RootPathCost=3
[2026-03-01T00:07:11.897225Z][Switch Switch3] Sending BPDU on NIC sw3-eth0: RootBridgeId=00:00:00:00:01:01, RootPathCost=3
[2026-03-01T00:07:11.897241Z][Switch Switch4] Sending BPDU on NIC sw4-eth0: RootBridgeId=00:00:00:00:01:01, RootPathCost=2
[2026-03-01T00:07:11.897259Z][Switch Switch2] Sending BPDU on NIC sw2-eth0: RootBridgeId=00:00:00:00:01:01, RootPathCost=2
[2026-03-01T00:07:11.897383Z][Switch Switch3] Sending BPDU on NIC sw3-eth2: RootBridgeId=00:00:00:00:01:01, RootPathCost=3
[2026-03-01T00:07:11.897396Z][Switch Switch2] Sending BPDU on NIC sw2-eth1: RootBridgeId=00:00:00:00:01:01, RootPathCost=2
[2026-03-01T00:07:11.897406Z][Switch Switch1] Sending BPDU on NIC sw1-eth1: RootBridgeId=00:00:00:00:01:01, RootPathCost=0
[2026-03-01T00:07:13.897598Z][Switch Switch1] Sending BPDU on NIC sw1-eth1: RootBridgeId=00:00:00:00:01:01, RootPathCost=0
[2026-03-01T00:07:13.897629Z][Switch Switch4] Sending BPDU on NIC sw4-eth0: RootBridgeId=00:00:00:00:01:01, RootPathCost=2
[2026-03-01T00:07:13.897642Z][Switch Switch1] Sending BPDU on NIC sw1-eth0: RootBridgeId=00:00:00:00:01:01, RootPathCost=0
[2026-03-01T00:07:13.897648Z][Switch Switch3] Sending BPDU on NIC sw3-eth1: RootBridgeId=00:00:00:00:01:01, RootPathCost=3
[2026-03-01T00:07:13.897651Z][Switch Switch3] Sending BPDU on NIC sw3-eth0: RootBridgeId=00:00:00:00:01:01, RootPathCost=3
[2026-03-01T00:07:13.897655Z][Switch Switch3] Sending BPDU on NIC sw3-eth2: RootBridgeId=00:00:00:00:01:01, RootPathCost=3
[2026-03-01T00:07:13.897659Z][Switch Switch2] Sending BPDU on NIC sw2-eth0: RootBridgeId=00:00:00:00:01:01, RootPathCost=2
[2026-03-01T00:07:13.897661Z][Switch Switch2] Sending BPDU on NIC sw2-eth1: RootBridgeId=00:00:00:00:01:01, RootPathCost=2

===============================================
   STP CONVERGENCE RESULTS
===============================================

Switch: Switch1
  Bridge ID: 00:00:00:00:01:01
  Root Bridge ID: 00:00:00:00:01:01
  Root Path Cost: 0
  STP State: 5
  Ports:
    Port 0 (sw1-eth0):
      Role: BLOCKING
      Root Path Cost: 0
    Port 1 (sw1-eth1):
      Role: BLOCKING
      Root Path Cost: 0

Switch: Switch2
  Bridge ID: 00:00:00:00:02:01
  Root Bridge ID: 00:00:00:00:01:01
  Root Path Cost: 2
  STP State: 5
  Ports:
    Port 0 (sw2-eth0):
      Role: BLOCKING
      Root Path Cost: 1
    Port 1 (sw2-eth1):
      Role: BLOCKING
      Root Path Cost: 2

Switch: Switch3
  Bridge ID: 00:00:00:00:03:01
  Root Bridge ID: 00:00:00:00:01:01
  Root Path Cost: 3
  STP State: 5
  Ports:
    Port 0 (sw3-eth0):
      Role: BLOCKING
      Root Path Cost: 2
    Port 1 (sw3-eth1):
      Role: BLOCKING
      Root Path Cost: 1
    Port 2 (sw3-eth2):
      Role: BLOCKING
      Root Path Cost: 3

Switch: Switch4
  Bridge ID: 00:00:00:00:04:01
  Root Bridge ID: 00:00:00:00:01:01
  Root Path Cost: 2
  STP State: 5
  Ports:
    Port 0 (sw4-eth0):
      Role: BLOCKING
      Root Path Cost: 2


===============================================
   SIMULATION COMPLETE
===============================================

Cable between sw3-eth2 and sw4-eth0 shutting down
[Switch Switch1] Shutting down goroutine
[Switch Switch2] Shutting down goroutine
[Switch Switch3] Shutting down goroutine
[Switch Switch4] Shutting down goroutine
Cable between sw1-eth0 and sw2-eth0 shutting down
Cable between sw2-eth1 and sw3-eth0 shutting down
Cable between sw1-eth1 and sw3-eth1 shutting down
```

## Srouce code setup

<details>
<summary>main.go source code</summary>

```go
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
	sw4.Stop()
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
		fmt.Printf("    Port %d (%s):\n", i, nic.ID)
		fmt.Printf("      Role: %s\n", portType)
		fmt.Printf("      Root Path Cost: %d\n", nic.StpInfo.RootPathCost)
	}
	fmt.Printf("\n")
}

```
</details>
