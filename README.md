# network-simulator
Golang code to simulate networking devices which compose a network

## Local commands
```
go run main.go

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
```
