package main

import (
	"fmt"
	"net/http"
)

type reqFn = func(r *http.Request)

var Funcs = make(map[string]reqFn) // пустой интерфейс может принять любое значение

func DBInsert(r *http.Request) {
	// логика вставки
	fmt.Println("DBInsert")
}

func DBDelete(r *http.Request) {
	// логика удаления
}

func main() {
	Funcs["DBInsert"] = DBInsert
	Funcs["DBDelete"] = DBDelete
	Funcs["DBChange"] = func(r *http.Request) {
		// логика изменения
	}
	// ...

	r := new(http.Request)
	Funcs["DBInsert"](r)
}
