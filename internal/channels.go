package internal

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sync"
)

// RoutineConvert2SHA256 calculates the checksum in SHA256 for data stored in InboundChannels.
func (elem *Element) RoutineConvert2SHA256() {
	defer elem.Wg.Done()

	for raw := range elem.Inbound {
		checksum := sha256.Sum256(raw.Text)
		// checksum := raw.Text
		raw.Buffer = checksum[:]
		raw.Unlock() // UnLock-1
	}
}

// RoutineWriter outputs the data stored in OutboundChannels with HEX-Dump.
func (elem *Element) RoutineWrite() {
	defer elem.Wg.Done()

	for raw := range elem.Outbound {
		raw.Lock() // Lock-3
		fmt.Println(hex.Dump(raw.Buffer))
		raw.Unlock() // UnLock-2
	}
}

func NewInboundElement() *QueueInboundElement {
	data := GetInbountElement()
	data.Text = GetMessageBuffer()
	data.Mutex = sync.Mutex{}

	return data
}

type QueueElement struct {
	sync.Mutex
	buffer *[MaxMessageSize]byte // slice holding the packet data
	packet []byte                // slice of "buffer" (always!)
}

// MaxMessageSize is 65535=(2^16)-1.
const MaxMessageSize = (1 << Bit16) - 1

// GetOutboundElement gets QueueOutboundElement.
func (device *Device) GetOutboundElement() *QueueOutboundElement {
	return device.pool.outboundElements.Get().(*QueueOutboundElement)
}
