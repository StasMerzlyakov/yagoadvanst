package main

import (
	"fmt"

	"golang.org/x/sync/errgroup"
)

type Name string

var names = []Name{"Anna", "Ivan", "Fedor", "Katya", "Gleb"}

// Здесь напишите метод для Name
func (n Name) Hello() error {
	fmt.Printf("Hello %s\n", n)
	return nil
}

func main() {
	g := &errgroup.Group{}

	for _, name := range names {
		g.Go(name.Hello)
	}

	g.Wait()
}
