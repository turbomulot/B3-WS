package mesh

import (
	"context"
	"fmt"
	"time"

	libp2p "github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
)

type Node struct {
	Host     host.Host
	Peers    map[string]string
	MsgQueue []Message
}

func NewNode(ctx context.Context) (*Node, error) {
	host, err := libp2p.New()
	if err != nil {
		return nil, err
	}
	return &Node{
		Host:     host,
		Peers:    make(map[string]string),
		MsgQueue: []Message{},
	}, nil
}

func (n *Node) SendMessage(msg Message) {
	msg.HopCount++
	if msg.HopCount <= msg.TTL {
		n.MsgQueue = append(n.MsgQueue, msg)
	}
}

func (n *Node) SendMessageWithParams(to string, payload []byte, msgType int) {
	msg := Message{
		ID:        fmt.Sprintf("msg-%d", time.Now().UnixNano()),
		From:      n.Host.ID().String(),
		To:        to,
		Type:      MsgType(msgType),
		Payload:   payload,
		HopCount:  0,
		TTL:       5, // Default TTL
		Timestamp: time.Now().Unix(),
	}
	n.SendMessage(msg)
}

func (n *Node) ReceiveMessage(msg Message) {
	if msg.To != n.Host.ID().String() {
		msg.HopCount++
		if msg.HopCount <= msg.TTL {
			n.MsgQueue = append(n.MsgQueue, msg)
		}
	}
}

func (n *Node) ProcessQueue() {
	ticker := time.NewTicker(2 * time.Second)
	go func() {
		for range ticker.C {
			for _, msg := range n.MsgQueue {
				fmt.Println("Forwarding message:", msg)
				// TODO: envoyer à voisins via BLE/Wi-Fi Direct
			}
			n.MsgQueue = []Message{}
		}
	}()
}