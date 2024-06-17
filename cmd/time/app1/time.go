package main

import (
	"fmt"
	"time"
)

func parseTime(layout, value string) {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	fmt.Println(t.Format(`02.01.06 15:04:05`))
}

func main() {
	parseTime(time.UnixDate, "Tue Jun 1 22:16:03 MSK 2022")
	parseTime("01/02/2006 15:04:05", "01/26/2023 14:43:00")
}
