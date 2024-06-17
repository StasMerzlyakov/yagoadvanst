package main

import (
	"fmt"
	"time"
)

func parseTime(str string) (time.Time, error) {
	return time.Parse(time.RFC3339, str)
}

func main() {
	currentTimeStr := "2021-09-19T15:59:41+03:00"
	// допишите код
	tm, err := parseTime(currentTimeStr)
	if err != nil {
		panic(err)
	}

	fmt.Println(tm.Format(time.RFC1123))
}
