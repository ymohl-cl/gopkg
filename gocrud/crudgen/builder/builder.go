package builder

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"

	"github.com/ymohl-cl/gopkg/gocrud/crudgen/extractor"
	"github.com/ymohl-cl/gopkg/gosource"
)

// driver supported
const (
	PSQL     = "postgres"
	PQSLName = "PSQL"
)

// Builder interface to each drivers
type Builder interface {
	buildHeader() error
	buildCreate() error
	buildRead() error
	buildUpdate() error
	buildDelete() error
	source() *gosource.GoSource
}

// New return a builder to create a new crud source file
func New(driver string, source string) (Builder, error) {
	var err error
	var sourcepath string
	var file *ast.File

	if sourcepath, err = filepath.Abs(source); err != nil {
		return nil, err
	}
	fset := token.NewFileSet()
	if file, err = parser.ParseFile(fset, sourcepath, nil, 0); err != nil {
		return nil, err
	}
	switch driver {
	case PSQL:
		var b Builder
		if b, err = newPSQL(file); err != nil {
			return nil, err
		}
		return b, nil
	default:
		return nil, errors.New("unsupported bdd driver")
	}
}

// Build the crud source file to the builder target
// if the builder is not initialized one error will be return
// if the build faile one error will be return with the context error
// if write file failed, one error will be return with the context error
func Build(b Builder) error {
	var err error
	var g *gosource.GoSource
	var buf []byte

	if b == nil {
		return errors.New("builder not initialized")
	}
	if err = b.buildHeader(); err != nil {
		return err
	}
	if err = b.buildCreate(); err != nil {
		return err
	}
	if err = b.buildRead(); err != nil {
		return err
	}
	if err = b.buildUpdate(); err != nil {
		return err
	}
	if err = b.buildDelete(); err != nil {
		return err
	}
	g = b.source()
	if buf, err = g.Bytes(); err != nil {
		return err
	}
	fmt.Println(string(buf))
	return nil
}

func createFunction(driver, driverName, driverType string, obj extractor.Object) *gosource.Function {
	fName := "Create" + driver + obj.Type
	f := gosource.NewFunction(fName)
	f.AddArg(driverName, driverType)
	f.AddArg(obj.Name, obj.Type)
	f.AddRet("_", obj.Type)
	f.AddRet("_", "error")
	return f
}

func deleteFunction(driver, driverName, driverType string, obj extractor.Object) *gosource.Function {
	fName := "Delete" + driver + obj.Type
	f := gosource.NewFunction(fName)
	f.AddArg(driverName, driverType)
	f.AddArg(obj.Name, obj.Type)
	f.AddRet("_", obj.Type)
	f.AddRet("_", "error")
	return f
}

func readFunction(driver, driverName, driverType string, obj extractor.Object) *gosource.Function {
	fName := "Read" + driver + obj.Type
	f := gosource.NewFunction(fName)
	f.AddArg(driverName, driverType)
	f.AddArg(obj.Name, obj.Type)
	f.AddRet("_", "[]"+obj.Type)
	f.AddRet("_", "error")
	return f
}
