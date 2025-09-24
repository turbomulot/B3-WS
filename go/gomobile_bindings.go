package mesh

import (
    "context"
)

// InitNode crée un nouveau nœud P2P
func InitNode() *Node {
    ctx := context.Background()
    node, _ := NewNode(ctx)
    go node.ProcessQueue()
    return node
}

// Wrapper pour envoyer un message
func (n *Node) SendMessageWrapper(to string, payload []byte, typ int) {
    n.SendMessageWithParams(to, payload, typ)
}

// Récupère la liste des pairs connus
func (n *Node) GetPeerList() []string {
    peers := []string{}
    for id := range n.Peers {
        peers = append(peers, id)
    }
    return peers
}
