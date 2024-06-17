package errcheckanalyzer

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"honnef.co/go/tools/staticcheck"
)

var mychecks []*analysis.Analyzer

func init() {

	checks := map[string]bool{
		"SA5000": true,
		"SA6000": true,
		"SA9004": true,
	}
	mychecks = []*analysis.Analyzer{
		{
			Name: "errcheck",
			Doc:  "check for unchecked errors",
			Run:  run,
		},
	}

	for _, v := range staticcheck.Analyzers {
		// добавляем в массив нужные проверки
		if checks[v.Analyzer.Name] {
			mychecks = append(mychecks, v.Analyzer)
		}
	}

	mychecks = append(mychecks, &analysis.Analyzer{
		Name: "errcheck",
		Doc:  "check for unchecked errors",
		Run:  run,
	})

}

var ErrCheckAnalyzer = &analysis.Analyzer{
	Name: "fullCheck",
	Doc:  "multycheck",
	Run:  runFull,
}

func runFull(pass *analysis.Pass) (interface{}, error) {

	multichecker.Main(mychecks...)

	var res []interface{}
	for _, an := range mychecks {
		ret, err := an.Run(pass)
		if err != nil {
			return res, err
		}
		res = append(res, ret)
	}
	return res, nil
}

var errorType = types.Universe.Lookup("error").Type().Underlying().(*types.Interface)

func isErrorType(t types.Type) bool {
	return types.Implements(t, errorType)
}

// isReturnError возвращает true, если среди возвращаемых значений есть ошибка.
func isReturnError(pass *analysis.Pass, call *ast.CallExpr) bool {
	for _, isError := range resultErrors(pass, call) {
		if isError {
			return true
		}
	}
	return false
}

// resultErrors возвращает булев массив со значениями true,
// если тип i-го возвращаемого значения соответствует ошибке.
func resultErrors(pass *analysis.Pass, call *ast.CallExpr) []bool {
	switch t := pass.TypesInfo.Types[call].Type.(type) {
	case *types.Named: // возвращается значение
		return []bool{isErrorType(t)}
	case *types.Pointer: // возвращается указатель
		return []bool{isErrorType(t)}
	case *types.Tuple: // возвращается несколько значений
		s := make([]bool, t.Len())
		for i := 0; i < t.Len(); i++ {
			switch mt := t.At(i).Type().(type) {
			case *types.Named:
				s[i] = isErrorType(mt)
			case *types.Pointer:
				s[i] = isErrorType(mt)
			}
		}
		return s
	}
	return []bool{false}
}

func run(pass *analysis.Pass) (interface{}, error) {
	expr := func(x *ast.ExprStmt) {
		// проверяем, что выражение представляет собой вызов функции,
		// у которой возвращаемая ошибка никак не обрабатывается
		if call, ok := x.X.(*ast.CallExpr); ok {
			if isReturnError(pass, call) {
				pass.Reportf(x.Pos(), "expression returns unchecked error")
			}
		}
	}
	tuplefunc := func(x *ast.AssignStmt) {
		// рассматриваем присваивание, при котором
		// вместо получения ошибок используется '_'
		// a, b, _ := tuplefunc()
		// проверяем, что это вызов функции
		if call, ok := x.Rhs[0].(*ast.CallExpr); ok {
			results := resultErrors(pass, call)
			for i := 0; i < len(x.Lhs); i++ {
				// перебираем все идентификаторы слева от присваивания
				if id, ok := x.Lhs[i].(*ast.Ident); ok && id.Name == "_" && results[i] {
					pass.Reportf(id.NamePos, "assignment with unchecked error")
				}
			}
		}
	}
	errfunc := func(x *ast.AssignStmt) {
		// множественное присваивание: a, _ := b, myfunc()
		// ищем ситуацию, когда функция справа возвращает ошибку,
		// а соответствующий идентификатор слева равен '_'
		for i := 0; i < len(x.Lhs); i++ {
			if id, ok := x.Lhs[i].(*ast.Ident); ok {
				// вызов функции справа
				if call, ok := x.Rhs[i].(*ast.CallExpr); ok {
					if id.Name == "_" && isReturnError(pass, call) {
						pass.Reportf(id.NamePos, "assignment with unchecked error")
					}
				}
			}
		}
	}

	// нужно обработать ситуацию вызова вида
	// go myfunc()
	goStmt := func(x *ast.GoStmt) {
		//if _, ok := x.Call.Fun.(*ast.Ident); ok {
		if isReturnError(pass, x.Call) {
			pass.Reportf(x.Pos(), "go call function returns unchecked error")
		}
		//}
	}

	deferStmt := func(x *ast.DeferStmt) {
		//if _, ok := x.Call.Fun.(*ast.Ident); ok {
		if isReturnError(pass, x.Call) {
			pass.Reportf(x.Pos(), "defer call function returns unchecked error")
		}
		//}
	}

	for _, file := range pass.Files {
		// функцией ast.Inspect проходим по всем узлам AST
		ast.Inspect(file, func(node ast.Node) bool {
			switch x := node.(type) {
			case *ast.ExprStmt: // выражение
				expr(x)
			case *ast.GoStmt:
				goStmt(x)
			case *ast.DeferStmt:
				deferStmt(x)
			case *ast.AssignStmt: // оператор присваивания
				// справа одно выражение x,y := myfunc()
				if len(x.Rhs) == 1 {
					tuplefunc(x)
				} else {
					// справа несколько выражений x,y := z,myfunc()
					errfunc(x)
				}
			}
			return true
		})
	}
	return nil, nil
}
