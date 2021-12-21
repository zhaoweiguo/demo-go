/*
From: 官方 Example: https://pkg.go.dev/go/parser@go1.17.1#example-ParseFile
*/
package parser

import (
	"github.com/stretchr/testify/assert"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"testing"
)

func TestFile(t *testing.T) {
	fset := token.NewFileSet() // positions are relative to fset

	src := `package foo

import (
	"fmt"
	"time"
)

func bar() {
	fmt.Println(time.Now())
}`

	// Parse src but stop after processing the imports.
	f, err := parser.ParseFile(fset, "", src, parser.ImportsOnly)
	assert.Nil(t, err)

	assert.Equal(t, `"fmt"`, f.Imports[0].Path.Value)
	assert.Equal(t, `"time"`, f.Imports[1].Path.Value)

	assert.Equal(t, "STRING", f.Imports[0].Path.Kind.String())
	assert.Equal(t, token.Pos(24), f.Imports[0].Path.ValuePos)

	// Print the imports from the file's AST.
	for _, s := range f.Imports {
		log.Println(s.Doc)
		log.Println(s.Path)
	}

	ast.Print(fset, f)

}
