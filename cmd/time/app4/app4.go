package main

import (
	"fmt"
	"time"
)

func main() {

	today := time.Now()

	day := 24 * time.Hour

	fmt.Println(today)
	fmt.Println(today.Truncate(day))
}
