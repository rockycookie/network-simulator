# network-simulator
Golang code to simulate networking devices which compose a network

## Local commands
```
% go run main.go
Switch: MySwitch
  Nic 1: ID=sw-eth0, Mac=00:11:22:33:44:01
  Nic 2: ID=sw-eth1, Mac=00:11:22:33:44:02
  Nic 3: ID=sw-eth2, Mac=00:11:22:33:44:03
  Nic 4: ID=sw-eth3, Mac=00:11:22:33:44:04

Host connections:
  Host1: Nic=host1-eth0, Mac=00:11:22:33:44:11
  Host2: Nic=host2-eth0, Mac=00:11:22:33:44:21

Cables:
  Cable1: host1-eth0 <-> sw-eth0
  Cable2: host2-eth0 <-> sw-eth1

*******************************
Starting frame transmission simulation
*******************************
[Host Host1] Started running
[Host Host2] Started running
[Host Host3] Started running
[Switch MySwitch] is running with 4 NICs

Host1 sending frame to Host2
Cable transmitting frame: SrcMac=00:11:22:33:44:11, DstMac=00:11:22:33:44:21, from NIC=host1-eth0, to NIC=sw-eth0 with 100ms delay
[2026-01-26T00:07:55.184815Z][Switch MySwitch] Switch received frame: SrcMac=00:11:22:33:44:11, DstMac=00:11:22:33:44:21, on NIC=sw-eth0
[2026-01-26T00:07:55.185021Z][Switch MySwitch] Switch does not know destination MAC 00:11:22:33:44:21; broadcasting frame
[2026-01-26T00:07:55.185032Z][Switch MySwitch] Switch sending frame: SrcMac=00:11:22:33:44:11, DstMac=00:11:22:33:44:21, via NIC=sw-eth1
[2026-01-26T00:07:55.185042Z][Switch MySwitch] Switch sending frame: SrcMac=00:11:22:33:44:11, DstMac=00:11:22:33:44:21, via NIC=sw-eth2
Cable transmitting frame: SrcMac=00:11:22:33:44:11, DstMac=00:11:22:33:44:21, from NIC=sw-eth2, to NIC=host3-eth0 with 100ms delay
Cable transmitting frame: SrcMac=00:11:22:33:44:11, DstMac=00:11:22:33:44:21, from NIC=sw-eth1, to NIC=host2-eth0 with 100ms delay
[2026-01-26T00:07:55.286336Z][Host Host3] Frame (SrcMac=00:11:22:33:44:11, DstMac=00:11:22:33:44:21) not for this host (DstMac=00:11:22:33:44:31); ignoring.
[2026-01-26T00:07:55.286432Z][Host Host2] Frame (SrcMac=00:11:22:33:44:11, DstMac=00:11:22:33:44:21) requests reply; sending reply frame.
Cable transmitting frame: SrcMac=00:11:22:33:44:21, DstMac=00:11:22:33:44:11, from NIC=host2-eth0, to NIC=sw-eth1 with 100ms delay
[2026-01-26T00:07:55.387798Z][Switch MySwitch] Switch received frame: SrcMac=00:11:22:33:44:21, DstMac=00:11:22:33:44:11, on NIC=sw-eth1
[2026-01-26T00:07:55.388036Z][Switch MySwitch] Switch sending frame: SrcMac=00:11:22:33:44:21, DstMac=00:11:22:33:44:11, via NIC=sw-eth0
Cable transmitting frame: SrcMac=00:11:22:33:44:21, DstMac=00:11:22:33:44:11, from NIC=sw-eth0, to NIC=host1-eth0 with 100ms delay
[2026-01-26T00:07:55.489169Z][Host Host1] Received NoReply frame (SrcMac=00:11:22:33:44:21, DstMac=00:11:22:33:44:11): Reply to Greeting

Simulation ended.
Cable between host3-eth0 and sw-eth2 shutting down
[Host Host1] Shutting down goroutine
[Host Host2] Shutting down goroutine
[Host Host3] Shutting down goroutine
[Switch MySwitch] Shutting down goroutine
Cable between host1-eth0 and sw-eth0 shutting down
Cable between host2-eth0 and sw-eth1 shutting down
```
