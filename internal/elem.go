package internal

import "sync"

// Data: InputText -> Stored in structure.
type Data struct {
	Text   []byte
	Buffer []byte
	Mu     sync.Mutex
}

// Element: InboundChannels -> RoutineProcessing -> OutboundChannels.
type Element struct {
	Inbound  chan *Data
	Outbound chan *Data
	Wg       sync.WaitGroup
}
