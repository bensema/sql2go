package core

import (
	"github.com/bensema/sql2go/database"
)

const (
	GODIR_Model   = "model"        // entity file
	GOFILE_ENTITY = "db_entity.go" // entity table file
	PkgEntity     = "model"        // entity package name

)

type S2G struct {
	Db      *database.DB
	OutPath string
}
