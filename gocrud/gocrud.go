package gocrud

import (
	"time"
)

// GoCRUD describe metadata to the specific bdd object
type GoCRUD struct {
	ID         string
	Version    string
	CreateDate time.Time
	UpdateDate time.Time
}

// Init the metadata
func (g *GoCRUD) Init() {
	g.Version = "0"
	g.CreateDate = time.Now()
	g.UpdateDate = time.Now()
}
