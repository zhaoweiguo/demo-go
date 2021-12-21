package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

func TestDemo2(t *testing.T) {
	fset := token.NewFileSet()
	fmt.Printf("[1]: %T => %v\n", fset, fset)

	pkgs, err := parser.ParseDir(fset, filepath.Join(os.Getenv("GOROOT"), "src", "net"), nil, parser.PackageClauseOnly)
	if err != nil {
		panic(err.Error())
	}
	for name := range pkgs {
		fmt.Println(name)
	}

	ast.Print(fset, pkgs)

}
