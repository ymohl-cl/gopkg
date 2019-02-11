package gosource

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFunction(t *testing.T) {
	// default: should be ok
	func() {
		// init
		fName := `main`
		f := NewFunction(fName)

		// assert
		if assert.NotNil(t, f) {
			assert.Equal(t, fName, f.name)
			assert.Equal(t, "", f.content)
		}
	}()
}

func TestSetContent(t *testing.T) {
	// default: should be ok
	func() {
		// init
		fName := `main`
		fContent := `\tfmt.Println("Hello world")`
		f := NewFunction(fName)

		// assert
		if assert.NotNil(t, f) {
			f.SetContent(fContent)
			assert.Equal(t, fName, f.name)
			assert.Equal(t, fContent, f.content)
		}
	}()
}

func TestAddArg(t *testing.T) {
	// default: should be ok
	func() {
		// init
		argName := `message`
		argType := `string`
		f := NewFunction(`main`)
		f.AddArg(argName, argType)

		// assert
		if assert.NotNil(t, f.args) {
			assert.Len(t, f.args, 1)
			assert.Equal(t, argName, f.args[0].Name)
			assert.Equal(t, argType, f.args[0].Type)
		}
	}()
}

func TestAddRet(t *testing.T) {
	// default: should be ok
	func() {
		// init
		retName := `err`
		retType := `error`
		f := NewFunction(`main`)
		f.AddRet(retName, retType)

		// assert
		if assert.NotNil(t, f.rets) {
			assert.Len(t, f.rets, 1)
			assert.Equal(t, retName, f.rets[0].Name)
			assert.Equal(t, retType, f.rets[0].Type)
		}
	}()
}
