package extractor

import (
	"errors"
	"fmt"
	"go/ast"
	"regexp"
	"strings"
)

const (
	errEmptyScope     = "scope file not define"
	errStructNotFound = "struct type not found"
	objSuffix         = "Obj"
)

// Object describe a struct type
type Object struct {
	Type      string
	Name      string
	TableName string
	Fields    []Field
}

// Field describe list fields to the struct
type Field struct {
	Name string
	Type string
	Tags []string
}

// Objects list extracted from the ast file parameter
func Objects(file *ast.File) ([]Object, error) {
	var objs []Object

	if file.Scope == nil {
		return nil, errors.New(errEmptyScope)
	}
	for name, t := range file.Scope.Objects {
		var astFields []*ast.Field
		var obj Object
		var ok bool

		if astFields, ok = extractFields(t.Decl); !ok {
			continue
		}
		obj.Type = name
		obj.Name = objectName(name)
		obj.TableName = objectTable(name)
		for _, astField := range astFields {
			fields, err := extractField(astField)
			if err != nil {
				return nil, err
			}
			obj.Fields = append(obj.Fields, fields...)
		}
		objs = append(objs, obj)
	}
	if len(objs) == 0 {
		return nil, errors.New(errStructNotFound)
	}
	return objs, nil
}

func objectName(objType string) string {
	objName := strings.ToLower(objType)
	if strings.Compare(objName, objType) == 0 {
		objName += objSuffix
	}
	return objName
}

func objectTable(objType string) string {
	tableName := strings.ToLower(objType)
	tableName += "s"
	return tableName
}

func extractField(f *ast.Field) ([]Field, error) {
	var fields []Field
	var err error

	for _, name := range f.Names {
		var tag []string
		if f.Tag != nil {
			if tag, err = extractTags(f.Tag.Value); err != nil {
				return nil, err
			}
		}
		field := Field{
			Name: name.Name,
			Type: fmt.Sprintf("%v", f.Type),
			Tags: tag,
		}
		fields = append(fields, field)
	}
	return fields, nil
}

func extractTags(value string) ([]string, error) {
	r, err := regexp.Compile(`crud:"([^"]+)"`)
	if err != nil {
		return nil, err
	}
	value = r.FindString(value)
	value = strings.Replace(value, "crud:", "", -1)
	value = strings.Trim(value, `"`)
	tags := strings.Split(value, ",")
	return tags, nil
}

func extractFields(decl interface{}) ([]*ast.Field, bool) {
	if decl == nil {
		return nil, false
	}
	switch ts := decl.(type) {
	case *ast.TypeSpec:
		switch st := ts.Type.(type) {
		case *ast.StructType:
			if st.Fields == nil {
				return nil, false
			}
			if len(st.Fields.List) == 0 {
				return nil, false
			}
			return st.Fields.List, true
		default:
			return nil, false
		}
	default:
		return nil, false
	}
}
