package internal

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

// RoutineConvert2SHA256 calculates the checksum in SHA256 for data stored in InboundChannels.
func (elem *Element) RoutineConvert2SHA256() {
	defer elem.Wg.Done()

	for raw := range elem.Inbound {
		checksum := sha256.Sum256(raw.Text)
		raw.Buffer = checksum[:]
		raw.Mu.Unlock() // UnLock-1
	}
}

// RoutineWriter outputs the data stored in OutboundChannels with HEX-Dump.
func (elem *Element) RoutineWrite() {
	defer elem.Wg.Done()

	for raw := range elem.Outbound {
		raw.Mu.Lock() // Lock-2
		fmt.Fprintln(os.Stdout, "raw text:", string(raw.Text))
		fmt.Fprintln(os.Stdout, hex.Dump(raw.Buffer))
		raw.Mu.Unlock() // UnLock-2
	}
}
