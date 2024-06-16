package examples

import (
	"math/rand"
	"os"
	"time"
)

func Delay() {
	if os.Getenv("EXAMPLE_DELAY") == "true" {
		time.Sleep(time.Duration(10*rand.Intn(100)) * time.Millisecond)
	}
}
