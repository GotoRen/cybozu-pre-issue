package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"runtime"
	"sync"

	"github.com/GotoRen/cyboze_pre_issue/internal/logger"
	"github.com/joho/godotenv"
)

type data struct {
	sync.Mutex
	text   string
	buffer []byte
}

func HASH(inc chan *data, w *sync.WaitGroup) {
	defer w.Done()

	for elme := range inc {

		checksum := sha256.Sum256([]byte(elme.text))
		elme.buffer = checksum[:]

		elme.Unlock()
	}
}

func Write(outc chan *data, w *sync.WaitGroup) {
	defer w.Done()

	for elme := range outc {
		elme.Lock()

		fmt.Println(hex.Dump(elme.buffer))

		elme.Unlock()
	}
}

func init() {
	err := godotenv.Load()
	if err != nil {
		logger.LogErr("Error loading .env file", "error", err)
	}
}

func main() {
	logger.InitZap()
	var wait sync.WaitGroup
	inc := make(chan *data)
	outc := make(chan *data)
	cpus := runtime.NumCPU()

	obj_path := fmt.Sprintf("tests/" + os.Getenv("FILE"))
	fmt.Println(obj_path)

	wait.Add(cpus)
	for i := 0; i < cpus; i++ {
		go HASH(inc, &wait)
	}

	wait.Add(1)
	go Write(outc, &wait)

	f, _ := os.Open(obj_path)
	defer f.Close()

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		line := sc.Text()

		data := data{
			text:   line,
			buffer: nil,
			Mutex:  sync.Mutex{},
		}
		data.Lock()

		inc <- &data
		outc <- &data
	}

	close(inc)
	close(outc)

	wait.Wait()
}
