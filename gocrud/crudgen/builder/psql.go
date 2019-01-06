package builder

import (
	"fmt"
	"go/ast"
	"reflect"
	"strconv"
	"strings"

	"github.com/ymohl-cl/gopkg/gocrud/crudgen/extractor"
	"github.com/ymohl-cl/gopkg/gosource"
)

const (
	driverName = "PSQL"
	describe   = "crud controllers to psql bdd service"
	typeDriver = "driver.Conn"
	varDriver  = "db"
	varSuffix  = "Obj"
)

// PSQLBuilder implement the builder interface
type PSQLBuilder struct {
	*gosource.GoSource
	file *ast.File
	objs []extractor.Object
}

func newPSQL(file *ast.File) (Builder, error) {
	var p PSQLBuilder
	var err error

	p.GoSource = gosource.New(file.Name.Name, describe)
	p.file = file

	if p.objs, err = extractor.Objects(p.file); err != nil {
		return nil, err
	}

	//	p.help()
	return &p, nil
}

func (p *PSQLBuilder) help() {
	fmt.Println("Doc: ", p.file.Doc)
	fmt.Println("Name: ", p.file.Name)
	fmt.Println("Decls: ", p.file.Decls)
	fmt.Println("Scope: ", p.file.Scope)
	fmt.Println("Scope Outer: ", p.file.Scope.Outer)
	fmt.Println("Scope Objects: ", p.file.Scope.Objects)
	fmt.Println("Scope User Objects: ", p.file.Scope.Objects["User"])
	fmt.Println("Scope User Objects Rdecl: ", reflect.TypeOf(p.file.Scope.Objects["User"].Decl))
	switch v := p.file.Scope.Objects["User"].Decl.(type) {
	case *ast.TypeSpec:
		fmt.Println(v.Name.Name)
		fmt.Println("reglect Type: ", reflect.TypeOf(v.Type))
		switch st := v.Type.(type) {
		case *ast.StructType:
			fmt.Println("fields", st.Fields)
			for _, f := range st.Fields.List {
				fmt.Println("Type: ", f.Type)
				fmt.Println("tag: ", f.Tag)
				if f.Tag != nil {
					fmt.Println("tag value: ", f.Tag.Value)
				}
				fmt.Println("Names: ", f.Names)
			}
		default:
			fmt.Println("type is not a struct type")
		}
	default:
		fmt.Println("type not found")
	}

	fmt.Println("Scope User Objects Rdata: ", reflect.TypeOf(p.file.Scope.Objects["User"].Data))
	fmt.Println("Scope User Objects Rtype: ", reflect.TypeOf(p.file.Scope.Objects["User"].Type))
	fmt.Println("Scope User Objects decl: ", p.file.Scope.Objects["User"].Decl)
	fmt.Println("Scope User Objects data: ", p.file.Scope.Objects["User"].Data)
	fmt.Println("Scope User Objects type: ", p.file.Scope.Objects["User"].Type)

	fmt.Println("Scope User Objects: ", p.file.Scope.Objects["Truc"].Decl)
	fmt.Println("Imports: ", p.file.Imports)
	fmt.Println("Package: ", p.file.Package)
	fmt.Println("Indent: ")
	for _, i := range p.file.Unresolved {
		fmt.Printf("\tNamePos: %v\n", i.NamePos)
		fmt.Printf("\tName: %s\n", i.Name)
		if i.Obj != nil {
			fmt.Printf("\tObj Name: %s\n", i.Obj.Name)
		} else {
			fmt.Printf("\tObject is nil\n")
		}
	}
	fmt.Println("Comments: ", p.file.Comments)
}

func (p *PSQLBuilder) buildHeader() error {
	p.SetComments(
		commentGenerator,
		commentIssue,
	)
	p.SetImports(
		"database/sql",
		"github.com/lib/pq",
	)
	return nil
}

