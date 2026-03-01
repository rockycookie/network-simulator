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
        - STP support
            - multi-switches root election
            - loop-free function by blocking certain switch port(s)
            - port-fast to ensure enpoint host connection to the switches
    - Host Ethernet features
        - Send and reply L2 frames using its unique NIC
            - Discard frames if their destination MAC does not match the host's NIC
    - Cable Ethernet features
        - Send frame with 100ms delay (to showcase which events are happening in parallel vs in sequential order)

## Notes
- L2 Ethernet Frame currently has flag `NeedReply` and attribute `Name` just to make life easier, as L3 logic is not currently imlemented yet

## Sample Setup & Run Results

- [Host-Switch-Host Communication](./docs/setup/host-to-host-via-switch.md)
- [4 Switches STP Root Election](./docs/setup/4-switches-stp-root-election.md)
- [Host-LoopedSwitches-Host Communication](./docs/setup/host-to-host-via-4-looped-switches.md)
