package internal

import (
	"sync"
	"sync/atomic"
)

// Bit16 is bit length, QueueSize defines the size of queue.
const (
	Bit16                      = 16
	QueueInboundSize           = 1024
	QueueOutboundSize          = 1024
	PreallocatedBuffersPerPool = 0 // Disable and allow for infinite memory growth
)

// MaxMessageSize is 65535=(2^16)-1.
const MaxMessageSize = (1 << Bit16) - 1

// WaitPool pools wait pooling.
type WaitPool struct {
	pool  sync.Pool
	cond  sync.Cond
	lock  sync.Mutex
	count uint32
	max   uint32
}

// PutMessageBuffer puts 65535bytes.
func (raw *Raw) PutMessageBuffer(msg *[MaxMessageSize]byte) {
	raw.pool.messageBuffers.Put(msg)
}

// Put adds x to the pool.
func (p *WaitPool) Put(x interface{}) {
	p.pool.Put(x)

	if p.max == 0 {
		return
	}

	atomic.AddUint32(&p.count, ^uint32(0))
	p.cond.Signal()
}

func (raw *Raw) PutInboundElement(elem *QueueInboundElement) {
	elem.clearPointers()
	raw.pool.inboundElements.Put(elem)
}

func (elem *QueueInboundElement) clearPointers() {
	elem = nil
}

func (elem *QueueOutboundElement) clearPointers() {
	elem = nil
}

func (raw *Raw) NewInboundElement() *QueueInboundElement {
	elem := raw.GetInboundElement()
	elem.buffer = raw.GetMessageBuffer()
	elem.Mutex = sync.Mutex{}

	return elem
}

func (raw *Raw) GetInboundElement() *QueueInboundElement {
	return raw.pool.inboundElements.Get().(*QueueInboundElement)
}

func (p *WaitPool) Get() interface{} {
	if p.max != 0 {
		p.lock.Lock()
		for atomic.LoadUint32(&p.count) >= p.max {
			p.cond.Wait()
		}
		atomic.AddUint32(&p.count, 1)
		p.lock.Unlock()
	}

	return p.pool.Get()
}

// GetMessageBuffer gets 65535bytes.
func (raw *Raw) GetMessageBuffer() *[MaxMessageSize]byte {
	return raw.pool.messageBuffers.Get().(*[MaxMessageSize]byte)
}
