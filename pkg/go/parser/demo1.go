package main

import(
	"go/parser"
	"go/ast"
	"fmt"
)

type funcCall struct {
	name string
	args []interface{}
}

func main() {
	in := "TrieRoute(\"abc\")"
	expr, err := parser.ParseExpr(in)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("[1]: [%T]->  %v\n", expr, expr)

	var call *funcCall
	ast.Inspect(expr, func(n ast.Node) bool {
//		var s string
		switch x := n.(type) {
		case *ast.BasicLit:
			fmt.Printf("3%v\n", x)
//			return false
		case *ast.CallExpr:
			call = &funcCall{}
			fmt.Printf("4%v\n", x)
		case *ast.Ident:
			call.name = x.Name
			fmt.Printf("5%v\n", x)
		default:
			if x != nil {
				fmt.Printf("6%v\n", x)
			}
			fmt.Println(7)
		}
		return true
	})


	src := "a + b"
	x, err := parser.ParseExpr(src)
	if err != nil {
		fmt.Errorf("ParseExpr(%s): %v", src, err)
	}
	fmt.Printf("[%T]%v\n", x, x)

/*	// sanity check
	if _, ok := x.(*ast.BinaryExpr); !ok {
		fmt.Errorf("ParseExpr(%s): got %T, want *ast.BinaryExpr", src, x)
	}
*/

	src = "a + *"
	expr, err = parser.ParseExpr(src)
	if err == nil {
		fmt.Errorf("ParseExpr(%s): got no error", src)
	}
	fmt.Printf("[%T]%v\n", expr, expr)



}





