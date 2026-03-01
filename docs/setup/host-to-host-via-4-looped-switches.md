# Host to Host communication via 4 looped switches

## Feature included

On top of [4 Switches STP Root Election](./4-switches-stp-root-election.md), added
- Deciding each switch ports role (ROOT, DESIGNATED, BLOCKING)
- Allow port-fast for endpoint connections
- refined logging and frames sending logic

## Execution console output

### STP logs only
```bash
% ./network-simulator -log stp

===============================================
   STP LOOP-FREE DEMONSTRATION
===============================================

Starting all devices...
[Switch Switch1] is running with 3 NICs
  Nic 0: ID=sw1-eth0, Mac=00:00:00:00:01:01
  Nic 1: ID=sw1-eth1, Mac=00:00:00:00:01:02
  Nic 2: ID=sw1-eth2, Mac=00:00:00:00:01:03
[Switch Switch2] is running with 2 NICs
  Nic 0: ID=sw2-eth0, Mac=00:00:00:00:02:01
  Nic 1: ID=sw2-eth1, Mac=00:00:00:00:02:02
[Switch Switch3] is running with 3 NICs
  Nic 0: ID=sw3-eth0, Mac=00:00:00:00:03:01
  Nic 1: ID=sw3-eth1, Mac=00:00:00:00:03:02
  Nic 2: ID=sw3-eth2, Mac=00:00:00:00:03:03
[Switch Switch4] is running with 2 NICs
  Nic 0: ID=sw4-eth0, Mac=00:00:00:00:04:01
  Nic 1: ID=sw4-eth1, Mac=00:00:00:00:04:02
[Host Host1] (Nic=host1-eth0, Mac=aa:aa:aa:aa:aa:01) Started running
[Host Host2] (Nic=host2-eth0, Mac=aa:aa:aa:aa:aa:02) Started running
Cable between sw1-eth0 <-> sw2-eth0 is connected
Cable between sw2-eth1 <-> sw3-eth0 is connected
Cable between sw1-eth1 <-> sw3-eth1 is connected
Cable between sw3-eth2 <-> sw4-eth0 is connected
Cable between host1-eth0 <-> sw1-eth2 is connected
Cable between host2-eth0 <-> sw4-eth1 is connected

===============================================
   STARTING STP ROOT ELECTION
===============================================

Waiting for STP convergence (16 seconds)...

[2026-03-01T09:52:47.568949Z][Switch1:sw1-eth1] Received BPDU: RootSwitchId=00:00:00:00:03:01, RootPathCost=0, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:52:47.568998Z][Switch4:sw4-eth0] Received BPDU with lower-Root-ID: RootSwitchId=00:00:00:00:03:01, RootPathCost=1 (was RootBridgeId=00:00:00:00:04:01, RootPathCost=0)
[2026-03-01T09:52:47.569127Z][Switch2:sw2-eth1] Received BPDU: RootSwitchId=00:00:00:00:03:01, RootPathCost=0, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:52:47.569168Z][Switch2:sw2-eth0] Received BPDU with lower-Root-ID: RootSwitchId=00:00:00:00:01:01, RootPathCost=1 (was RootBridgeId=00:00:00:00:02:01, RootPathCost=0)
[2026-03-01T09:52:47.669141Z][Switch1:sw1-eth0] Received BPDU: RootSwitchId=00:00:00:00:02:01, RootPathCost=0, LocalSwitchId=00:00:00:00:02:01
[2026-03-01T09:52:47.669176Z][Switch3:sw3-eth1] Received BPDU with lower-Root-ID: RootSwitchId=00:00:00:00:01:01, RootPathCost=1 (was RootBridgeId=00:00:00:00:03:01, RootPathCost=0)
[2026-03-01T09:52:47.669283Z][Switch3:sw3-eth0] Received BPDU with lower-Root-ID: RootSwitchId=00:00:00:00:02:01, RootPathCost=1 (was RootBridgeId=00:00:00:00:03:01, RootPathCost=0)
[2026-03-01T09:52:47.669297Z][Switch3:sw3-eth2] Received BPDU: RootSwitchId=00:00:00:00:04:01, RootPathCost=0, LocalSwitchId=00:00:00:00:04:01
[2026-03-01T09:52:49.569092Z][Switch3:sw3-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=0, LocalSwitchId=00:00:00:00:01:01
[2026-03-01T09:52:49.569185Z][Switch3:sw3-eth0] Received BPDU with lower-Root-ID: RootSwitchId=00:00:00:00:01:01, RootPathCost=2 (was RootBridgeId=00:00:00:00:02:01, RootPathCost=1)
[2026-03-01T09:52:49.569181Z][Switch4:sw4-eth0] Received BPDU with lower-Root-ID: RootSwitchId=00:00:00:00:01:01, RootPathCost=2 (was RootBridgeId=00:00:00:00:03:01, RootPathCost=1)
[2026-03-01T09:52:49.569118Z][Switch1:sw1-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:02:01
[2026-03-01T09:52:49.67021Z][Switch1:sw1-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:52:49.670488Z][Switch2:sw2-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=0, LocalSwitchId=00:00:00:00:01:01
[2026-03-01T09:52:49.670495Z][Switch3:sw3-eth2] Received BPDU: RootSwitchId=00:00:00:00:03:01, RootPathCost=1, LocalSwitchId=00:00:00:00:04:01
[2026-03-01T09:52:49.670513Z][Switch2:sw2-eth1] Received BPDU with lower-Root-ID: RootSwitchId=00:00:00:00:01:01, RootPathCost=2 (was RootBridgeId=00:00:00:00:02:01, RootPathCost=0)
[2026-03-01T09:52:51.570257Z][Switch2:sw2-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=0, LocalSwitchId=00:00:00:00:01:01
[2026-03-01T09:52:51.570279Z][Switch3:sw3-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:02:01
[2026-03-01T09:52:51.570424Z][Switch3:sw3-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=0, LocalSwitchId=00:00:00:00:01:01
[2026-03-01T09:52:51.57032Z][Switch4:sw4-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:52:51.671367Z][Switch1:sw1-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:02:01
[2026-03-01T09:52:51.671537Z][Switch2:sw2-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:52:51.671547Z][Switch3:sw3-eth2] Received BPDU with lower-Root-ID: RootSwitchId=00:00:00:00:01:01, RootPathCost=3 (was RootBridgeId=00:00:00:00:03:01, RootPathCost=0)
[2026-03-01T09:52:51.671562Z][Switch1:sw1-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:52:53.571426Z][Switch4:sw4-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:52:53.571421Z][Switch1:sw1-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:52:53.571473Z][Switch2:sw2-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=0, LocalSwitchId=00:00:00:00:01:01
[2026-03-01T09:52:53.571555Z][Switch2:sw2-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:52:53.672555Z][Switch3:sw3-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:02:01
[2026-03-01T09:52:53.672719Z][Switch3:sw3-eth2] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=2, LocalSwitchId=00:00:00:00:04:01
[2026-03-01T09:52:53.672734Z][Switch3:sw3-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=0, LocalSwitchId=00:00:00:00:01:01
[2026-03-01T09:52:53.672593Z][Switch1:sw1-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:02:01
[2026-03-01T09:52:55.572522Z][Switch3:sw3-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:02:01
[2026-03-01T09:52:55.572656Z][Switch3:sw3-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=0, LocalSwitchId=00:00:00:00:01:01
[2026-03-01T09:52:55.572575Z][Switch4:sw4-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:52:55.572522Z][Switch1:sw1-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:02:01
[2026-03-01T09:52:55.672967Z][Switch1:sw1-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:52:55.672969Z][Switch2:sw2-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:52:55.673124Z][Switch2:sw2-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=0, LocalSwitchId=00:00:00:00:01:01
[2026-03-01T09:52:55.673033Z][Switch3:sw3-eth2] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=2, LocalSwitchId=00:00:00:00:04:01
[2026-03-01T09:52:57.573427Z][Switch1:sw1-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:52:57.573559Z][Switch1:sw1-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:02:01
[2026-03-01T09:52:57.573579Z][Switch2:sw2-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:52:57.573595Z][Switch4:sw4-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:52:57.674523Z][Switch3:sw3-eth2] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=2, LocalSwitchId=00:00:00:00:04:01
[2026-03-01T09:52:57.674598Z][Switch3:sw3-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=0, LocalSwitchId=00:00:00:00:01:01
[2026-03-01T09:52:57.674607Z][Switch3:sw3-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:02:01
[2026-03-01T09:52:57.674561Z][Switch2:sw2-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=0, LocalSwitchId=00:00:00:00:01:01
[2026-03-01T09:52:59.57455Z][Switch1:sw1-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:52:59.574691Z][Switch1:sw1-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:02:01
[2026-03-01T09:52:59.574704Z][Switch2:sw2-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:52:59.574715Z][Switch4:sw4-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:52:59.674767Z][Switch2:sw2-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=0, LocalSwitchId=00:00:00:00:01:01
[2026-03-01T09:52:59.674758Z][Switch3:sw3-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:02:01
[2026-03-01T09:52:59.674829Z][Switch3:sw3-eth2] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=2, LocalSwitchId=00:00:00:00:04:01
[2026-03-01T09:52:59.674833Z][Switch3:sw3-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=0, LocalSwitchId=00:00:00:00:01:01
[2026-03-01T09:53:01.575692Z][Switch1:sw1-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:02:01
[2026-03-01T09:53:01.575826Z][Switch1:sw1-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:53:01.575717Z][Switch3:sw3-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:02:01
[2026-03-01T09:53:01.575849Z][Switch3:sw3-eth2] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=2, LocalSwitchId=00:00:00:00:04:01
[2026-03-01T09:53:01.675881Z][Switch3:sw3-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=0, LocalSwitchId=00:00:00:00:01:01
[2026-03-01T09:53:01.675903Z][Switch2:sw2-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=0, LocalSwitchId=00:00:00:00:01:01
[2026-03-01T09:53:01.675952Z][Switch2:sw2-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:53:01.675914Z][Switch4:sw4-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:53:02.468749Z][Switch1] ------> I am the root switch, all ports set to DESIGNATED.
[2026-03-01T09:53:02.46883Z][Switch2:sw2-eth0] ------> Port set to ROOT (RootSwitchId=00:00:00:00:01:01, RootPathCost=1)
[2026-03-01T09:53:02.46891Z][Switch2:sw2-eth1] ------> Port set to DESIGNATED (LocalRootPathCost=1, OtherSwitchRootPathCost=1, LocalSwitchId=00:00:00:00:02:01, OtherSwitchId=00:00:00:00:03:01)
[2026-03-01T09:53:02.468928Z][Switch3:sw3-eth0] ------> Port set to BLOCKING (LocalRootPathCost=1, OtherSwitchRootPathCost=1, LocalSwitchId=00:00:00:00:03:01, OtherSwitchId=00:00:00:00:02:01)
[2026-03-01T09:53:02.468946Z][Switch3:sw3-eth1] ------> Port set to ROOT (RootSwitchId=00:00:00:00:01:01, RootPathCost=1)
[2026-03-01T09:53:02.468954Z][Switch3:sw3-eth2] ------> Port set to DESIGNATED (LocalRootPathCost=1, OtherSwitchRootPathCost=2, LocalSwitchId=00:00:00:00:03:01, OtherSwitchId=00:00:00:00:04:01)
[2026-03-01T09:53:02.468862Z][Switch4:sw4-eth0] ------> Port set to ROOT (RootSwitchId=00:00:00:00:01:01, RootPathCost=2)
[2026-03-01T09:53:02.469064Z][Switch4:sw4-eth1] ------> Port set to BLOCKING (LocalRootPathCost=2, OtherSwitchRootPathCost=0, LocalSwitchId=00:00:00:00:04:01, OtherSwitchId=)

===============================================
   STP CONVERGENCE RESULTS
===============================================

Switch: Switch1
  Bridge ID: 00:00:00:00:01:01
  Root Switch ID: 00:00:00:00:01:01
  Root Path Cost: 0
  STP State: 4
  Ports:
    Port 0 (sw1-eth0):
      Role: DESIGNATED
      Root Path Cost: 0
    Port 1 (sw1-eth1):
      Role: DESIGNATED
      Root Path Cost: 0
    Port 2 (sw1-eth2):
      Role: DESIGNATED (PortFast)
      Root Path Cost: 0

Switch: Switch2
  Bridge ID: 00:00:00:00:02:01
  Root Switch ID: 00:00:00:00:01:01
  Root Path Cost: 1
  STP State: 4
  Ports:
    Port 0 (sw2-eth0):
      Role: ROOT
      Root Path Cost: 1
    Port 1 (sw2-eth1):
      Role: DESIGNATED
      Root Path Cost: 2

Switch: Switch3
  Bridge ID: 00:00:00:00:03:01
  Root Switch ID: 00:00:00:00:01:01
  Root Path Cost: 1
  STP State: 4
  Ports:
    Port 0 (sw3-eth0):
      Role: BLOCKING
      Root Path Cost: 2
    Port 1 (sw3-eth1):
      Role: ROOT
      Root Path Cost: 1
    Port 2 (sw3-eth2):
      Role: DESIGNATED
      Root Path Cost: 3

Switch: Switch4
  Bridge ID: 00:00:00:00:04:01
  Root Switch ID: 00:00:00:00:01:01
  Root Path Cost: 2
  STP State: 4
  Ports:
    Port 0 (sw4-eth0):
      Role: ROOT
      Root Path Cost: 2
    Port 1 (sw4-eth1):
      Role: BLOCKING (PortFast)
      Root Path Cost: 0


===============================================
   STP BLOCKING DEMONSTRATION
===============================================

Host1 sending frame to Host2...
[2026-03-01T09:53:03.471672Z][Host1] Sending frame (SrcMac=aa:aa:aa:aa:aa:01, DstMac=aa:aa:aa:aa:aa:02, Name=Greeting Hi, NeedReply=true)
[2026-03-01T09:53:03.674022Z][Switch3:sw3-eth0] Port is in STP BLOCKING state, cannot send frame to aa:aa:aa:aa:aa:02
[2026-03-01T09:53:03.775143Z][Switch3:sw3-eth0] Received frame to aa:aa:aa:aa:aa:02 but port is in STP BLOCKING state, dropping frame
[2026-03-01T09:53:03.876245Z][Host2] Received frame (SrcMac=aa:aa:aa:aa:aa:01, DstMac=aa:aa:aa:aa:aa:02, Name=Greeting Hi, NeedReply=true)
[2026-03-01T09:53:03.87641Z][Host2] Sending frame (SrcMac=aa:aa:aa:aa:aa:02, DstMac=aa:aa:aa:aa:aa:01, Name=Reply to Greeting Hi, NeedReply=false)
[2026-03-01T09:53:04.27974Z][Host1] Received frame (SrcMac=aa:aa:aa:aa:aa:02, DstMac=aa:aa:aa:aa:aa:01, Name=Reply to Greeting Hi, NeedReply=false)

===============================================
   SIMULATION COMPLETE
===============================================

Cable between sw3-eth2 and sw4-eth0 shutting down
[Switch Switch1] Shutting down goroutine
[Switch Switch2] Shutting down goroutine
[Switch Switch3] Shutting down goroutine
[Switch Switch4] Shutting down goroutine
[Host Host1] Shutting down goroutine
[Host Host2] Shutting down goroutine
Cable between host1-eth0 and sw1-eth2 shutting down
Cable between host2-eth0 and sw4-eth1 shutting down
Cable between sw1-eth0 and sw2-eth0 shutting down
Cable between sw2-eth1 and sw3-eth0 shutting down
Cable between sw1-eth1 and sw3-eth1 shutting down
```

