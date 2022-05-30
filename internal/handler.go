package internal

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sync"
)

func (elem *Element) Convert2SHA256() {
	defer elem.Wg.Done()

	for raw := range elem.Inbound {
		checksum := sha256.Sum256([]byte(raw.Text))
		// checksum := []byte(elem.Text)
		raw.Buffer = checksum[:]
		raw.Unlock()
	}
}

func Write(outc chan *Data, w *sync.WaitGroup) {
	defer w.Done()

	for elme := range outc {
		elme.Lock()

		fmt.Println(hex.Dump(elme.Buffer))
		// fmt.Println(string(elme.Buffer))

		elme.Unlock()
	}
}
