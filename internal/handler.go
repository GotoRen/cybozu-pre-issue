package internal

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"runtime"

	"github.com/GotoRen/cyboze_pre_issue/internal/logger"
)

type Raw struct {
	pool struct {
		messageBuffers   *WaitPool
		inboundElements  *WaitPool
		outboundElements *WaitPool
	}

	queue struct {
		inbound  *inboundQueue
		outbound *outboundQueue
	}
}

func (raw *Raw) RoutineSHA256Converter() {
	// defer raw.queue.Done()

	for elem := range raw.queue.inbound.c {
		elem.checksum = sha256.Sum256(elem.data)
		elem.Unlock()
	}
}

// func (raw *Raw) RoutineSequentialWriter() {
// 	for elme := range raw.queue.outbound.c {
// 		elme.Lock()
// 		// 書き出し
// 		fmt.Println(hex.Dump(elme.checksum[:]))
// 		elme.Unlock()
// 	}
// }

// NewRoutine creates a new routine.
func (raw *Raw) NewRoutine() {
	cpus := runtime.NumCPU() // start worker threads.

	// popularate
	raw.PopulatePools()

	// create queues
	raw.queue.inbound = newInboundQueue()
	raw.queue.outbound = newOutboundQueue()

	// auto draining
	// raw.queue.inbound = newAutodrainingInboundQueue(raw)
	// raw.queue.outbound = newAutodrainingOutboundQueue(raw)

	raw.queue.inbound.wg.Add(cpus) // One for each RoutineSHA256Converter.

	for i := 0; i < cpus; i++ {
		go raw.RoutineSHA256Converter() // convertion routine.
	}

	raw.queue.inbound.wg.Add(1)
	// go raw.RoutineSequentialWriter() // Writeルーチン
}

func (raw *Raw) RoutineSequentialReader(fp *os.File) {
	defer raw.queue.inbound.wg.Done()
	var elem *QueueInboundElement

	for {
		if elem != nil {
			// 削除
			raw.PutMessageBuffer(elem.buffer) // PutMessageBuffer puts 65535bytes.
			raw.PutInboundElement(elem)
		}
		elem = raw.NewInboundElement()
		size, err := fp.Read(elem.buffer[:])

		if err != nil {
			logger.LogErr("Failed to read line", "error", err)

			// 全て削除
			raw.PutMessageBuffer(elem.buffer) // PutMessageBuffer puts 65535bytes.
			raw.PutInboundElement(elem)
			return
		}

		if size == 0 {
			logger.LogErr("Reading size is too small", "error", size)
			continue
		}

		// elem.Lock()

		if elem != nil {
			raw.StagePacket(elem)
			elem = nil
			raw.SendStagedPackets()
		}
	}

	// close(elem.Inbound)
	// close(elem.Outbound)

	// elem.Wg.Wait() // main goroutine waits for other goroutines that have been added.
}

// StagePacket stating peer queue from outboundElement.
func (raw *Raw) StagePacket(elem *QueueInboundElement) {
	for {
		select {
		case raw.queue.inbound.c <- elem:
			return
		default:
		}
		select {
		// case tooOld := <-peer.queue.staged:
		// 	peer.device.PutMessageBuffer(tooOld.buffer)
		// 	peer.device.PutOutboundElement(tooOld)
		default:
		}
	}
}

func (raw *Raw) SendStagedPackets() {
	for {
		select {
		case elem := <-raw.queue.outbound.c:
			elem.Lock()

			if elem != nil {
				// 書き出し
				fmt.Println(hex.Dump(elem.checksum[:]))
			} else {
				// 消去
				raw.PutMessageBuffer(elem.buffer)
				raw.PutOutboundElement(elem)
			}

			elem.Unlock()
		default:
			return
		}
	}
}

func (raw *Raw) PutOutboundElement(elem *QueueOutboundElement) {
	elem.clearPointers()
	raw.pool.outboundElements.Put(elem)
}
