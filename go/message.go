package mesh

type MsgType int

const (
	Text MsgType = iota
	Image
	Audio
)

type Message struct {
	ID        string
	From      string
	To        string
	Type      MsgType
	Payload   []byte
	HopCount  int
	TTL       int
	Timestamp int64
}