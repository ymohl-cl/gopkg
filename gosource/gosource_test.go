package gosource

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// default: should be ok
	func() {
		// init
		packageName := "main"
		describe := []string{"init example and start"}
		g := New(packageName, describe...)

		// assert
		if assert.NotNil(t, g) {
			assert.Equal(t, packageName, g.name)
			assert.Equal(t, describe, g.comments)
			assert.NotNil(t, g.consts)
			assert.NotNil(t, g.globals)
		}
	}()
}

func TestAddComments(t *testing.T) {
	// default: should be ok
	func() {
		// init
		comment1 := "comment_1"
		comment2 := "comment_2"
		g := &GoSource{}
		g.AddComments(comment1, comment2)

		// assert
		if assert.Len(t, g.comments, 2) {
			assert.Equal(t, comment1, g.comments[0])
			assert.Equal(t, comment2, g.comments[1])
		}
	}()
}

func TestAddImports(t *testing.T) {
	// default: should be ok
	func() {
		// init
		import1 := "net/http"
		import2 := "fmt"
		g := &GoSource{}
		g.AddImports(import1, import2)

		// assert
		if assert.Len(t, g.imports, 2) {
			assert.Equal(t, import1, g.imports[0])
			assert.Equal(t, import2, g.imports[1])
		}
	}()
}

func TestAddConst(t *testing.T) {
	// should return an error because value is nil
	func() {
		// init
		g := New("main", "init example")
		err := g.AddConst("message", "string", nil)

		// assert
		if assert.Error(t, err) {
			assert.Equal(t, "value not define to the constant definition", err.Error())
		}
	}()

	// default: should be ok
	func() {
		// init
		name := "message"
		nameType := "string"
		value := "message example"
		g := New("main", "init example")
		err := g.AddConst(name, nameType, value)

		// assert
		if assert.NoError(t, err) {
			assert.Equal(t, nameType, g.consts[name].name)
			assert.Equal(t, value, g.consts[name].value)
		}
	}()
}

func TestAddGlobal(t *testing.T) {
	// should return an error because value is nil
	func() {
		// init
		g := New("main", "init example")
		err := g.AddGlobal("message", "string", nil)

		// assert
		if assert.Error(t, err) {
			assert.Equal(t, "value not define to the global definition", err.Error())
		}
	}()

	// default: should be ok
	func() {
		// init
		name := "message"
		nameType := "string"
		value := "message example"
		g := New("main", "init example")
		err := g.AddGlobal(name, nameType, value)

		// assert
		if assert.NoError(t, err) {
			assert.Equal(t, nameType, g.globals[name].name)
			assert.Equal(t, value, g.globals[name].value)
		}
	}()
}

func TestAddFunction(t *testing.T) {
	// should return an error because function is nil
	func() {
		// init
		g := &GoSource{}
		err := g.AddFunction(nil)

		// assert
		if assert.Error(t, err) {
			assert.Equal(t, "function not define on the function definition", err.Error())
		}
	}()

	// default: should be ok
	func() {
		// init
		g := &GoSource{}
		fName := `function_test`
		fContent := `\tfmt.Println("Hello world")`
		err := g.AddFunction(&Function{
			name:    fName,
			content: fContent,
		})

		// assert
		if assert.NoError(t, err) {
			assert.Len(t, g.funcs, 1)
			assert.Equal(t, fName, g.funcs[0].name)
			assert.Equal(t, fContent, g.funcs[0].content)
		}
	}()
}