### All logs (STP + Ethernet)
<details>

```bash
% ./network-simulator -log all

===============================================
   STP LOOP-FREE DEMONSTRATION
===============================================

Starting all devices...
[Switch Switch1] is running with 3 NICs
  Nic 0: ID=sw1-eth0, Mac=00:00:00:00:01:01
  Nic 1: ID=sw1-eth1, Mac=00:00:00:00:01:02
  Nic 2: ID=sw1-eth2, Mac=00:00:00:00:01:03
[Switch Switch2] is running with 2 NICs
  Nic 0: ID=sw2-eth0, Mac=00:00:00:00:02:01
  Nic 1: ID=sw2-eth1, Mac=00:00:00:00:02:02
[Switch Switch3] is running with 3 NICs
  Nic 0: ID=sw3-eth0, Mac=00:00:00:00:03:01
  Nic 1: ID=sw3-eth1, Mac=00:00:00:00:03:02
  Nic 2: ID=sw3-eth2, Mac=00:00:00:00:03:03
[Switch Switch4] is running with 2 NICs
  Nic 0: ID=sw4-eth0, Mac=00:00:00:00:04:01
  Nic 1: ID=sw4-eth1, Mac=00:00:00:00:04:02
[Host Host1] (Nic=host1-eth0, Mac=aa:aa:aa:aa:aa:01) Started running
[Host Host2] (Nic=host2-eth0, Mac=aa:aa:aa:aa:aa:02) Started running
Cable between sw1-eth0 <-> sw2-eth0 is connected
Cable between sw2-eth1 <-> sw3-eth0 is connected
Cable between sw1-eth1 <-> sw3-eth1 is connected
Cable between sw3-eth2 <-> sw4-eth0 is connected
Cable between host1-eth0 <-> sw1-eth2 is connected
Cable between host2-eth0 <-> sw4-eth1 is connected

===============================================
   STARTING STP ROOT ELECTION
===============================================

Waiting for STP convergence (16 seconds)...

[2026-03-01T09:55:27.72806Z][Cable(host2-eth0 <-> sw4-eth1)] Cable transmitting frame: SrcMac=00:00:00:00:04:02, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:27.728096Z][Cable(sw3-eth2 <-> sw4-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:04:01, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:27.728109Z][Cable(sw1-eth1 <-> sw3-eth1)] Cable transmitting frame: SrcMac=00:00:00:00:03:02, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:27.728071Z][Cable(sw2-eth1 <-> sw3-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:02:02, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:27.728129Z][Cable(sw1-eth0 <-> sw2-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:01:01, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:27.728115Z][Cable(host1-eth0 <-> sw1-eth2)] Cable transmitting frame: SrcMac=00:00:00:00:01:03, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:27.829245Z][Cable(sw1-eth1 <-> sw3-eth1)] Cable transmitting frame: SrcMac=00:00:00:00:01:02, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:27.829299Z][Cable(sw2-eth1 <-> sw3-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:03:01, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:27.829302Z][Host1] Received frame (SrcMac=00:00:00:00:01:03, DstMac=01:80:C2:00:00:00, Name=, NeedReply=false) not for this host (DstMac=aa:aa:aa:aa:aa:01); ignoring.
[2026-03-01T09:55:27.829302Z][Host2] Received frame (SrcMac=00:00:00:00:04:02, DstMac=01:80:C2:00:00:00, Name=, NeedReply=false) not for this host (DstMac=aa:aa:aa:aa:aa:02); ignoring.
[2026-03-01T09:55:27.829308Z][Switch3:sw3-eth0] Received frame: SrcMac=00:00:00:00:02:02, DstMac=01:80:C2:00:00:00, Cable=Cable(sw2-eth1 <-> sw3-eth0)
[2026-03-01T09:55:27.829318Z][Cable(sw3-eth2 <-> sw4-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:03:03, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:27.829316Z][Cable(sw1-eth0 <-> sw2-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:02:01, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:27.829318Z][Switch1:sw1-eth1] Received frame: SrcMac=00:00:00:00:03:02, DstMac=01:80:C2:00:00:00, Cable=Cable(sw1-eth1 <-> sw3-eth1)
[2026-03-01T09:55:27.829429Z][Switch1:sw1-eth1] Received BPDU: RootSwitchId=00:00:00:00:03:01, RootPathCost=0, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:55:27.829332Z][Switch2:sw2-eth0] Received frame: SrcMac=00:00:00:00:01:01, DstMac=01:80:C2:00:00:00, Cable=Cable(sw1-eth0 <-> sw2-eth0)
[2026-03-01T09:55:27.829342Z][Switch3:sw3-eth2] Received frame: SrcMac=00:00:00:00:04:01, DstMac=01:80:C2:00:00:00, Cable=Cable(sw3-eth2 <-> sw4-eth0)
[2026-03-01T09:55:27.829334Z][Switch3:sw3-eth0] Received BPDU with lower-Root-ID: RootSwitchId=00:00:00:00:02:01, RootPathCost=1 (was RootBridgeId=00:00:00:00:03:01, RootPathCost=0)
[2026-03-01T09:55:27.829502Z][Switch3:sw3-eth2] Received BPDU: RootSwitchId=00:00:00:00:04:01, RootPathCost=0, LocalSwitchId=00:00:00:00:04:01
[2026-03-01T09:55:27.829466Z][Switch2:sw2-eth0] Received BPDU with lower-Root-ID: RootSwitchId=00:00:00:00:01:01, RootPathCost=1 (was RootBridgeId=00:00:00:00:02:01, RootPathCost=0)
[2026-03-01T09:55:27.930454Z][Switch2:sw2-eth1] Received frame: SrcMac=00:00:00:00:03:01, DstMac=01:80:C2:00:00:00, Cable=Cable(sw2-eth1 <-> sw3-eth0)
[2026-03-01T09:55:27.930459Z][Switch3:sw3-eth1] Received frame: SrcMac=00:00:00:00:01:02, DstMac=01:80:C2:00:00:00, Cable=Cable(sw1-eth1 <-> sw3-eth1)
[2026-03-01T09:55:27.930595Z][Switch3:sw3-eth1] Received BPDU with lower-Root-ID: RootSwitchId=00:00:00:00:01:01, RootPathCost=1 (was RootBridgeId=00:00:00:00:03:01, RootPathCost=0)
[2026-03-01T09:55:27.930535Z][Switch4:sw4-eth0] Received frame: SrcMac=00:00:00:00:03:03, DstMac=01:80:C2:00:00:00, Cable=Cable(sw3-eth2 <-> sw4-eth0)
[2026-03-01T09:55:27.930676Z][Switch4:sw4-eth0] Received BPDU with lower-Root-ID: RootSwitchId=00:00:00:00:03:01, RootPathCost=1 (was RootBridgeId=00:00:00:00:04:01, RootPathCost=0)
[2026-03-01T09:55:27.930523Z][Switch1:sw1-eth0] Received frame: SrcMac=00:00:00:00:02:01, DstMac=01:80:C2:00:00:00, Cable=Cable(sw1-eth0 <-> sw2-eth0)
[2026-03-01T09:55:27.930695Z][Switch1:sw1-eth0] Received BPDU: RootSwitchId=00:00:00:00:02:01, RootPathCost=0, LocalSwitchId=00:00:00:00:02:01
[2026-03-01T09:55:27.930581Z][Switch2:sw2-eth1] Received BPDU: RootSwitchId=00:00:00:00:03:01, RootPathCost=0, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:55:29.728546Z][Cable(host1-eth0 <-> sw1-eth2)] Cable transmitting frame: SrcMac=00:00:00:00:01:03, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:29.728544Z][Cable(sw1-eth0 <-> sw2-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:02:01, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:29.728641Z][Cable(sw1-eth1 <-> sw3-eth1)] Cable transmitting frame: SrcMac=00:00:00:00:01:02, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:29.728701Z][Cable(sw3-eth2 <-> sw4-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:04:01, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:29.72872Z][Cable(sw2-eth1 <-> sw3-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:03:01, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:29.728689Z][Cable(host2-eth0 <-> sw4-eth1)] Cable transmitting frame: SrcMac=00:00:00:00:04:02, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:29.829713Z][Cable(sw3-eth2 <-> sw4-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:03:03, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:29.82985Z][Switch3:sw3-eth2] Received frame: SrcMac=00:00:00:00:04:01, DstMac=01:80:C2:00:00:00, Cable=Cable(sw3-eth2 <-> sw4-eth0)
[2026-03-01T09:55:29.829876Z][Switch3:sw3-eth2] Received BPDU: RootSwitchId=00:00:00:00:03:01, RootPathCost=1, LocalSwitchId=00:00:00:00:04:01
[2026-03-01T09:55:29.82998Z][Host2] Received frame (SrcMac=00:00:00:00:04:02, DstMac=01:80:C2:00:00:00, Name=, NeedReply=false) not for this host (DstMac=aa:aa:aa:aa:aa:02); ignoring.
[2026-03-01T09:55:29.830014Z][Host1] Received frame (SrcMac=00:00:00:00:01:03, DstMac=01:80:C2:00:00:00, Name=, NeedReply=false) not for this host (DstMac=aa:aa:aa:aa:aa:01); ignoring.
[2026-03-01T09:55:29.830033Z][Cable(sw1-eth0 <-> sw2-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:01:01, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:29.830048Z][Switch1:sw1-eth0] Received frame: SrcMac=00:00:00:00:02:01, DstMac=01:80:C2:00:00:00, Cable=Cable(sw1-eth0 <-> sw2-eth0)
[2026-03-01T09:55:29.830064Z][Switch1:sw1-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:02:01
[2026-03-01T09:55:29.830078Z][Cable(sw1-eth1 <-> sw3-eth1)] Cable transmitting frame: SrcMac=00:00:00:00:03:02, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:29.830089Z][Switch3:sw3-eth1] Received frame: SrcMac=00:00:00:00:01:02, DstMac=01:80:C2:00:00:00, Cable=Cable(sw1-eth1 <-> sw3-eth1)
[2026-03-01T09:55:29.830105Z][Switch3:sw3-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=0, LocalSwitchId=00:00:00:00:01:01
[2026-03-01T09:55:29.830117Z][Cable(sw2-eth1 <-> sw3-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:02:02, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:29.830131Z][Switch2:sw2-eth1] Received frame: SrcMac=00:00:00:00:03:01, DstMac=01:80:C2:00:00:00, Cable=Cable(sw2-eth1 <-> sw3-eth0)
[2026-03-01T09:55:29.830147Z][Switch2:sw2-eth1] Received BPDU with lower-Root-ID: RootSwitchId=00:00:00:00:01:01, RootPathCost=2 (was RootBridgeId=00:00:00:00:02:01, RootPathCost=0)
[2026-03-01T09:55:29.931049Z][Switch4:sw4-eth0] Received frame: SrcMac=00:00:00:00:03:03, DstMac=01:80:C2:00:00:00, Cable=Cable(sw3-eth2 <-> sw4-eth0)
[2026-03-01T09:55:29.931245Z][Switch4:sw4-eth0] Received BPDU with lower-Root-ID: RootSwitchId=00:00:00:00:01:01, RootPathCost=2 (was RootBridgeId=00:00:00:00:03:01, RootPathCost=1)
[2026-03-01T09:55:29.931321Z][Switch3:sw3-eth0] Received frame: SrcMac=00:00:00:00:02:02, DstMac=01:80:C2:00:00:00, Cable=Cable(sw2-eth1 <-> sw3-eth0)
[2026-03-01T09:55:29.931343Z][Switch3:sw3-eth0] Received BPDU with lower-Root-ID: RootSwitchId=00:00:00:00:01:01, RootPathCost=2 (was RootBridgeId=00:00:00:00:02:01, RootPathCost=1)
[2026-03-01T09:55:29.931362Z][Switch2:sw2-eth0] Received frame: SrcMac=00:00:00:00:01:01, DstMac=01:80:C2:00:00:00, Cable=Cable(sw1-eth0 <-> sw2-eth0)
[2026-03-01T09:55:29.931377Z][Switch2:sw2-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=0, LocalSwitchId=00:00:00:00:01:01
[2026-03-01T09:55:29.931396Z][Switch1:sw1-eth1] Received frame: SrcMac=00:00:00:00:03:02, DstMac=01:80:C2:00:00:00, Cable=Cable(sw1-eth1 <-> sw3-eth1)
[2026-03-01T09:55:29.931418Z][Switch1:sw1-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:55:31.729687Z][Cable(sw1-eth0 <-> sw2-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:02:01, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:31.729677Z][Cable(sw3-eth2 <-> sw4-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:03:03, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:31.729716Z][Cable(sw1-eth1 <-> sw3-eth1)] Cable transmitting frame: SrcMac=00:00:00:00:03:02, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:31.729732Z][Cable(sw2-eth1 <-> sw3-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:02:02, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:31.72975Z][Cable(host2-eth0 <-> sw4-eth1)] Cable transmitting frame: SrcMac=00:00:00:00:04:02, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:31.729771Z][Cable(host1-eth0 <-> sw1-eth2)] Cable transmitting frame: SrcMac=00:00:00:00:01:03, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:31.830909Z][Host1] Received frame (SrcMac=00:00:00:00:01:03, DstMac=01:80:C2:00:00:00, Name=, NeedReply=false) not for this host (DstMac=aa:aa:aa:aa:aa:01); ignoring.
[2026-03-01T09:55:31.830932Z][Cable(sw1-eth1 <-> sw3-eth1)] Cable transmitting frame: SrcMac=00:00:00:00:01:02, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:31.831045Z][Switch3:sw3-eth0] Received frame: SrcMac=00:00:00:00:02:02, DstMac=01:80:C2:00:00:00, Cable=Cable(sw2-eth1 <-> sw3-eth0)
[2026-03-01T09:55:31.831063Z][Host2] Received frame (SrcMac=00:00:00:00:04:02, DstMac=01:80:C2:00:00:00, Name=, NeedReply=false) not for this host (DstMac=aa:aa:aa:aa:aa:02); ignoring.
[2026-03-01T09:55:31.831007Z][Switch1:sw1-eth0] Received frame: SrcMac=00:00:00:00:02:01, DstMac=01:80:C2:00:00:00, Cable=Cable(sw1-eth0 <-> sw2-eth0)
[2026-03-01T09:55:31.831075Z][Switch3:sw3-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:02:01
[2026-03-01T09:55:31.830998Z][Cable(sw1-eth0 <-> sw2-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:01:01, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:31.831032Z][Cable(sw2-eth1 <-> sw3-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:03:01, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:31.830952Z][Switch1:sw1-eth1] Received frame: SrcMac=00:00:00:00:03:02, DstMac=01:80:C2:00:00:00, Cable=Cable(sw1-eth1 <-> sw3-eth1)
[2026-03-01T09:55:31.831088Z][Switch1:sw1-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:02:01
[2026-03-01T09:55:31.831347Z][Switch1:sw1-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:55:31.831022Z][Switch4:sw4-eth0] Received frame: SrcMac=00:00:00:00:03:03, DstMac=01:80:C2:00:00:00, Cable=Cable(sw3-eth2 <-> sw4-eth0)
[2026-03-01T09:55:31.831367Z][Switch4:sw4-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:55:31.831017Z][Cable(sw3-eth2 <-> sw4-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:04:01, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:31.932161Z][Switch3:sw3-eth2] Received frame: SrcMac=00:00:00:00:04:01, DstMac=01:80:C2:00:00:00, Cable=Cable(sw3-eth2 <-> sw4-eth0)
[2026-03-01T09:55:31.932314Z][Switch3:sw3-eth2] Received BPDU with lower-Root-ID: RootSwitchId=00:00:00:00:01:01, RootPathCost=3 (was RootBridgeId=00:00:00:00:03:01, RootPathCost=0)
[2026-03-01T09:55:31.932209Z][Switch3:sw3-eth1] Received frame: SrcMac=00:00:00:00:01:02, DstMac=01:80:C2:00:00:00, Cable=Cable(sw1-eth1 <-> sw3-eth1)
[2026-03-01T09:55:31.93235Z][Switch3:sw3-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=0, LocalSwitchId=00:00:00:00:01:01
[2026-03-01T09:55:31.93217Z][Switch2:sw2-eth0] Received frame: SrcMac=00:00:00:00:01:01, DstMac=01:80:C2:00:00:00, Cable=Cable(sw1-eth0 <-> sw2-eth0)
[2026-03-01T09:55:31.932377Z][Switch2:sw2-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=0, LocalSwitchId=00:00:00:00:01:01
[2026-03-01T09:55:31.932225Z][Switch2:sw2-eth1] Received frame: SrcMac=00:00:00:00:03:01, DstMac=01:80:C2:00:00:00, Cable=Cable(sw2-eth1 <-> sw3-eth0)
[2026-03-01T09:55:31.932396Z][Switch2:sw2-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:55:33.730917Z][Cable(sw1-eth0 <-> sw2-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:02:01, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:33.731047Z][Cable(sw3-eth2 <-> sw4-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:04:01, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:33.731183Z][Cable(host1-eth0 <-> sw1-eth2)] Cable transmitting frame: SrcMac=00:00:00:00:01:03, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:33.731215Z][Cable(sw1-eth1 <-> sw3-eth1)] Cable transmitting frame: SrcMac=00:00:00:00:03:02, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:33.731231Z][Cable(sw2-eth1 <-> sw3-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:02:02, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:33.731323Z][Cable(host2-eth0 <-> sw4-eth1)] Cable transmitting frame: SrcMac=00:00:00:00:04:02, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:33.831325Z][Cable(sw2-eth1 <-> sw3-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:03:01, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:33.831435Z][Switch1:sw1-eth1] Received frame: SrcMac=00:00:00:00:03:02, DstMac=01:80:C2:00:00:00, Cable=Cable(sw1-eth1 <-> sw3-eth1)
[2026-03-01T09:55:33.83146Z][Host2] Received frame (SrcMac=00:00:00:00:04:02, DstMac=01:80:C2:00:00:00, Name=, NeedReply=false) not for this host (DstMac=aa:aa:aa:aa:aa:02); ignoring.
[2026-03-01T09:55:33.831378Z][Cable(sw1-eth0 <-> sw2-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:01:01, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:33.831388Z][Switch1:sw1-eth0] Received frame: SrcMac=00:00:00:00:02:01, DstMac=01:80:C2:00:00:00, Cable=Cable(sw1-eth0 <-> sw2-eth0)
[2026-03-01T09:55:33.831415Z][Host1] Received frame (SrcMac=00:00:00:00:01:03, DstMac=01:80:C2:00:00:00, Name=, NeedReply=false) not for this host (DstMac=aa:aa:aa:aa:aa:01); ignoring.
[2026-03-01T09:55:33.831426Z][Cable(sw1-eth1 <-> sw3-eth1)] Cable transmitting frame: SrcMac=00:00:00:00:01:02, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:33.831465Z][Switch1:sw1-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:55:33.831681Z][Switch1:sw1-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:02:01
[2026-03-01T09:55:33.831304Z][Cable(sw3-eth2 <-> sw4-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:03:03, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:33.831479Z][Switch3:sw3-eth0] Received frame: SrcMac=00:00:00:00:02:02, DstMac=01:80:C2:00:00:00, Cable=Cable(sw2-eth1 <-> sw3-eth0)
[2026-03-01T09:55:33.831712Z][Switch3:sw3-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:02:01
[2026-03-01T09:55:33.831362Z][Switch3:sw3-eth2] Received frame: SrcMac=00:00:00:00:04:01, DstMac=01:80:C2:00:00:00, Cable=Cable(sw3-eth2 <-> sw4-eth0)
[2026-03-01T09:55:33.83173Z][Switch3:sw3-eth2] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=2, LocalSwitchId=00:00:00:00:04:01
[2026-03-01T09:55:33.932564Z][Switch4:sw4-eth0] Received frame: SrcMac=00:00:00:00:03:03, DstMac=01:80:C2:00:00:00, Cable=Cable(sw3-eth2 <-> sw4-eth0)
[2026-03-01T09:55:33.932637Z][Switch4:sw4-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:55:33.932596Z][Switch2:sw2-eth1] Received frame: SrcMac=00:00:00:00:03:01, DstMac=01:80:C2:00:00:00, Cable=Cable(sw2-eth1 <-> sw3-eth0)
[2026-03-01T09:55:33.932654Z][Switch2:sw2-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:55:33.932603Z][Switch3:sw3-eth1] Received frame: SrcMac=00:00:00:00:01:02, DstMac=01:80:C2:00:00:00, Cable=Cable(sw1-eth1 <-> sw3-eth1)
[2026-03-01T09:55:33.932586Z][Switch2:sw2-eth0] Received frame: SrcMac=00:00:00:00:01:01, DstMac=01:80:C2:00:00:00, Cable=Cable(sw1-eth0 <-> sw2-eth0)
[2026-03-01T09:55:33.932663Z][Switch3:sw3-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=0, LocalSwitchId=00:00:00:00:01:01
[2026-03-01T09:55:33.932679Z][Switch2:sw2-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=0, LocalSwitchId=00:00:00:00:01:01
[2026-03-01T09:55:35.732056Z][Cable(sw2-eth1 <-> sw3-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:03:01, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:35.732113Z][Cable(host1-eth0 <-> sw1-eth2)] Cable transmitting frame: SrcMac=00:00:00:00:01:03, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:35.732131Z][Cable(sw3-eth2 <-> sw4-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:03:03, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:35.732144Z][Cable(sw1-eth1 <-> sw3-eth1)] Cable transmitting frame: SrcMac=00:00:00:00:03:02, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:35.732158Z][Cable(sw1-eth0 <-> sw2-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:01:01, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:35.732173Z][Cable(host2-eth0 <-> sw4-eth1)] Cable transmitting frame: SrcMac=00:00:00:00:04:02, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:35.833345Z][Cable(sw3-eth2 <-> sw4-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:04:01, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:35.833345Z][Host2] Received frame (SrcMac=00:00:00:00:04:02, DstMac=01:80:C2:00:00:00, Name=, NeedReply=false) not for this host (DstMac=aa:aa:aa:aa:aa:02); ignoring.
[2026-03-01T09:55:35.833389Z][Cable(sw1-eth1 <-> sw3-eth1)] Cable transmitting frame: SrcMac=00:00:00:00:01:02, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:35.8334Z][Switch1:sw1-eth1] Received frame: SrcMac=00:00:00:00:03:02, DstMac=01:80:C2:00:00:00, Cable=Cable(sw1-eth1 <-> sw3-eth1)
[2026-03-01T09:55:35.833413Z][Cable(sw1-eth0 <-> sw2-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:02:01, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:35.83342Z][Switch2:sw2-eth0] Received frame: SrcMac=00:00:00:00:01:01, DstMac=01:80:C2:00:00:00, Cable=Cable(sw1-eth0 <-> sw2-eth0)
[2026-03-01T09:55:35.833543Z][Switch2:sw2-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=0, LocalSwitchId=00:00:00:00:01:01
[2026-03-01T09:55:35.833441Z][Switch2:sw2-eth1] Received frame: SrcMac=00:00:00:00:03:01, DstMac=01:80:C2:00:00:00, Cable=Cable(sw2-eth1 <-> sw3-eth0)
[2026-03-01T09:55:35.833704Z][Switch2:sw2-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:55:35.833476Z][Switch4:sw4-eth0] Received frame: SrcMac=00:00:00:00:03:03, DstMac=01:80:C2:00:00:00, Cable=Cable(sw3-eth2 <-> sw4-eth0)
[2026-03-01T09:55:35.833729Z][Switch4:sw4-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:55:35.833433Z][Cable(sw2-eth1 <-> sw3-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:02:02, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:35.833459Z][Host1] Received frame (SrcMac=00:00:00:00:01:03, DstMac=01:80:C2:00:00:00, Name=, NeedReply=false) not for this host (DstMac=aa:aa:aa:aa:aa:01); ignoring.
[2026-03-01T09:55:35.833502Z][Switch1:sw1-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:55:35.933885Z][Switch3:sw3-eth1] Received frame: SrcMac=00:00:00:00:01:02, DstMac=01:80:C2:00:00:00, Cable=Cable(sw1-eth1 <-> sw3-eth1)
[2026-03-01T09:55:35.934047Z][Switch3:sw3-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=0, LocalSwitchId=00:00:00:00:01:01
[2026-03-01T09:55:35.934078Z][Switch3:sw3-eth2] Received frame: SrcMac=00:00:00:00:04:01, DstMac=01:80:C2:00:00:00, Cable=Cable(sw3-eth2 <-> sw4-eth0)
[2026-03-01T09:55:35.934093Z][Switch3:sw3-eth2] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=2, LocalSwitchId=00:00:00:00:04:01
[2026-03-01T09:55:35.934135Z][Switch1:sw1-eth0] Received frame: SrcMac=00:00:00:00:02:01, DstMac=01:80:C2:00:00:00, Cable=Cable(sw1-eth0 <-> sw2-eth0)
[2026-03-01T09:55:35.934169Z][Switch1:sw1-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:02:01
[2026-03-01T09:55:35.934188Z][Switch3:sw3-eth0] Received frame: SrcMac=00:00:00:00:02:02, DstMac=01:80:C2:00:00:00, Cable=Cable(sw2-eth1 <-> sw3-eth0)
[2026-03-01T09:55:35.934369Z][Switch3:sw3-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:02:01
[2026-03-01T09:55:37.733203Z][Cable(sw3-eth2 <-> sw4-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:04:01, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:37.73321Z][Cable(sw1-eth0 <-> sw2-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:02:01, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:37.733243Z][Cable(sw2-eth1 <-> sw3-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:03:01, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:37.733262Z][Cable(sw1-eth1 <-> sw3-eth1)] Cable transmitting frame: SrcMac=00:00:00:00:03:02, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:37.733277Z][Cable(host1-eth0 <-> sw1-eth2)] Cable transmitting frame: SrcMac=00:00:00:00:01:03, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:37.7333Z][Cable(host2-eth0 <-> sw4-eth1)] Cable transmitting frame: SrcMac=00:00:00:00:04:02, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:37.833908Z][Host2] Received frame (SrcMac=00:00:00:00:04:02, DstMac=01:80:C2:00:00:00, Name=, NeedReply=false) not for this host (DstMac=aa:aa:aa:aa:aa:02); ignoring.
[2026-03-01T09:55:37.834062Z][Cable(sw3-eth2 <-> sw4-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:03:03, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:37.834085Z][Switch3:sw3-eth2] Received frame: SrcMac=00:00:00:00:04:01, DstMac=01:80:C2:00:00:00, Cable=Cable(sw3-eth2 <-> sw4-eth0)
[2026-03-01T09:55:37.834088Z][Cable(sw2-eth1 <-> sw3-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:02:02, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:37.834105Z][Switch3:sw3-eth2] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=2, LocalSwitchId=00:00:00:00:04:01
[2026-03-01T09:55:37.834121Z][Cable(sw1-eth1 <-> sw3-eth1)] Cable transmitting frame: SrcMac=00:00:00:00:01:02, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:37.834131Z][Switch1:sw1-eth1] Received frame: SrcMac=00:00:00:00:03:02, DstMac=01:80:C2:00:00:00, Cable=Cable(sw1-eth1 <-> sw3-eth1)
[2026-03-01T09:55:37.834664Z][Switch1:sw1-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:55:37.834156Z][Cable(sw1-eth0 <-> sw2-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:01:01, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:37.834165Z][Switch1:sw1-eth0] Received frame: SrcMac=00:00:00:00:02:01, DstMac=01:80:C2:00:00:00, Cable=Cable(sw1-eth0 <-> sw2-eth0)
[2026-03-01T09:55:37.834698Z][Switch1:sw1-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:02:01
[2026-03-01T09:55:37.834147Z][Host1] Received frame (SrcMac=00:00:00:00:01:03, DstMac=01:80:C2:00:00:00, Name=, NeedReply=false) not for this host (DstMac=aa:aa:aa:aa:aa:01); ignoring.
[2026-03-01T09:55:37.834172Z][Switch2:sw2-eth1] Received frame: SrcMac=00:00:00:00:03:01, DstMac=01:80:C2:00:00:00, Cable=Cable(sw2-eth1 <-> sw3-eth0)
[2026-03-01T09:55:37.834724Z][Switch2:sw2-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:55:37.93524Z][Switch3:sw3-eth0] Received frame: SrcMac=00:00:00:00:02:02, DstMac=01:80:C2:00:00:00, Cable=Cable(sw2-eth1 <-> sw3-eth0)
[2026-03-01T09:55:37.935366Z][Switch3:sw3-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:02:01
[2026-03-01T09:55:37.935271Z][Switch3:sw3-eth1] Received frame: SrcMac=00:00:00:00:01:02, DstMac=01:80:C2:00:00:00, Cable=Cable(sw1-eth1 <-> sw3-eth1)
[2026-03-01T09:55:37.935397Z][Switch3:sw3-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=0, LocalSwitchId=00:00:00:00:01:01
[2026-03-01T09:55:37.935219Z][Switch2:sw2-eth0] Received frame: SrcMac=00:00:00:00:01:01, DstMac=01:80:C2:00:00:00, Cable=Cable(sw1-eth0 <-> sw2-eth0)
[2026-03-01T09:55:37.935418Z][Switch2:sw2-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=0, LocalSwitchId=00:00:00:00:01:01
[2026-03-01T09:55:37.935294Z][Switch4:sw4-eth0] Received frame: SrcMac=00:00:00:00:03:03, DstMac=01:80:C2:00:00:00, Cable=Cable(sw3-eth2 <-> sw4-eth0)
[2026-03-01T09:55:37.935483Z][Switch4:sw4-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:55:39.733635Z][Cable(host2-eth0 <-> sw4-eth1)] Cable transmitting frame: SrcMac=00:00:00:00:04:02, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:39.733639Z][Cable(sw2-eth1 <-> sw3-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:02:02, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:39.733676Z][Cable(sw1-eth0 <-> sw2-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:02:01, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:39.733694Z][Cable(sw1-eth1 <-> sw3-eth1)] Cable transmitting frame: SrcMac=00:00:00:00:03:02, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:39.733706Z][Cable(sw3-eth2 <-> sw4-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:03:03, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:39.733718Z][Cable(host1-eth0 <-> sw1-eth2)] Cable transmitting frame: SrcMac=00:00:00:00:01:03, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:39.834887Z][Cable(sw1-eth0 <-> sw2-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:01:01, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:39.834886Z][Host1] Received frame (SrcMac=00:00:00:00:01:03, DstMac=01:80:C2:00:00:00, Name=, NeedReply=false) not for this host (DstMac=aa:aa:aa:aa:aa:01); ignoring.
[2026-03-01T09:55:39.834929Z][Cable(sw1-eth1 <-> sw3-eth1)] Cable transmitting frame: SrcMac=00:00:00:00:01:02, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:39.834941Z][Switch1:sw1-eth1] Received frame: SrcMac=00:00:00:00:03:02, DstMac=01:80:C2:00:00:00, Cable=Cable(sw1-eth1 <-> sw3-eth1)
[2026-03-01T09:55:39.835073Z][Switch1:sw1-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:55:39.83496Z][Switch4:sw4-eth0] Received frame: SrcMac=00:00:00:00:03:03, DstMac=01:80:C2:00:00:00, Cable=Cable(sw3-eth2 <-> sw4-eth0)
[2026-03-01T09:55:39.835102Z][Switch4:sw4-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:55:39.834981Z][Cable(sw2-eth1 <-> sw3-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:03:01, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:39.834987Z][Switch3:sw3-eth0] Received frame: SrcMac=00:00:00:00:02:02, DstMac=01:80:C2:00:00:00, Cable=Cable(sw2-eth1 <-> sw3-eth0)
[2026-03-01T09:55:39.835191Z][Switch3:sw3-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:02:01
[2026-03-01T09:55:39.834951Z][Cable(sw3-eth2 <-> sw4-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:04:01, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:39.834973Z][Host2] Received frame (SrcMac=00:00:00:00:04:02, DstMac=01:80:C2:00:00:00, Name=, NeedReply=false) not for this host (DstMac=aa:aa:aa:aa:aa:02); ignoring.
[2026-03-01T09:55:39.834997Z][Switch1:sw1-eth0] Received frame: SrcMac=00:00:00:00:02:01, DstMac=01:80:C2:00:00:00, Cable=Cable(sw1-eth0 <-> sw2-eth0)
[2026-03-01T09:55:39.835226Z][Switch1:sw1-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:02:01
[2026-03-01T09:55:39.936147Z][Switch3:sw3-eth1] Received frame: SrcMac=00:00:00:00:01:02, DstMac=01:80:C2:00:00:00, Cable=Cable(sw1-eth1 <-> sw3-eth1)
[2026-03-01T09:55:39.936272Z][Switch3:sw3-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=0, LocalSwitchId=00:00:00:00:01:01
[2026-03-01T09:55:39.936185Z][Switch2:sw2-eth1] Received frame: SrcMac=00:00:00:00:03:01, DstMac=01:80:C2:00:00:00, Cable=Cable(sw2-eth1 <-> sw3-eth0)
[2026-03-01T09:55:39.936311Z][Switch2:sw2-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:55:39.936136Z][Switch3:sw3-eth2] Received frame: SrcMac=00:00:00:00:04:01, DstMac=01:80:C2:00:00:00, Cable=Cable(sw3-eth2 <-> sw4-eth0)
[2026-03-01T09:55:39.936333Z][Switch3:sw3-eth2] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=2, LocalSwitchId=00:00:00:00:04:01
[2026-03-01T09:55:39.9362Z][Switch2:sw2-eth0] Received frame: SrcMac=00:00:00:00:01:01, DstMac=01:80:C2:00:00:00, Cable=Cable(sw1-eth0 <-> sw2-eth0)
[2026-03-01T09:55:39.93635Z][Switch2:sw2-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=0, LocalSwitchId=00:00:00:00:01:01
[2026-03-01T09:55:41.734779Z][Cable(host2-eth0 <-> sw4-eth1)] Cable transmitting frame: SrcMac=00:00:00:00:04:02, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:41.73477Z][Cable(sw3-eth2 <-> sw4-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:04:01, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:41.734808Z][Cable(sw2-eth1 <-> sw3-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:02:02, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:41.734822Z][Cable(sw1-eth0 <-> sw2-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:02:01, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:41.734837Z][Cable(sw1-eth1 <-> sw3-eth1)] Cable transmitting frame: SrcMac=00:00:00:00:03:02, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:41.734852Z][Cable(host1-eth0 <-> sw1-eth2)] Cable transmitting frame: SrcMac=00:00:00:00:01:03, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:41.836034Z][Cable(sw2-eth1 <-> sw3-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:03:01, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:41.836028Z][Host1] Received frame (SrcMac=00:00:00:00:01:03, DstMac=01:80:C2:00:00:00, Name=, NeedReply=false) not for this host (DstMac=aa:aa:aa:aa:aa:01); ignoring.
[2026-03-01T09:55:41.836067Z][Cable(sw1-eth0 <-> sw2-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:01:01, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:41.836076Z][Switch1:sw1-eth0] Received frame: SrcMac=00:00:00:00:02:01, DstMac=01:80:C2:00:00:00, Cable=Cable(sw1-eth0 <-> sw2-eth0)
[2026-03-01T09:55:41.836204Z][Switch1:sw1-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:02:01
[2026-03-01T09:55:41.836096Z][Switch1:sw1-eth1] Received frame: SrcMac=00:00:00:00:03:02, DstMac=01:80:C2:00:00:00, Cable=Cable(sw1-eth1 <-> sw3-eth1)
[2026-03-01T09:55:41.836233Z][Switch1:sw1-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:55:41.836117Z][Cable(sw3-eth2 <-> sw4-eth0)] Cable transmitting frame: SrcMac=00:00:00:00:03:03, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:41.836125Z][Switch3:sw3-eth2] Received frame: SrcMac=00:00:00:00:04:01, DstMac=01:80:C2:00:00:00, Cable=Cable(sw3-eth2 <-> sw4-eth0)
[2026-03-01T09:55:41.836841Z][Switch3:sw3-eth2] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=2, LocalSwitchId=00:00:00:00:04:01
[2026-03-01T09:55:41.836089Z][Cable(sw1-eth1 <-> sw3-eth1)] Cable transmitting frame: SrcMac=00:00:00:00:01:02, DstMac=01:80:C2:00:00:00 with 100ms delay
[2026-03-01T09:55:41.83611Z][Host2] Received frame (SrcMac=00:00:00:00:04:02, DstMac=01:80:C2:00:00:00, Name=, NeedReply=false) not for this host (DstMac=aa:aa:aa:aa:aa:02); ignoring.
[2026-03-01T09:55:41.836135Z][Switch3:sw3-eth0] Received frame: SrcMac=00:00:00:00:02:02, DstMac=01:80:C2:00:00:00, Cable=Cable(sw2-eth1 <-> sw3-eth0)
[2026-03-01T09:55:41.836877Z][Switch3:sw3-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:02:01
[2026-03-01T09:55:41.937281Z][Switch3:sw3-eth1] Received frame: SrcMac=00:00:00:00:01:02, DstMac=01:80:C2:00:00:00, Cable=Cable(sw1-eth1 <-> sw3-eth1)
[2026-03-01T09:55:41.937408Z][Switch3:sw3-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=0, LocalSwitchId=00:00:00:00:01:01
[2026-03-01T09:55:41.937438Z][Switch2:sw2-eth1] Received frame: SrcMac=00:00:00:00:03:01, DstMac=01:80:C2:00:00:00, Cable=Cable(sw2-eth1 <-> sw3-eth0)
[2026-03-01T09:55:41.937456Z][Switch2:sw2-eth1] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:55:41.93747Z][Switch2:sw2-eth0] Received frame: SrcMac=00:00:00:00:01:01, DstMac=01:80:C2:00:00:00, Cable=Cable(sw1-eth0 <-> sw2-eth0)
[2026-03-01T09:55:41.937478Z][Switch2:sw2-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=0, LocalSwitchId=00:00:00:00:01:01
[2026-03-01T09:55:41.937493Z][Switch4:sw4-eth0] Received frame: SrcMac=00:00:00:00:03:03, DstMac=01:80:C2:00:00:00, Cable=Cable(sw3-eth2 <-> sw4-eth0)
[2026-03-01T09:55:41.941892Z][Switch4:sw4-eth0] Received BPDU: RootSwitchId=00:00:00:00:01:01, RootPathCost=1, LocalSwitchId=00:00:00:00:03:01
[2026-03-01T09:55:42.729224Z][Switch1] ------> I am the root switch, all ports set to DESIGNATED.
[2026-03-01T09:55:42.729506Z][Switch3:sw3-eth0] ------> Port set to BLOCKING (LocalRootPathCost=1, OtherSwitchRootPathCost=1, LocalSwitchId=00:00:00:00:03:01, OtherSwitchId=00:00:00:00:02:01)
[2026-03-01T09:55:42.729546Z][Switch3:sw3-eth1] ------> Port set to ROOT (RootSwitchId=00:00:00:00:01:01, RootPathCost=1)
[2026-03-01T09:55:42.729561Z][Switch3:sw3-eth2] ------> Port set to DESIGNATED (LocalRootPathCost=1, OtherSwitchRootPathCost=2, LocalSwitchId=00:00:00:00:03:01, OtherSwitchId=00:00:00:00:04:01)
[2026-03-01T09:55:42.729763Z][Switch2:sw2-eth0] ------> Port set to ROOT (RootSwitchId=00:00:00:00:01:01, RootPathCost=1)
[2026-03-01T09:55:42.729796Z][Switch2:sw2-eth1] ------> Port set to DESIGNATED (LocalRootPathCost=1, OtherSwitchRootPathCost=1, LocalSwitchId=00:00:00:00:02:01, OtherSwitchId=00:00:00:00:03:01)
[2026-03-01T09:55:42.729931Z][Switch4:sw4-eth0] ------> Port set to ROOT (RootSwitchId=00:00:00:00:01:01, RootPathCost=2)
[2026-03-01T09:55:42.729954Z][Switch4:sw4-eth1] ------> Port set to BLOCKING (LocalRootPathCost=2, OtherSwitchRootPathCost=0, LocalSwitchId=00:00:00:00:04:01, OtherSwitchId=)

===============================================
   STP CONVERGENCE RESULTS
===============================================

Switch: Switch1
  Bridge ID: 00:00:00:00:01:01
  Root Switch ID: 00:00:00:00:01:01
  Root Path Cost: 0
  STP State: 4
  Ports:
    Port 0 (sw1-eth0):
      Role: DESIGNATED
      Root Path Cost: 0
    Port 1 (sw1-eth1):
      Role: DESIGNATED
      Root Path Cost: 0
    Port 2 (sw1-eth2):
      Role: DESIGNATED (PortFast)
      Root Path Cost: 0

Switch: Switch2
  Bridge ID: 00:00:00:00:02:01
  Root Switch ID: 00:00:00:00:01:01
  Root Path Cost: 1
  STP State: 4
  Ports:
    Port 0 (sw2-eth0):
      Role: ROOT
      Root Path Cost: 1
    Port 1 (sw2-eth1):
      Role: DESIGNATED
      Root Path Cost: 2

Switch: Switch3
  Bridge ID: 00:00:00:00:03:01
  Root Switch ID: 00:00:00:00:01:01
  Root Path Cost: 1
  STP State: 4
  Ports:
    Port 0 (sw3-eth0):
      Role: BLOCKING
      Root Path Cost: 2
    Port 1 (sw3-eth1):
      Role: ROOT
      Root Path Cost: 1
    Port 2 (sw3-eth2):
      Role: DESIGNATED
      Root Path Cost: 3

Switch: Switch4
  Bridge ID: 00:00:00:00:04:01
  Root Switch ID: 00:00:00:00:01:01
  Root Path Cost: 2
  STP State: 4
  Ports:
    Port 0 (sw4-eth0):
      Role: ROOT
      Root Path Cost: 2
    Port 1 (sw4-eth1):
      Role: BLOCKING (PortFast)
      Root Path Cost: 0


===============================================
   STP BLOCKING DEMONSTRATION
===============================================

Host1 sending frame to Host2...
[2026-03-01T09:55:43.729686Z][Host1] Sending frame (SrcMac=aa:aa:aa:aa:aa:01, DstMac=aa:aa:aa:aa:aa:02, Name=Greeting Hi, NeedReply=true)
[2026-03-01T09:55:43.729712Z][Host1:host1-eth0] Sending frame: SrcMac=aa:aa:aa:aa:aa:01, DstMac=aa:aa:aa:aa:aa:02, Cable=Cable(host1-eth0 <-> sw1-eth2)
[2026-03-01T09:55:43.729759Z][Cable(host1-eth0 <-> sw1-eth2)] Cable transmitting frame: SrcMac=aa:aa:aa:aa:aa:01, DstMac=aa:aa:aa:aa:aa:02 with 100ms delay
[2026-03-01T09:55:43.830877Z][Switch1:sw1-eth2] Received frame: SrcMac=aa:aa:aa:aa:aa:01, DstMac=aa:aa:aa:aa:aa:02, Cable=Cable(host1-eth0 <-> sw1-eth2)
[2026-03-01T09:55:43.83099Z][Switch1:sw1-eth2] Learned MAC aa:aa:aa:aa:aa:01
[2026-03-01T09:55:43.830996Z][Switch1] Does not know destination MAC aa:aa:aa:aa:aa:02; broadcasting frame
[2026-03-01T09:55:43.831Z][Switch1:sw1-eth0] Sending frame: SrcMac=aa:aa:aa:aa:aa:01, DstMac=aa:aa:aa:aa:aa:02, Cable=Cable(sw1-eth0 <-> sw2-eth0)
[2026-03-01T09:55:43.831037Z][Switch1:sw1-eth1] Sending frame: SrcMac=aa:aa:aa:aa:aa:01, DstMac=aa:aa:aa:aa:aa:02, Cable=Cable(sw1-eth1 <-> sw3-eth1)
[2026-03-01T09:55:43.831044Z][Cable(sw1-eth1 <-> sw3-eth1)] Cable transmitting frame: SrcMac=aa:aa:aa:aa:aa:01, DstMac=aa:aa:aa:aa:aa:02 with 100ms delay
[2026-03-01T09:55:43.831054Z][Cable(sw1-eth0 <-> sw2-eth0)] Cable transmitting frame: SrcMac=aa:aa:aa:aa:aa:01, DstMac=aa:aa:aa:aa:aa:02 with 100ms delay
[2026-03-01T09:55:43.9321Z][Switch2:sw2-eth0] Received frame: SrcMac=aa:aa:aa:aa:aa:01, DstMac=aa:aa:aa:aa:aa:02, Cable=Cable(sw1-eth0 <-> sw2-eth0)
[2026-03-01T09:55:43.932158Z][Switch2:sw2-eth0] Learned MAC aa:aa:aa:aa:aa:01
[2026-03-01T09:55:43.93216Z][Switch2] Does not know destination MAC aa:aa:aa:aa:aa:02; broadcasting frame
[2026-03-01T09:55:43.932132Z][Switch3:sw3-eth1] Received frame: SrcMac=aa:aa:aa:aa:aa:01, DstMac=aa:aa:aa:aa:aa:02, Cable=Cable(sw1-eth1 <-> sw3-eth1)
[2026-03-01T09:55:43.932178Z][Switch3:sw3-eth1] Learned MAC aa:aa:aa:aa:aa:01
[2026-03-01T09:55:43.932182Z][Switch3] Does not know destination MAC aa:aa:aa:aa:aa:02; broadcasting frame
[2026-03-01T09:55:43.932186Z][Switch3:sw3-eth0] Port is in STP BLOCKING state, cannot send frame to aa:aa:aa:aa:aa:02
[2026-03-01T09:55:43.932164Z][Switch2:sw2-eth1] Sending frame: SrcMac=aa:aa:aa:aa:aa:01, DstMac=aa:aa:aa:aa:aa:02, Cable=Cable(sw2-eth1 <-> sw3-eth0)
[2026-03-01T09:55:43.932199Z][Cable(sw2-eth1 <-> sw3-eth0)] Cable transmitting frame: SrcMac=aa:aa:aa:aa:aa:01, DstMac=aa:aa:aa:aa:aa:02 with 100ms delay
[2026-03-01T09:55:43.93219Z][Switch3:sw3-eth2] Sending frame: SrcMac=aa:aa:aa:aa:aa:01, DstMac=aa:aa:aa:aa:aa:02, Cable=Cable(sw3-eth2 <-> sw4-eth0)
[2026-03-01T09:55:43.932214Z][Cable(sw3-eth2 <-> sw4-eth0)] Cable transmitting frame: SrcMac=aa:aa:aa:aa:aa:01, DstMac=aa:aa:aa:aa:aa:02 with 100ms delay
[2026-03-01T09:55:44.033295Z][Switch4:sw4-eth0] Received frame: SrcMac=aa:aa:aa:aa:aa:01, DstMac=aa:aa:aa:aa:aa:02, Cable=Cable(sw3-eth2 <-> sw4-eth0)
[2026-03-01T09:55:44.033395Z][Switch4:sw4-eth0] Learned MAC aa:aa:aa:aa:aa:01
[2026-03-01T09:55:44.033402Z][Switch4] Does not know destination MAC aa:aa:aa:aa:aa:02; broadcasting frame
[2026-03-01T09:55:44.03341Z][Switch4:sw4-eth1] Sending frame: SrcMac=aa:aa:aa:aa:aa:01, DstMac=aa:aa:aa:aa:aa:02, Cable=Cable(host2-eth0 <-> sw4-eth1)
[2026-03-01T09:55:44.033421Z][Cable(host2-eth0 <-> sw4-eth1)] Cable transmitting frame: SrcMac=aa:aa:aa:aa:aa:01, DstMac=aa:aa:aa:aa:aa:02 with 100ms delay
[2026-03-01T09:55:44.033314Z][Switch3:sw3-eth0] Received frame to aa:aa:aa:aa:aa:02 but port is in STP BLOCKING state, dropping frame
[2026-03-01T09:55:44.133875Z][Host2] Received frame (SrcMac=aa:aa:aa:aa:aa:01, DstMac=aa:aa:aa:aa:aa:02, Name=Greeting Hi, NeedReply=true)
[2026-03-01T09:55:44.133918Z][Host2] Sending frame (SrcMac=aa:aa:aa:aa:aa:02, DstMac=aa:aa:aa:aa:aa:01, Name=Reply to Greeting Hi, NeedReply=false)
[2026-03-01T09:55:44.133923Z][Host2:host2-eth0] Sending frame: SrcMac=aa:aa:aa:aa:aa:02, DstMac=aa:aa:aa:aa:aa:01, Cable=Cable(host2-eth0 <-> sw4-eth1)
[2026-03-01T09:55:44.133933Z][Cable(host2-eth0 <-> sw4-eth1)] Cable transmitting frame: SrcMac=aa:aa:aa:aa:aa:02, DstMac=aa:aa:aa:aa:aa:01 with 100ms delay
[2026-03-01T09:55:44.235129Z][Switch4:sw4-eth1] Received frame: SrcMac=aa:aa:aa:aa:aa:02, DstMac=aa:aa:aa:aa:aa:01, Cable=Cable(host2-eth0 <-> sw4-eth1)
[2026-03-01T09:55:44.235391Z][Switch4:sw4-eth1] Learned MAC aa:aa:aa:aa:aa:02
[2026-03-01T09:55:44.235413Z][Switch4:sw4-eth0] Sending frame: SrcMac=aa:aa:aa:aa:aa:02, DstMac=aa:aa:aa:aa:aa:01, Cable=Cable(sw3-eth2 <-> sw4-eth0)
[2026-03-01T09:55:44.235431Z][Cable(sw3-eth2 <-> sw4-eth0)] Cable transmitting frame: SrcMac=aa:aa:aa:aa:aa:02, DstMac=aa:aa:aa:aa:aa:01 with 100ms delay
[2026-03-01T09:55:44.336591Z][Switch3:sw3-eth2] Received frame: SrcMac=aa:aa:aa:aa:aa:02, DstMac=aa:aa:aa:aa:aa:01, Cable=Cable(sw3-eth2 <-> sw4-eth0)
[2026-03-01T09:55:44.336753Z][Switch3:sw3-eth2] Learned MAC aa:aa:aa:aa:aa:02
[2026-03-01T09:55:44.336771Z][Switch3:sw3-eth1] Sending frame: SrcMac=aa:aa:aa:aa:aa:02, DstMac=aa:aa:aa:aa:aa:01, Cable=Cable(sw1-eth1 <-> sw3-eth1)
[2026-03-01T09:55:44.336794Z][Cable(sw1-eth1 <-> sw3-eth1)] Cable transmitting frame: SrcMac=aa:aa:aa:aa:aa:02, DstMac=aa:aa:aa:aa:aa:01 with 100ms delay
[2026-03-01T09:55:44.437995Z][Switch1:sw1-eth1] Received frame: SrcMac=aa:aa:aa:aa:aa:02, DstMac=aa:aa:aa:aa:aa:01, Cable=Cable(sw1-eth1 <-> sw3-eth1)
[2026-03-01T09:55:44.438122Z][Switch1:sw1-eth1] Learned MAC aa:aa:aa:aa:aa:02
[2026-03-01T09:55:44.438138Z][Switch1:sw1-eth2] Sending frame: SrcMac=aa:aa:aa:aa:aa:02, DstMac=aa:aa:aa:aa:aa:01, Cable=Cable(host1-eth0 <-> sw1-eth2)
[2026-03-01T09:55:44.438155Z][Cable(host1-eth0 <-> sw1-eth2)] Cable transmitting frame: SrcMac=aa:aa:aa:aa:aa:02, DstMac=aa:aa:aa:aa:aa:01 with 100ms delay
[2026-03-01T09:55:44.539245Z][Host1] Received frame (SrcMac=aa:aa:aa:aa:aa:02, DstMac=aa:aa:aa:aa:aa:01, Name=Reply to Greeting Hi, NeedReply=false)

===============================================
   SIMULATION COMPLETE
===============================================

Cable between sw3-eth2 and sw4-eth0 shutting down
[Switch Switch1] Shutting down goroutine
[Switch Switch2] Shutting down goroutine
[Switch Switch3] Shutting down goroutine
[Switch Switch4] Shutting down goroutine
[Host Host1] Shutting down goroutine
[Host Host2] Shutting down goroutine
Cable between host1-eth0 and sw1-eth2 shutting down
Cable between host2-eth0 and sw4-eth1 shutting down
Cable between sw1-eth0 and sw2-eth0 shutting down
Cable between sw2-eth1 and sw3-eth0 shutting down
Cable between sw1-eth1 and sw3-eth1 shutting down
```
</details>

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

```
</details>