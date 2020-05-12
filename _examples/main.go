package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mivinci/loghs"
)

func main() {
	log := loghs.New(os.Stdout)
	t := time.Now()
	log.Info("hey")
	fmt.Println(time.Since(t))
}
