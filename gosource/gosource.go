// Package gosource allows write a go source file from a generator code.
// Do not set a go syntax, gosource make it for you.
// For a common usage, golint and gofmt pass on the generated code.
// For the next if you ask it:
// - Instruction to the function with auto indent
// - Comments to all kinds of blocks (const, globals ...)
// - Refer to github.com/ymohl-cl/gopkg to propose issues and pr
package gosource

import (
	"bytes"
	"errors"
)

// GoSource describe the content to a go source file.
type GoSource struct {
	buf      bytes.Buffer
	name     string
	comments []string
	imports  []string
	consts   map[string]Type
	globals  map[string]Type
	funcs    []*Function
}

// New return a GoSource package initialized.
// You can add many comments to describe the package.
//
//	func ExampleNew() {
//		source := New("main",
//			"generated by gosource example",
//			"DO NOT EDIT THIS FILE")
// 		Output:
// 		// Package main generated by gosource example.'
//		// DO NOT EDIT THIS FILE.'
//		// package main
func New(packageName string, comments ...string) *GoSource {
	return &GoSource{
		name:     packageName,
		comments: comments,
		consts:   make(map[string]Type),
		globals:  make(map[string]Type),
	}
}

// AddComments to your package.
// Note: each entries will be write on only one line.
func (g *GoSource) AddComments(comments ...string) {
	for _, v := range comments {
		g.comments = append(g.comments, v)
	}
}

// AddImports to the package.
// Note: each entries will define the complete import.
//
//	func ExampleAddImports() {
//		source := New()
//		source.AddImports("github.com/ymohl-cl/gopkg/gosource")
// 		Output:
//		import "github.com/ymohl-cl/gopkg/gosource"
func (g *GoSource) AddImports(imports ...string) {
	for _, v := range imports {
		g.imports = append(g.imports, v)
	}
}

// AddConst in the GoSource.
//
//	func ExampleAddConst() {
//		source := New()
//		source.AddConst("message", "string", "Hello world !")
// 		Output:
//		const (
//			message string = "Hello world !"
//		)
func (g *GoSource) AddConst(name, nameType string, value interface{}) error {
	if value == nil {
		return errors.New("value not define to the constant definition")
	}
	g.consts[name] = Type{
		name:  nameType,
		value: value,
	}
	return nil
}

// AddGlobal in the GoSource.
//
//	func ExampleAddGlobal() {
//		source := New()
//		source.AddGlobal("message", "string", "Hello world !")
// 		Output:
//		message string = "Hello world !"
func (g *GoSource) AddGlobal(name, nameType string, value interface{}) error {
	if value == nil {
		return errors.New("value not define to the global definition")
	}
	g.globals[name] = Type{
		name:  nameType,
		value: value,
	}
	return nil
}

// AddFunction in the GoSource.
func (g *GoSource) AddFunction(f *Function) error {
	if f == nil {
		return errors.New("function not define on the function definition")
	}
	g.funcs = append(g.funcs, f)
	return nil
}
