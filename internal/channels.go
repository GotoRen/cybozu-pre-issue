package internal

import (
	"sync"
)

type outboundQueue struct {
	c  chan *QueueOutboundElement
	wg sync.WaitGroup
}

func newOutboundQueue() *outboundQueue {
	q := &outboundQueue{
		c: make(chan *QueueOutboundElement, QueueOutboundSize),
	}
	q.wg.Add(1)

	go func() {
		q.wg.Wait()
		close(q.c)
	}()

	return q
}

type inboundQueue struct {
	c  chan *QueueInboundElement
	wg sync.WaitGroup
}

func newInboundQueue() *inboundQueue {
	q := &inboundQueue{
		c: make(chan *QueueInboundElement, QueueInboundSize),
	}
	q.wg.Add(1)

	go func() {
		q.wg.Wait()
		close(q.c)
	}()

	return q
}

type QueueInboundElement struct {
	sync.Mutex
	buffer   *[MaxMessageSize]byte
	data     []byte
	checksum [32]byte
}

type QueueOutboundElement struct {
	sync.Mutex
	buffer   *[MaxMessageSize]byte
	data     []byte
	checksum [32]byte
}

// PopulatePools creates new buffers.
func (raw *Raw) PopulatePools() {
	raw.pool.messageBuffers = NewWaitPool(PreallocatedBuffersPerPool, func() interface{} {
		return new([MaxMessageSize]byte)
	})
	raw.pool.inboundElements = NewWaitPool(PreallocatedBuffersPerPool, func() interface{} {
		return new(QueueInboundElement)
	})
	raw.pool.outboundElements = NewWaitPool(PreallocatedBuffersPerPool, func() interface{} {
		return new(QueueOutboundElement)
	})
}

// NewWaitPool initializes wait pool.
func NewWaitPool(max uint32, newFunc func() interface{}) *WaitPool {
	p := &WaitPool{pool: sync.Pool{New: newFunc}, max: max}
	p.cond = sync.Cond{L: &p.lock}

	return p
}
