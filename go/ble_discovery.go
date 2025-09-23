package main

import (
	"fmt"
)

// Advertising the node via BLE
func (n *Node) StartBLEAdvertising() {
	fmt.Println("Starting BLE advertising (pseudo-code)")
	// TODO: use a Go BLE library to advertise PeerID
}

// Scan for nearby nodes
func (n *Node) ScanBLENeighbors() {
	fmt.Println("Scanning BLE neighbors (pseudo-code)")
	// TODO: detect peers and add to n.Peers map
}
