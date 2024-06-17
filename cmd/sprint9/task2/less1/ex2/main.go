package main

import (
	"crypto/rand"
	"fmt"
)

const Alphabet = 26

func GenCryptoKey(n int) (string, error) {
	// допишите код
	// ...

	res := make([]byte, n)

	data := make([]byte, n)
	j := 0

	_, err := rand.Read(data)
	if err != nil {
		return ``, err
	}

	for i := 0; i < n; i++ {
		for {
			if j == n {
				j = 0
				_, err := rand.Read(data)
				if err != nil {
					return ``, err
				}
			}
			if data[j] >= Alphabet {
				j++

			} else {
				break
			}
		}

		res[i] = 'A' + data[j]
		j++
	}

	return string(res), nil

}

func main() {
	for i := 16; i <= 64; i += 16 {
		key, err := GenCryptoKey(i)
		if err != nil {
			panic(err)
		}
		fmt.Println(key)
	}
}
