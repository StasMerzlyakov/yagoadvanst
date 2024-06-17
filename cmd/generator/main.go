package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	inCh := gen(2, 3)
	ch1 := square(inCh)
	ch2 := square(inCh)
	for n := range fanIn(ch1, ch2) {
		fmt.Println(n)
	}
}

func gen(nums ...int) chan int {
	outCh := make(chan int)
	go func() {
		defer close(outCh)
		for _, n := range nums {
			outCh <- n
		}
	}()

	return outCh
}

func square(inCh chan int) chan int {
	outCh := make(chan int)
	go func() {
		defer close(outCh)
		for n := range inCh {
			outCh <- n * n
		}
	}()

	return outCh
}

// эту функцию нужно реализовать
func fanIn(chs ...chan int) chan int {
	outCh := make(chan int)

	time.Sleep(123)

	var wg sync.WaitGroup
	for _, ch := range chs {
		wg.Add(1)
		inCh := ch
		go func() {
			defer wg.Done()
			for v := range inCh {
				outCh <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		defer close(outCh)
	}()

	return outCh
}
