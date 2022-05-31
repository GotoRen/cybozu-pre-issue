package internal

import (
	"runtime"
	"sync"
)

// Type of struct: InputText -> Stored in structure
type Raw struct {
	sync.Mutex
	Buffer []byte
}

// Type of struct: InboundChannels -> RoutineProcessing -> OutboundChannels
type Element struct {
	Wg       sync.WaitGroup
	Inbound  chan *Raw
	Outbound chan *Raw
}

// SHA256Converter calls RoutineConvert2SHA256 as a goroutine.
func (elem *Element) SHA256Converter() {
	// start worker threads.
	cpus := runtime.NumCPU()

	// one for each RoutineConvert2SHA256.
	elem.Wg.Add(cpus)

	for i := 0; i < cpus; i++ {
		go elem.RoutineConvert2SHA256()
	}
}

// Writer calls RoutineWriter as a goroutine.
func (elem *Element) Writer() {
	elem.Wg.Add(1)
	go elem.RoutineWrite()
}
