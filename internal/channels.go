package internal

import (
	"runtime"
	"sync"
)

type Data struct {
	Text   string
	Buffer []byte
	sync.Mutex
}

type Element struct {
	Inbound  chan *Data
	Outbound chan *Data
	Wg       sync.WaitGroup
}

func (elem *Element) RoutineSHA256Converter() {
	// start workers
	cpus := runtime.NumCPU()

	elem.Wg.Add(cpus) // One for each RoutineSHA256Converter.

	for i := 0; i < cpus; i++ {
		go elem.Convert2SHA256()
	}
}

func (elem *Element) RoutineWriter() {
	elem.Wg.Add(1)
	go Write(elem.Outbound, &elem.Wg)
}
