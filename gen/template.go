package gen

import "github.com/bensema/sql2go/gen/internal"

const (
	TplMarkdown = "template/markdown.tmpl"
	TplModel    = "template/model.tmpl"
)

func Asset(name string) ([]byte, error) {
	return internal.Asset(name)
}
