package parser

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go/ast"
	"go/parser"
	"testing"
)

type funcCall struct {
	name string
	args []interface{}
}

func TestA(t *testing.T) {
	in := "TrieRoute(\"abc\")"
	expr, err := parser.ParseExpr(in)
	assert.Nil(t, err)
	fmt.Printf("[1]: [%T]->  %v\n", expr, expr)

	var call *funcCall

	F := func(n ast.Node) bool {
		//		var s string
		switch x := n.(type) {
		case *ast.BasicLit:
			fmt.Printf("==>BasicLit: %v\n", x)
			//			return false
		case *ast.CallExpr:
			call = &funcCall{}
			fmt.Printf("==>CallExpr: %v\n", x)
		case *ast.Ident:
			call.name = x.Name
			fmt.Printf("==>Ident: %v\n", x)
		default:
			if x != nil {
				fmt.Printf("==>default: %v\n", x)
			}
			fmt.Println("...default")
		}
		return true
	}
	ast.Inspect(expr, F)

	/*	// sanity check
		if _, ok := x.(*ast.BinaryExpr); !ok {
			fmt.Errorf("ParseExpr(%s): got %T, want *ast.BinaryExpr", src, x)
		}
	*/

}

//
func TestPlusNormal(t *testing.T) {
	src := "a + b"
	x, err := parser.ParseExpr(src)
	assert.Nil(t, err)
	fmt.Printf("[%T]=>  %v\n", x, x) //[*ast.BinaryExpr]=>  &{a 3 + b}
}

func TestPlusStar(t *testing.T) {
	src := "a + *"
	x, err := parser.ParseExpr(src)
	assert.NotNil(t, err)
	fmt.Printf("[%T]=>  %v\n", x, x)
}
