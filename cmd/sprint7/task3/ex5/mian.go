package main

import "fmt"

func main() {
	var NonconstructedSlice []int
	EmptySlice := make([]int, 0)
	fmt.Printf("%v\n", NonconstructedSlice == nil)
	fmt.Printf("%v\n", EmptySlice == nil)
}
