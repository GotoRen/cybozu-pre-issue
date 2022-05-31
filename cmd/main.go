package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

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

	// Input
	fp, err := os.Open(obj_path)
	if err != nil {
		logger.LogErr("Failed to open and read input text", "error", err)
	}

	// Ref: https://qiita.com/ren510dev/items/38fe6d09831d08fde537
	r := bufio.NewReaderSize(fp, 4096)
	defer fp.Close()

	elem := &internal.Element{
		Inbound:  make(chan *internal.Raw),
		Outbound: make(chan *internal.Raw),
	}

	// routine start.
	elem.SHA256Converter()
	elem.Writer()

	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			logger.LogErr("Failed to read the row", "error", err)
		}

		raw := internal.Raw{
			Buffer: line,
		}

		raw.Lock() // Lock-1
		elem.Inbound <- &raw
		elem.Outbound <- &raw
	}

	close(elem.Inbound)
	close(elem.Outbound)

	elem.Wg.Wait() // main goroutine waits for other goroutines that have been added.
}
