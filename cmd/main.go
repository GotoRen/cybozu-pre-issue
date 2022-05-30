package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/GotoRen/cyboze_pre_issue/internal"
	"github.com/GotoRen/cyboze_pre_issue/internal/logger"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logger.LogErr("Error loading .env file", "error", err)
	}
}

func main() {
	logger.InitZap()

	obj_path := fmt.Sprintf("tests/" + os.Getenv("FILE"))
	logger.LogDebug("[DEBUG]", "Input file path", obj_path)

	f, err := os.Open(obj_path)
	if err != nil {
		logger.LogErr("Failed to open and read input text", "error", err)
	}
	defer f.Close()

	elem := &internal.Element{
		Inbound:  make(chan *internal.Data),
		Outbound: make(chan *internal.Data),
	}

	elem.RoutineSHA256Converter()
	elem.RoutineWriter()

	// ref: https://qiita.com/ren510dev/items/38fe6d09831d08fde537
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()

		// fmt.Println(line)
		data := internal.Data{
			Text:   line,
			Buffer: nil,
			Mutex:  sync.Mutex{},
		}
		data.Lock()

		elem.Inbound <- &data
		elem.Outbound <- &data
	}

	close(elem.Inbound)
	close(elem.Outbound)

	elem.Wg.Wait() // main goroutine waits for other goroutines that have been added.
}
