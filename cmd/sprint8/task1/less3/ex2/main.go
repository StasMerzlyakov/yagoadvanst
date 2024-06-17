package main

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

func main() {
	src := `package main
    
func main() {
     os.Exit(1)
}`

	// допишите код
	// ...

	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}
	ast.Print(fset, f)
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.Ident:
			if x.Name == "id" {
				x.Name = "Ident"
			}
		}
		return true
	})

	// печатаем дерево
	printer.Fprint(os.Stdout, fset, f)

}
