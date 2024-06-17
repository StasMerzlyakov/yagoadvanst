package pkg1

import "fmt"

func mulfunc(i int) (int, error) {
	return i * 2, nil
}

func errCheckFunc() {

	// формулируем ожидания: анализатор должен находить ошибку,
	// описанную в комментарии want
	go createError() // want "go call function returns unchecked error"

	defer createError() // want "defer call function returns unchecked error"

	defer func() {
		createError() // want "expression returns unchecked error"
	}()

	defer func() {
		_ = createError() // want "assignment with unchecked error"
	}()

	mulfunc(5)           // want "expression returns unchecked error"
	res, _ := mulfunc(5) // want "assignment with unchecked error"
	fmt.Println(res)     // want "expression returns unchecked error"
}

func createError() error {
	return fmt.Errorf("hello err")
}
