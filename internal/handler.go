package internal

import (
	"runtime"
	"sync"
)

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

// PutData places the received data.
func PutData(b []byte) *Data {
	return &Data{
		Text:   b,
		Buffer: nil,
		Mu:     sync.Mutex{},
	}
}
