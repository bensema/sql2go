package core

import (
	"bytes"
	"fmt"
	"github.com/bensema/sql2go"
	"github.com/bensema/sql2go/database"
	"log"
	"path"
	"strings"
	"text/template"
)

// 表字段信息
type FieldsInfo struct {
	TableDesc    database.TableDesc
	FormatFields string
	DbOriField   string
}

// 生成实体的请求结构
type EntityReq struct {
	Index         int    // 序列号
	TableName     string // 表名称
	TableNameGo   string // 表名称
	TableNameGoV2 string // 表名称 驼峰首字母小些
	TableComment  string // 表注释
	Path          string // 文件路径
	EntityPath    string // 实体路径
	Pkg           string // 命名空间名称
	EntityPkg     string // entity实体的空间名称
	FormatList    []string
	TableDesc     []*FieldsInfo
}

// 生成结构实体文件
func (s2g *S2G) createModel(formatList []string) (err error) {
	// 表结构文件路径
	createDir(path.Join(s2g.OutPath, ProjectBB))

	createDir(path.Join(s2g.OutPath, ProjectBB, GODIR_Model))
	createDir(path.Join(s2g.OutPath, ProjectBB, GODIRService))
	createDir(path.Join(s2g.OutPath, ProjectBB, GODIRServer))
	createDir(path.Join(s2g.OutPath, ProjectBB, GODIRDao))

	createDir(path.Join(s2g.OutPath, ProjectBB, GODIRServer, GODIRHttp))

	modelBizPath := path.Join(s2g.OutPath, ProjectBB, GODIR_Model, "model_biz.go")
	modelReqBizPath := path.Join(s2g.OutPath, ProjectBB, GODIR_Model, "model_req_biz.go")
	modelReplyBizPath := path.Join(s2g.OutPath, ProjectBB, GODIR_Model, "model_reply_biz.go")
	modelPageBizPath := path.Join(s2g.OutPath, ProjectBB, GODIR_Model, "page_biz.go")
	modelopCodeBizPath := path.Join(s2g.OutPath, ProjectBB, GODIR_Model, "op_code_biz.go")

	bizPath := path.Join(s2g.OutPath, ProjectBB, GODIRDao, "biz.go")
	serviceBizPath := path.Join(s2g.OutPath, ProjectBB, GODIRService, "service_biz.go")

	// 将表结构写入文件
	tables, err := s2g.Db.FindTables()
	if err != nil {
		fmt.Println(err)
	}
	reqs := []EntityReq{}
	for idx, table := range tables {
		idx++
		// 查询表结构信息
		tableDesc, err := s2g.Db.GetTableColumns(table.Name)
		if err != nil {
			log.Fatal("CreateEntityErr>>", err)
			continue
		}
		fieldsInfos := []*FieldsInfo{}
		for _, val := range tableDesc {
			fieldsInfos = append(fieldsInfos, &FieldsInfo{
				TableDesc:    *val,
				FormatFields: formatField(val.ColumnName, formatList),
				DbOriField:   val.ColumnName,
			})
		}

		req := new(EntityReq)
		req.Index = idx
		req.TableName = table.Name
		req.TableNameGo = sql2go.Capitalize(table.Name)
		req.TableNameGoV2 = sql2go.CapitalizeV2(table.Name)
		req.TableComment = table.Comment
		req.TableDesc = fieldsInfos
		req.FormatList = formatList
		req.EntityPkg = PkgModel
		reqs = append(reqs, *req)
	}

	err = s2g.GenCommon(reqs, modelBizPath, "content", TplModelBiz)
	if err != nil {
		log.Fatal("Create Model error>>", err)
	}

	err = s2g.GenCommon(reqs, modelReqBizPath, "content", TplModelReqBiz)
	if err != nil {
		log.Fatal("Create Model req error>>", err)
	}

	err = s2g.GenCommon(reqs, modelReplyBizPath, "content", TplModelReplyBiz)
	if err != nil {
		log.Fatal("Create Model reply error>>", err)
	}

	err = s2g.GenCommon(reqs, modelPageBizPath, "content", TplModelPageBiz)
	if err != nil {
		log.Fatal("Create Model page error>>", err)
	}
	err = s2g.GenCommon(reqs, bizPath, "content", TplDaoBiz, TplProject)
	if err != nil {
		log.Fatal("Create dao biz error>>", err)
	}

	err = s2g.GenCommon(reqs, serviceBizPath, "content", TplServiceBiz, TplProject)
	if err != nil {
		log.Fatal("Create service bizPath error>>", err)
	}

	err = s2g.GenCommon(reqs, modelopCodeBizPath, "content", TplModelOpCodeBiz)
	if err != nil {
		log.Fatal("Create op_code error>>", err)
	}
	return
}

func (s2g *S2G) GenCommon(req []EntityReq, savePath string, templateName string, filenames ...string) (err error) {
	tpl, err := template.ParseFiles(filenames...)
	if err != nil {
		fmt.Printf("GenCommon err: %s", err)
		return
	}

	content := bytes.NewBuffer([]byte{})
	err = tpl.ExecuteTemplate(content, templateName, req)
	if err != nil {
		fmt.Println(err)
	}
	// 表信息写入文件
	con := strings.Replace(content.String(), "&#34;", `"`, -1)
	err = WriteFile(savePath, con)
	if err != nil {
		return
	}
	return
}

//拼接特殊字符串
func formatField(field string, formats []string) string {
	if len(formats) == 0 {
		return ""
	}
	buf := bytes.Buffer{}
	for key := range formats {
		buf.WriteString(fmt.Sprintf(`%s:"%s" `, formats[key], field))
	}
	return "`" + strings.TrimRight(buf.String(), " ") + "`"
}
