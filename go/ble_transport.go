package mesh

type BLEEvent struct {
	Type    string
	Data    []byte
	PeerID  string
	Message string
}

var BLECallback func(event BLEEvent)

func SetBLECallback(cb func(event BLEEvent)) {
	BLECallback = cb
}