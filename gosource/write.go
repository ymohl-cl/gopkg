package gosource

import (
	"fmt"
	"go/format"
	"os"
)

// Bytes return the content source file to the GoSource.
// If GoSource don't describe a go file, err is return and byte will be empty.
func (g *GoSource) Bytes() ([]byte, error) {
	var content []byte
	var err error

	g.printHeader()
	if len(g.imports) > 0 {
		g.printImport()
	}
	if len(g.consts) > 0 {
		g.printConst()
	}
	if len(g.globals) > 0 {
		g.printGlobal()
	}
	g.printFunction()

	if content, err = format.Source(g.buf.Bytes()); err != nil {
		fmt.Println(string(g.buf.Bytes()))
		return nil, err
	}
	return content, nil
}

// CreateFile write the source in the new file targetted by the filepath.
// if an error occured, it will be returned.
func (g *GoSource) CreateFile(filepath string) error {
	var err error
	var f *os.File
	var b []byte

	if f, err = os.Create(filepath); err != nil {
		return err
	}
	defer f.Close()
	if b, err = g.Bytes(); err != nil {
		return err
	}
	if _, err = f.Write(b); err != nil {
		return err
	}
	return nil
}

// writeN with a line return at the end of print.
func (g *GoSource) writeN(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format+"\n", args...)
}

// write without a line return at the end of print.
func (g *GoSource) write(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}

// printHeader section.
func (g *GoSource) printHeader() {
	for i, c := range g.comments {
		if i == 0 {
			g.writeN("// Package %s %s", g.name, c)
		} else {
			g.writeN("// %s", c)
		}
	}
	g.writeN("package %s", g.name)
	g.writeN("")
}

// printImport section.
func (g *GoSource) printImport() {
	var nb int

	nb = len(g.imports)
	if nb == 0 {
		return
	}
	if nb == 1 {
		g.writeN("import \"%s\"", g.imports[0])
		return
	}

	g.writeN("import (")
	for _, i := range g.imports {
		g.writeN("\t\"%s\"", i)
	}
	g.writeN(")")
	g.writeN("")
}

// printConst section.
func (g *GoSource) printConst() {
	g.writeN("const (")
	for name, t := range g.consts {
		if t.name == "string" {
			g.writeN("\t%s %s = \"%v\"", name, t.name, t.value)
		} else {
			g.writeN("\t%s %s = %v", name, t.name, t.value)
		}
	}
	g.writeN(")")
	g.writeN("")
}

// printGlobal section.
func (g *GoSource) printGlobal() {
	for name, t := range g.globals {
		if t.name == "string" {
			g.writeN("var %s %s = \"%v\"", name, t.name, t.value)
		} else {
			g.writeN("var %s %s = %v", name, t.name, t.value)
		}
	}
	g.writeN("")
}

// printFunction section.
func (g *GoSource) printFunction() {
	for _, f := range g.funcs {
		g.write("func %s(", f.name)
		first := true
		for _, p := range f.args {
			if !first {
				g.write(", ")
			}
			g.write("%s %s", p.Name, p.Type)
			first = false
		}
		if len(f.rets) > 0 {
			g.write(") (")
			first = true
			for _, p := range f.rets {
				if !first {
					g.write(", ")
				}
				g.write("%s %s", p.Name, p.Type)
				first = false
			}
		}
		g.writeN(") {")
		g.write(f.content)
		g.writeN("}")
		g.writeN("")
	}
}
