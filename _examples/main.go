package main

import (
	"github.com/mivinci/loghs"
)

func main() {
	// logger := loghs.New(os.Stdout)
	// t := time.Now()
	// logger.Debug().Time("2006/01/02 15:04:05").Message("hello")
	// since := time.Since(t)
	// fmt.Println()
	// fmt.Println(since)

	// stdlog := log.New(os.Stdout, "DEBUG ", log.LstdFlags)
	// t := time.Now()
	// stdlog.Printf("hello")
	// since := time.Since(t)
	// fmt.Println(since)

	loghs.Info("hello")
}
