package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/GotoRen/cybozu-pre-issue/internal"
	"github.com/GotoRen/cybozu-pre-issue/pkg/config"
	"github.com/GotoRen/cybozu-pre-issue/pkg/logger"
)

func main() {
	// Loading env config.
	cfg, err := config.Get()
	if err != nil {
		log.Fatal(err)
	}

	// Initializing the logger.
	if err := logger.InitZap(cfg); err != nil {
		log.Fatal(err)
	}

	// Executing the main logic.
	if err := run(cfg); err != nil {
		log.Fatal(err)
	}
}

const bufSize = 4096

func run(cfg *config.Config) error {
	objPath := fmt.Sprintf("tests/" + cfg.FileName)
	logger.LogDebug("Input file path", objPath)

	// Input
	fp, err := os.Open(objPath)
	if err != nil {
		return fmt.Errorf("failed to open and read input text: %w", err)
	}
	defer fp.Close()

	// Ref: https://www.ren510.dev/blog/competitive-programming-tle
	r := bufio.NewReaderSize(fp, bufSize)

	elem := &internal.Element{
		Inbound:  make(chan *internal.Data),
		Outbound: make(chan *internal.Data),
		Wg:       sync.WaitGroup{},
	}

	// routine start.
	elem.SHA256Converter()
	elem.Writer()

	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			logger.LogErr("failed to read the row", "error", err)

			return fmt.Errorf("failed to read the row: %w", err)
		}

		datum := internal.PutData(line)

		datum.Mu.Lock() // Lock-1
		elem.Inbound <- datum
		elem.Outbound <- datum
	}

	close(elem.Inbound)
	close(elem.Outbound)

	elem.Wg.Wait() // main goroutine waits for other goroutines that have been added.

	return nil
}
