package internal

import (
	"runtime"
	"sync"
)

// Type of struct: InputText -> Stored in structure
type Data struct {
	Text   []byte
	Buffer []byte
	sync.Mutex
}

// Type of struct: InboundChannels -> RoutineProcessing -> OutboundChannels
type Element struct {
	Inbound  chan *Data
	Outbound chan *Data
	Wg       sync.WaitGroup
}

// SHA256Converter calls RoutineConvert2SHA256 as a goroutine.
func (elem *Element) SHA256Converter() {
	// start workers
	cpus := runtime.NumCPU()

	// One for each RoutineConvert2SHA256.
	elem.Wg.Add(cpus)

	for i := 0; i < cpus; i++ {
		go elem.RoutineConvert2SHA256()
	}
}

// Writer calls RoutineWriter as a goroutine.
func (elem *Element) Writer() {
	elem.Wg.Add(1)
	go elem.RoutineWriter()
}

// PutData places the received data.
func PutData(b []byte) (datum Data) {
	datum = Data{
		Text:   b,
		Buffer: nil,
	}
	return
}
