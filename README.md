# network-simulator
Golang code to simulate networking devices which compose a network

## Local commands
```
% go run main.go

Switch: Switch1
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

Host1 is sending a frame to Host2

Cable transmitting frame: SrcMac=00:11:22:33:44:11, DstMac=00:11:22:33:44:21, from NIC=host1-eth0, to NIC=sw-eth0
Switch 'Switch1' received frame: SrcMac=00:11:22:33:44:11, DstMac=00:11:22:33:44:21, on NIC=sw-eth0
Switch 'Switch1' does not know destination MAC 00:11:22:33:44:21; broadcasting frame
Switch 'Switch1' sending frame: SrcMac=00:11:22:33:44:11, DstMac=00:11:22:33:44:21, via NIC=sw-eth1
Cable transmitting frame: SrcMac=00:11:22:33:44:11, DstMac=00:11:22:33:44:21, from NIC=sw-eth1, to NIC=host2-eth0
[2026-01-25T23:00:45Z][Host Host2] Received frame: SrcMac=00:11:22:33:44:11, DstMac=00:11:22:33:44:21
Cable transmitting frame: SrcMac=00:11:22:33:44:21, DstMac=00:11:22:33:44:11, from NIC=host2-eth0, to NIC=sw-eth1
Switch 'Switch1' received frame: SrcMac=00:11:22:33:44:21, DstMac=00:11:22:33:44:11, on NIC=sw-eth1
Switch 'Switch1' sending frame: SrcMac=00:11:22:33:44:21, DstMac=00:11:22:33:44:11, via NIC=sw-eth0
Cable transmitting frame: SrcMac=00:11:22:33:44:21, DstMac=00:11:22:33:44:11, from NIC=sw-eth0, to NIC=host1-eth0
[2026-01-25T23:00:45Z][Host Host1] Received frame: SrcMac=00:11:22:33:44:21, DstMac=00:11:22:33:44:11
Switch 'Switch1' sending frame: SrcMac=00:11:22:33:44:11, DstMac=00:11:22:33:44:21, via NIC=sw-eth2
Cable transmitting frame: SrcMac=00:11:22:33:44:11, DstMac=00:11:22:33:44:21, from NIC=sw-eth2, to NIC=host3-eth0
[2026-01-25T23:00:45Z][Host Host3] Received frame: SrcMac=00:11:22:33:44:11, DstMac=00:11:22:33:44:21
[2026-01-25T23:00:45Z][Host Host3] Frame not for this host (DstMac=00:11:22:33:44:21); ignoring.
```
