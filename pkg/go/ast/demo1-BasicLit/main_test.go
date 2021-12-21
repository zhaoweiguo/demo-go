package BasicLit

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

/*
手工造数据构造`ast.BasicLit结构`
*/
func TestInt(t *testing.T) {
	var lit9527 = &ast.BasicLit{
		Kind:  token.INT,
		Value: "9527",
	}
	ast.Print(nil, lit9527)
}

func TestFloat(t *testing.T) {
	var lit9527 = &ast.BasicLit{
		Kind:  token.FLOAT,
		Value: "9527",
	}
	ast.Print(nil, lit9527)
}

/*
通过parser.ParseExpr函数构造`ast.BasicLit结构`
*/
func TestParseInt(t *testing.T) {
	expr, _ := parser.ParseExpr(`9527`)
	ast.Print(nil, expr)
}
func TestParseString(t *testing.T) {
	expr, _ := parser.ParseExpr(`"9527"`)
	ast.Print(nil, expr)
}

/*
手工造数据构造`ast.Indent结构`
*/
func TestParseX(t *testing.T) {
	expr, _ := parser.ParseExpr(`x`)
	ast.Print(nil, expr)
}

/*
通过parser.ParseExpr函数构造`ast.Indent结构`
*/
func TestIndent(t *testing.T) {
	indent := ast.NewIdent(`x`)
	ast.Print(nil, indent)
}

func TestBinaryExpr(t *testing.T) {
	expr, _ := parser.ParseExpr(`1+2*3`)
	ast.Print(nil, expr)
}
