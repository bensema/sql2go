package core

import (
	"github.com/bensema/sql2go/database"
)

const (
	GODIR_Model    = "model"    // entity file
	GODIR_Internal = "internal" // entity file
	GODIRService   = "service"  // entity file
	GODIRDao       = "dao"      // entity file
	GODIRHttp      = "http"     // entity file
	PkgModel       = "model"    // entity package name

)

type S2G struct {
	Db      *database.DB
	OutPath string
}
