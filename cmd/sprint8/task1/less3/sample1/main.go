package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

func main() {
	// исходный код, который будем разбирать
	src := `package main
import "fmt"

func main() {
    fmt.Println("Hello, world!")
}`

	// дерево разбора AST ассоциируется с набором исходных файлов FileSet
	fset := token.NewFileSet()
	// парсер может работать с файлом
	// или исходным кодом, переданным в виде строки
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}
	// печатаем дерево
	ast.Print(fset, f)

	ast.Inspect(f, func(n ast.Node) bool {
		// проверяем, какой конкретный тип лежит в узле
		switch x := n.(type) {
		case *ast.CallExpr:
			// ast.CallExpr представляет вызов функции или метода
			fmt.Printf("CallExpr %v: ", fset.Position(x.Fun.Pos()))
			printer.Fprint(os.Stdout, fset, x)
			fmt.Println()
		case *ast.FuncDecl:
			// ast.FuncDecl представляет декларацию функции
			fmt.Printf("FuncDecl %s %v: ", x.Name.Name, fset.Position(x.Pos()))
			printer.Fprint(os.Stdout, fset, x)
			fmt.Println()
		}
		return true
	})
}
