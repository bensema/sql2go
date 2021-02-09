package core

import (
	"github.com/bensema/sql2go/database"
)

const (
	GODIR_Model    = "model"    // entity file
	GODIR_Internal = "internal" // entity file
	PkgModel       = "model"    // entity package name

)

type S2G struct {
	Db      *database.DB
	OutPath string
}
