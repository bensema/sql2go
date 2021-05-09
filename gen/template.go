package gen

import "github.com/bensema/sql2go/gen/internal"

const (
	TplMarkdown = "template/markdown.tmpl"
	TplOpCode   = "template/op_code.tmpl"

	TplModelBiz      = "template/model/model_biz.tmpl"
	TplModelReqBiz   = "template/model/model_req_biz.tmpl"
	TplModelReplyBiz = "template/model/model_reply_biz.tmpl"
	TplModelPageBiz  = "template/model/page_biz.tmpl"
	TplServiceBiz    = "template/service/service_biz.tmpl"
	TplBiz           = "template/dao/biz.tmpl"
	TplCurdBiz       = "template/dao/internal/curd_biz.tmpl"
	TplCurdCommonBiz = "template/dao/internal/curd_common.tmpl"
)

func Asset(name string) ([]byte, error) {
	return internal.Asset(name)
}
