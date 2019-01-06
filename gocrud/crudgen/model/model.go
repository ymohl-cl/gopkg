package model

import (
	"github.com/ymohl-cl/gopkg/gocrud"
)

// User Documentation
type User struct {
	gocrud.GoCRUD
	Pseudo string `crud:"uniq"`
	Age    string `crud:"uniq"`
	Email  string `crud:"uniq,auto" json:"omitempty"`
}

// Truc to test
type Truc struct {
	gocrud.GoCRUD
	Test string
}
