package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Slice[T any] []T

func (s *Slice[T]) Map(f func(T) T) *Slice[T] {
	for k, v := range *s {
		(*s)[k] = f(v)
	}
	return s
}

// constraints.Ordered описывает строки, целые числа и числа с плавающей точкой
// для них точно определён оператор +
func Double[T constraints.Ordered](v T) T {
	return v + v
}

func (s *Slice[T]) Reduce(r T, f func(a, e T) T) T {
	for _, v := range *s {
		r = f(r, v)
	}
	return r
}

func Sum[T constraints.Ordered](a, e T) T {
	return a + e
}

func (s *Slice[T]) Filter(allow func(e T) bool) *Slice[T] {
	var res Slice[T]
	for _, v := range *s {
		if allow(v) {
			res = append(res, v)
		}
	}
	*s = res
	return s
}

func isEven(i int) bool {
	return i%2 == 0
}

func main() {
	var si = Slice[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res := si.Filter(isEven).Map(Double[int]).Reduce(0, Sum[int])
	fmt.Println(res)
}
