package main

type IntegerG interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~int64 | ~int32
}

type Vec[T IntegerG] []T

func Check[S ~[]T, T IntegerG](s S, i T) {
	// ...
}

// var v Vec = []int64{10}
var vector Vec[uint8]

func main() {
	Check([]float32{0.56}, 3.14)
	Check([]int32{7, 8}, 10)
	Check(Vec[int64]{100, 200, 300}, 5)
	Check[Vec, int64](Vec{0, 33}, 5)
	Check[Vec, int64](Vec{0, 33}, 5)
}
