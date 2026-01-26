# network-simulator
Golang code to simulate networking devices which compose a network

## Feature Supported
- Devices
    - Host
        - each with one Network Interface Card (NIC)
    - Cable
        - connects 2 NICs
    - Switch
        - each with an arry of NICs
- L2 Ethernet features
    - Switch
        - Broadcast for L2 flooding request
        - MAC address table caches mapping of MAC to switch ports (NICs)
        - Broadcast for Ethernet frames whose destination MAC does not exist in the MAC address table
    - Host Ethernet features
        - Send and reply L2 frames using its unique NIC
            - Discard frames if their destination MAC does not match the host's NIC
    - Cable Ethernet features
        - Send frame with 100ms delay (to showcase which events are happening in parallel vs in sequential order)

## Notes
- L2 Ethernet Frame currently has flag `NeedReply` and attribute `Name` just to make life easier, as L3 logic is not currently imlemented yet

## Sample Run

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
