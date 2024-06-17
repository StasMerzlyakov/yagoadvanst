package main

import (
	"bytes"
	"fmt"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

type Data struct {
	ID     int    `toml:"id"`
	Name   string `toml:"name"`
	Values []byte `toml:"values"`
}

const yamlData = `
id: 101
name: Gopher
values:
- 11
- 22
- 33
`

func main() {
	// вставьте недостающий код
	// 1) десериализуйте yamlData в переменную типа Data

	var data Data
	if err := yaml.Unmarshal([]byte(yamlData), &data); err != nil {
		panic(err)
	}

	// 2) преобразуйте полученную переменную в TOML

	var buff bytes.Buffer

	if err := toml.NewEncoder(&buff).Encode(data); err != nil {
		panic(err)
	}

	// 3) выведите в консоль результат
	fmt.Printf("%v\n", buff.String())
}
