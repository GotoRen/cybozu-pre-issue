package internal

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// RoutineConvert2SHA256 calculates the checksum in SHA256 for data stored in InboundChannels.
func (elem *Element) RoutineConvert2SHA256() {
	defer elem.Wg.Done()

	for raw := range elem.Inbound {
		checksum := sha256.Sum256(raw.Buffer)
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
