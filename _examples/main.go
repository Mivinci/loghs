package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mivinci/loghs"
)

func main() {
	logger := loghs.New(os.Stdout)
	t := time.Now()
	logger.Debug().String("hey").Int64(24).Message("xjj")
	since := time.Since(t)
	fmt.Println()
	fmt.Println(since)

	// stdlog := log.New(os.Stdout, "DEBUG ", log.LstdFlags)
	// t := time.Now()
	// stdlog.Printf("hello")
	// since := time.Since(t)
	// fmt.Println(since)

	// loghs.Info("hello")
}
