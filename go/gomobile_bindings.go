package main

import "C"
import (
	"context"
	"fmt"
	"time"
)

//export InitNode
func InitNode() *Node {
	ctx := context.Background()
	node, _ := NewNode(ctx)
	node.ProcessQueue()
	node.StartBLEAdvertising()
	node.ScanBLENeighbors()
	node.StartWiFiDirect()
	return node
}

//export Send
func Send(node *Node, to string, payload []byte, typ int) {
	msg := Message{
		ID:        fmt.Sprintf("%d", time.Now().UnixNano()),
		From:      node.Host.ID().String(),
		To:        to,
		Type:      MsgType(typ),
		Payload:   payload,
		HopCount:  0,
		TTL:       5,
		Timestamp: time.Now().Unix(),
	}
	node.SendMessage(msg)
}
