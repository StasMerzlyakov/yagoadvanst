package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		delta := time.Since(start).Truncate(time.Second)
		fmt.Printf("deltaSec: %v\n", delta)
	}
}