func (p *PSQLBuilder) buildCreate() error {
	var err error

	for _, obj := range p.objs {
		f := createFunction(driverName, varDriver, typeDriver, obj)
		var params string
		var values string
		var queryValues string
		for i, f := range obj.Fields {
			if i > 0 {
				values += ", "
				params += ", "
				queryValues += ", "
			}
			params += strings.ToLower(f.Name)
			values += "$" + strconv.Itoa(i+1)
			queryValues += obj.Name + "." + f.Name
		}

		params += ", version, createDate, updateDate"
		nbFields := len(obj.Fields)
		values += ", $" + strconv.Itoa(nbFields+1) + ", $" + strconv.Itoa(nbFields+2) + ", $" + strconv.Itoa(nbFields+3)
		queryValues += ", " + obj.Name + ".Version, " + obj.Name + ".CreateDate, " + obj.Name + ".UpdateDate"

		declarative := "\tvar id int\n"
		declarative += "\tvar sqlStatement string\n"
		declarative += "\tvar err error\n"
		metadata := "\n\t" + obj.Name + ".Init()\n"
		statement := "\tsqlStatement = `INSERT INTO " + obj.TableName + " (" + params + ") VALUES (" + values + ") RETURNING id`\n"
		query := "\terr = db.QueryRow(sqlStatement, " + queryValues + ").Scan(&id)\n"
		check := "\tif err != nil {\n"
		check += "\t\treturn " + obj.Type + "{}, err\n"
		check += "\t}\n"
		ret := "\t" + obj.Name + ".id = id\n"
		ret += "\treturn " + obj.Name + ", nil\n"

		fContent := declarative + metadata + statement + query + check + ret
		f.SetContent(fContent)
		if err = p.AddFunction(f); err != nil {
			return err
		}
	}
	return nil
}

func (p *PSQLBuilder) buildRead() error {
	var err error

	for _, obj := range p.objs {
		f := readFunction(driverName, varDriver, typeDriver, obj)
		var params string
		var values string
		var queryValues string
		for i, f := range obj.Fields {
			if i > 0 {
				values += ", "
				params += ", "
				//queryValues += ", "
			}
			paramName := strings.ToLower(f.Name)
			params += paramName
			values += paramName + " = ?"
			values += "$" + strconv.Itoa(i+1)
			queryValues += obj.Name + "." + f.Name
		}

		params += ", id, version, createDate, updateDate"
		nbFields := len(obj.Fields)
		values += ", $" + strconv.Itoa(nbFields+1) + ", $" + strconv.Itoa(nbFields+2) + ", $" + strconv.Itoa(nbFields+3)
		queryValues += ", " + obj.Name + ".Version, " + obj.Name + ".CreateDate, " + obj.Name + ".UpdateDate"

		declarative := "\tvar id int\n"
		declarative += "\tvar sqlStatement string\n"
		declarative += "\tvar err error\n"
		metadata := "\n\t" + obj.Name + ".Init()\n"
		statement := "\tsqlStatement = `INSERT INTO " + obj.TableName + " (" + params + ") VALUES (" + values + ") RETURNING id`\n"
		query := "\terr = db.QueryRow(sqlStatement, " + queryValues + ").Scan(&id)\n"
		check := "\tif err != nil {\n"
		check += "\t\treturn " + obj.Type + "{}, err\n"
		check += "\t}\n"
		ret := "\t" + obj.Name + ".id = id\n"
		ret += "\treturn " + obj.Name + ", nil\n"

		fContent := declarative + metadata + statement + query + check + ret
		f.SetContent(fContent)
		if err = p.AddFunction(f); err != nil {
			return err
		}
	}
	return nil

	/*
		var (
			id int
			name string
		)
		rows, err := db.Query("select id, name from users where id = ?", 1)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&id, &name)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(id, name)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	*/
	return nil
}

func (p *PSQLBuilder) buildUpdate() error {
	return nil
}

func (p *PSQLBuilder) buildDelete() error {
	var err error

	for _, obj := range p.objs {
		f := deleteFunction(driverName, varDriver, typeDriver, obj)

		declarative := "\tvar sqlStatement string\n"
		declarative += "\tvar err error\n"
		statement := "\n\tsqlStatement = `DELETE FROM " + obj.TableName + " WHERE id = $1`\n"
		query := "\t_, err = db.Exex(sqlStatement, " + obj.Name + ".ID)\n"
		check := "\tif err != nil {\n"
		check += "\t\treturn " + obj.Type + "{}, err\n"
		check += "\t}\n"
		ret := "\treturn " + obj.Name + ", nil\n"

		fContent := declarative + statement + query + check + ret
		f.SetContent(fContent)
		if err = p.AddFunction(f); err != nil {
			return err
		}
	}
	return nil
}

func (p PSQLBuilder) source() *gosource.GoSource {
	return p.GoSource
}
