package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ymohl-cl/gopkg/gocrud/crudgen/builder"
)

var (
	source = flag.String("source", "", "Model Go source file.")
	bdd    = flag.String("bdd", "", "bdd's driver (supported: psql).")
)

func init() {
	flag.Parse()
	if *source == "" {
		fmt.Println("source file can't be ignored")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *bdd == "" {
		fmt.Println("bdd driver can't be ignored")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func main() {
	var err error
	var b builder.Builder

	if b, err = builder.New(*bdd, *source); err != nil {
		panic(err)
	}
	if err = builder.Build(b); err != nil {
		panic(err)
	}
}
