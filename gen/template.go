package gen

import "github.com/bensema/sql2go/gen/internal"

const (
	TplMarkdown   = "template/markdown.tmpl"
	TplModel      = "template/model.tmpl"
	TplModelReq   = "template/model_req.tmpl"
	TplModelReply = "template/model_reply.tmpl"
	TplModelPage  = "template/page.tmpl"
	TplBiz        = "template/biz.tmpl"
	TplService    = "template/service.tmpl"
	TplOpCode     = "template/op_code.tmpl"
	TplHttpBiz    = "template/http_biz.tmpl"
	TplAdminCmd   = "template/admin_cmd.tmpl"
	TplCurd       = "template/curd.tmpl"
	TplCurd1      = "template/curd_1.tmpl"
	TplCurdCommon = "template/curd_common.tmpl"
)

func Asset(name string) ([]byte, error) {
	return internal.Asset(name)
}
