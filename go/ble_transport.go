package main

import "fmt"

// BLEEvent représente un peer détecté via BLE
type BLEEvent struct {
	PeerID string
}

// BLECallback est la fonction à appeler quand un peer est détecté
var BLECallback func(event BLEEvent)

// SetBLECallback permet à Android de passer une fonction de callback
func SetBLECallback(cb func(event BLEEvent)) {
	BLECallback = cb
}

// AddPeerFromBLE est appelé depuis Android quand un peer est découvert
func (n *Node) AddPeerFromBLE(peerID string, addr string) {
	if _, exists := n.Peers[peerID]; !exists {
		n.Peers[peerID] = addr
		fmt.Println("Nouveau peer BLE ajouté:", peerID, addr)
		// Appeler callback Go → Android si nécessaire
		if BLECallback != nil {
			BLECallback(BLEEvent{PeerID: peerID})
		}
	}
}
