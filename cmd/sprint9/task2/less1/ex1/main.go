package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"

	"time"
)

func generate(ctx context.Context, rnd *rand.Rand, n int, ch chan<- int) {

	for n > 0 {
		n--
		val := rnd.Intn(6) + 1
		select {
		case <-ctx.Done():
			return
		case ch <- val:
		}
	}
}

func summ(ctx context.Context, ch <-chan int) {
	cnt := map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
	}

	for {
		select {
		case <-ctx.Done():
			return
		case vl, ok := <-ch:
			if ok {
				cnt[vl] = cnt[vl] + 1
				cnt[0] = cnt[0] + 1
			} else {
				count := cnt[0]
				if count > 0 {
					all := float32(cnt[0])

					fmt.Printf("n= %d [%f, %f, %f, %f, %f, %f]\n", count,
						float32(cnt[1])/all, float32(cnt[2])/all,
						float32(cnt[3])/all, float32(cnt[4])/all,
						float32(cnt[5])/all, float32(cnt[6])/all)
					return

				} else {
					fmt.Println("no data")
				}
			}
		}
	}
}

func doTest(count int) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	var wg sync.WaitGroup

	ctx := context.Background()

	ch := make(chan int)

	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			close(ch)
		}()

		generate(ctx, rnd, count, ch)
	}()

	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()

		summ(ctx, ch)
	}()

	wg.Wait()
}

func main() {
	doTest(100)
	doTest(1000)
	doTest(10000)
}
