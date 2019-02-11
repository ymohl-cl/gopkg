package gosource

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var expectedSRC = `// Package main test the writer GoSource
// example test main
package main

import "fmt"

const (
	message string = "Hello world"
)

var nbr int = 1

func test(n int, m string) (err error) {
	for i := 0; i < n; i++ {
		fmt.Println(m)
	}
	return nil
}

func main() {
	_ = test(nbr, message)
}
`

func TestBytes(t *testing.T) {
	// should return an error because file don't match with go syntax
	func() {
		// init
		g := New("", "test")
		f := NewFunction("main")
		f.SetContent(`fmt.Println()`)
		g.AddFunction(f)
		content, err := g.Bytes()

		// assert
		if assert.Error(t, err) {
			assert.Len(t, content, 0)
			assert.Equal(t, "4:1: expected 'IDENT', found 'func'", err.Error())
		}
	}()

	// default: should be ok
	func() {
		// init
		g := New("main", "test the writer GoSource")
		g.AddComments("example test main")
		g.AddImports("fmt")
		g.AddConst("message", "string", "Hello world")
		g.AddGlobal("nbr", "int", 1)
		f := NewFunction("test")
		f.SetContent("\tfor i := 0; i < n; i++ {\n\t\tfmt.Println(m)\n\t}\n\treturn nil\n")
		f.AddArg("n", "int")
		f.AddArg("m", "string")
		f.AddRet("err", "error")
		g.AddFunction(f)
		f = NewFunction("main")
		f.SetContent("\t_ = test(nbr, message)\n")
		g.AddFunction(f)
		content, err := g.Bytes()

		// assert
		if assert.NoError(t, err) {
			assert.Equal(t, expectedSRC, string(content))
		}
	}()
}
