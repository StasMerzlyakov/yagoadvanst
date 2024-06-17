package main

import (
	"fmt"
	"regexp"

	"golang.org/x/time/rate"
)

func main() {
	emails := []string{"$€§@yandex.com", "ivan@mail.ru", "john@gmailyahoo", "fedor@gmail.com", "stepan@yahoo.com", "commanderpike@gmail.com", "greta@abcd@gmail_yahoo.com", "abc.def@mail.c", "abc.def@mail#archive.com", "abc.def@mail.org"}
	// компилируем регулярку валидных адресов
	// https://help.xmatters.com/ondemand/trial/valid_email_format.htm
	rx := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	// помещаем в тот же несущий массив
	valid := emails[:0]
	// проходимся в цикле
	for _, v := range emails {
		// проверяем адрес регуляркой
		if rx.MatchString(v) {
			valid = append(valid, v)
		}
	}
	fmt.Println(valid)

	_ = rate.NewLimiter(rate.Limit(10), 5)

}
