package main

import (
	"fmt"
	"time"
)

func main() {

	birthday := time.Date(2093, time.November, 26, 0, 0, 0, 0, time.Local)

	days := time.Until(birthday).Truncate(24 * time.Hour)

	fmt.Println(days / 24)

	/*birthday := time.Date(2093, time.November, 26, 0, 0, 0, 0, time.Local)
	duration := time.Until(birthday)
	days := int(duration.Hours() / 24)
	// альтернативный вариант
	// days := int(duration / time.Hour / 24)
	fmt.Println(days) */

}
