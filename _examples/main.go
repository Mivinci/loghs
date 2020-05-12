package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// log := loghs.New(os.Stdout)
	t := time.Now()
	// loghs.Info("hey")
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.Println("message")
	since := time.Since(t)
	fmt.Println()
	fmt.Println(since)
}
