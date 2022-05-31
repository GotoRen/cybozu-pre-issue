package main

import (
	"fmt"
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
	var raw internal.Raw

	fmt.Println("Hello, World!")

	/* ---------------------------------------------------------------------------- */
	obj_path := fmt.Sprintf("tests/" + os.Getenv("FILE"))
	logger.LogDebug("[DEBUG]", "Input file path", obj_path)

	// Input
	fp, err := os.Open(obj_path)
	if err != nil {
		logger.LogErr("Failed to open and read input text", "error", err)
	}
	defer fp.Close()

	// Ref: https://qiita.com/ren510dev/items/38fe6d09831d08fde537
	// r := bufio.NewReaderSize(fp, 4096)

	/* ---------------------------------------------------------------------------- */
	raw.NewRoutine()
	raw.RoutineSequentialReader(fp)

	// elem := &internal.Element{
	// 	Inbound:  make(chan *internal.Raw),
	// 	Outbound: make(chan *internal.Raw),
	// }

	// routine start.
	// raw.SHA256Converter()
	// raw.Writer()

}
