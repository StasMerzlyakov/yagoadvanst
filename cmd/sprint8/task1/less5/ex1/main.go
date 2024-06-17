package main

import "sort"

func main() {
	sl := []string{"foo", "bar", "buzz"}
	sl = sort.StringSlice(sl) // sort.StringSlice — это не функция, а тип, выражение не отсортирует sl
	// чтобы отсортировать, нужно сделать sort.StringSlice(sl).Sort()
}
