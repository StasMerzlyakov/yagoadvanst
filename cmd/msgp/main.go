package main

import "fmt"

type AccountBalance struct {
	AccountIdHash []byte           `msg:"account_id_hash"`
	Amounts       []CurrencyAmount `msg:"amounts"`
	IsBlocked     bool             `msg:"is_blocked"`
}

type CurrencyAmount struct {
	Amount   int64 // здесь не будем использовать структурные теги
	Decimals int8
	Symbol   string
}

type User struct {
	Name   string `msg:"name"`
	Age    int
	Phone  string `json:"-"`
	Email  string `msg:"-"`
	Active bool   `msg:",omitempty"`
}

//go:generate msgp

func main() {
	balance := AccountBalance{
		AccountIdHash: []byte{0x10, 0x20, 0x0A, 0x0B},
		Amounts: []CurrencyAmount{
			{Amount: 1000000, Decimals: 2, Symbol: "RUB"},
			{Amount: 2510, Decimals: 2, Symbol: "USD"},
		},
		IsBlocked: true,
	}
	// сериализуем значение переменной balance
	msgpBz, err := balance.MarshalMsg(nil)
	if err != nil {
		panic(err)
	}

	var balanceCopy AccountBalance
	// декодируем данные в переменную типа AccountBalance
	if _, err := balanceCopy.UnmarshalMsg(msgpBz); err != nil {
		panic(err)
	}

	// для сравнения выведем оригинальное и полученное значения
	fmt.Printf("balanceInit: %+v\n", balance)
	fmt.Printf("balanceCopy: %+v\n", balanceCopy)

}
