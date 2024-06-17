package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
func main() {
	rnd := rand.New(rand.NewSource(77))
	// всегда будет одна и та же последовательность
	// 460 3733561740 1284141128648234027
	fmt.Println(rnd.Intn(1000), rnd.Uint32(), rnd.Uint64())

	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	// будут печататься разные числа
	fmt.Println(rnd.Intn(99), rnd.Uint32())
}
*/

/*
// RandomString возвращает случайную шестнадцатеричную строку.
// Длина строки будет равна 2*n.
func RandomString(rnd *rand.Rand, n int) (string, error) {
	b := make([]byte, n)
	_, err := rnd.Read(b)
	if err != nil {
		return ``, err
	}
	return hex.EncodeToString(b), nil
}

func main() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	rand := func(n int) string {
		s, err := RandomString(rnd, n)
		if err != nil {
			log.Fatal(err)
		}
		return s
	}
	fmt.Println(rand(4), rand(5), rand(6))
}
*/

var (
	alphabet = []rune(`23456789@#$%^&/?+=!` +
		`ABCEFGHJKLMNPQRSTUVWXYZ` +
		`abcdefghijkmnopqrstuvwxyz`)
	alength = len(alphabet)
	rnd     *rand.Rand
)

// RandPsw возвращает случайный пароль указанной длины.
func RandPsw(count int) string {
	b := make([]rune, count)
	for i := range b {
		b[i] = alphabet[rnd.Intn(alength)]
	}
	return string(b)
}

func main() {
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println(RandPsw(6), RandPsw(8), RandPsw(11))
}
