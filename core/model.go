package core

import (
	"bytes"
	"fmt"
	"github.com/bensema/sql2go"
	"github.com/bensema/sql2go/database"
	"github.com/bensema/sql2go/gen"
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
	createDir(path.Join(s2g.OutPath, GODIR_Model))
	createDir(path.Join(s2g.OutPath, GODIRService))
	createDir(path.Join(s2g.OutPath, GODIRHttp))
	createDir(path.Join(s2g.OutPath, GODIRDao))
	filePath := path.Join(s2g.OutPath, GODIR_Model, "model.go")
	filePathModelReq := path.Join(s2g.OutPath, GODIR_Model, "model_req.go")
	filePathModelReply := path.Join(s2g.OutPath, GODIR_Model, "model_reply.go")
	filePathModelPage := path.Join(s2g.OutPath, GODIR_Model, "page.go")
	biz := path.Join(s2g.OutPath, GODIRDao, "biz.go")
	serviceBiz := path.Join(s2g.OutPath, GODIRService, "service_biz.go")
	opCode := path.Join(s2g.OutPath, GODIR_Model, "op_code_biz.go")
	httpBiz := path.Join(s2g.OutPath, GODIRHttp, "http_biz.go")
	admCmd := path.Join(s2g.OutPath, GODIRHttp, "adm_cmd.go")
	bbAdminApiBiz := path.Join(s2g.OutPath, GODIRHttp, "bb_admin_api_biz.go")
	bbAdminBiz := path.Join(s2g.OutPath, GODIRHttp, "bb_admin_biz.go")
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

	// 生成基础信息
	err = s2g.generateModel(reqs, filePath)
	if err != nil {
		log.Fatal("Create Model error>>", err)
	}
	err = s2g.generateModelReq(reqs, filePathModelReq)
	if err != nil {
		log.Fatal("Create Model req error>>", err)
	}
	err = s2g.generateModelReply(reqs, filePathModelReply)
	if err != nil {
		log.Fatal("Create Model reply error>>", err)
	}
	err = s2g.generateModelPage(reqs, filePathModelPage)
	if err != nil {
		log.Fatal("Create Model page error>>", err)
	}

	err = s2g.generateModelBiz(reqs, biz)
	if err != nil {
		log.Fatal("Create biz error>>", err)
	}

	err = s2g.generateModelServiceBiz(reqs, serviceBiz)
	if err != nil {
		log.Fatal("Create service biz error>>", err)
	}

	err = s2g.generateModelOpCode(reqs, opCode)
	if err != nil {
		log.Fatal("Create op_code error>>", err)
	}

	err = s2g.generateModelHttpBiz(reqs, httpBiz)
	if err != nil {
		log.Fatal("Create http biz error>>", err)
	}

	err = s2g.generateModelAdmCmd(reqs, admCmd)
	if err != nil {
		log.Fatal("Create admin cmd error>>", err)
	}

	err = s2g.generateModelBBAdminBiz(reqs, bbAdminBiz)
	if err != nil {
		log.Fatal("Create bb admin biz error>>", err)
	}

	err = s2g.generateModelBBAdminApiBiz(reqs, bbAdminApiBiz)
	if err != nil {
		log.Fatal("Create bb admin api biz error>>", err)
	}

	return
}

// 创建结构实体
func (s2g *S2G) generateModel(req []EntityReq, filePath string) (err error) {

	// 加载模板文件
	tplByte, err := gen.Asset(gen.TplModel)
	if err != nil {
		return
	}
	tpl, err := template.New("model").Parse(string(tplByte))
	if err != nil {
		return
	}

	content := bytes.NewBuffer([]byte{})
	err = tpl.Execute(content, req)
	if err != nil {
		fmt.Println(err)
	}
	// 表信息写入文件
	con := strings.Replace(content.String(), "&#34;", `"`, -1)
	err = WriteFile(filePath, con)
	if err != nil {
		return
	}
	return
}

func (s2g *S2G) generateModelReq(req []EntityReq, filePath string) (err error) {

	// 加载模板文件
	tplByte, err := gen.Asset(gen.TplModelReq)
	if err != nil {
		return
	}
	tpl, err := template.New("model_req").Parse(string(tplByte))
	if err != nil {
		return
	}

	content := bytes.NewBuffer([]byte{})
	err = tpl.Execute(content, req)
	if err != nil {
		fmt.Println(err)
	}
	// 表信息写入文件
	con := strings.Replace(content.String(), "&#34;", `"`, -1)
	err = WriteFile(filePath, con)
	if err != nil {
		return
	}
	return
}

func (s2g *S2G) generateModelReply(req []EntityReq, filePath string) (err error) {

	// 加载模板文件
	tplByte, err := gen.Asset(gen.TplModelReply)
	if err != nil {
		return
	}
	tpl, err := template.New("model_reply").Parse(string(tplByte))
	if err != nil {
		return
	}

	content := bytes.NewBuffer([]byte{})
	err = tpl.Execute(content, req)
	if err != nil {
		fmt.Println(err)
	}
	// 表信息写入文件
	con := strings.Replace(content.String(), "&#34;", `"`, -1)
	err = WriteFile(filePath, con)
	if err != nil {
		return
	}
	return
}

func (s2g *S2G) generateModelPage(req []EntityReq, filePath string) (err error) {

	// 加载模板文件
	tplByte, err := gen.Asset(gen.TplModelPage)
	if err != nil {
		return
	}
	tpl, err := template.New("page").Parse(string(tplByte))
	if err != nil {
		return
	}

	content := bytes.NewBuffer([]byte{})
	err = tpl.Execute(content, req)
	if err != nil {
		fmt.Println(err)
	}
	// 表信息写入文件
	con := strings.Replace(content.String(), "&#34;", `"`, -1)
	err = WriteFile(filePath, con)
	if err != nil {
		return
	}
	return
}

func (s2g *S2G) generateModelBiz(req []EntityReq, filePath string) (err error) {

	// 加载模板文件
	tplByte, err := gen.Asset(gen.TplBiz)
	if err != nil {
		return
	}
	tpl, err := template.New("page").Parse(string(tplByte))
	if err != nil {
		return
	}

	content := bytes.NewBuffer([]byte{})
	err = tpl.Execute(content, req)
	if err != nil {
		fmt.Println(err)
	}
	// 表信息写入文件
	con := strings.Replace(content.String(), "&#34;", `"`, -1)
	err = WriteFile(filePath, con)
	if err != nil {
		return
	}
	return
}

func (s2g *S2G) generateModelServiceBiz(req []EntityReq, filePath string) (err error) {

	// 加载模板文件
	tplByte, err := gen.Asset(gen.TplService)
	if err != nil {
		return
	}
	tpl, err := template.New("page").Parse(string(tplByte))
	if err != nil {
		return
	}

	content := bytes.NewBuffer([]byte{})
	err = tpl.Execute(content, req)
	if err != nil {
		fmt.Println(err)
	}
	// 表信息写入文件
	con := strings.Replace(content.String(), "&#34;", `"`, -1)
	err = WriteFile(filePath, con)
	if err != nil {
		return
	}
	return
}

func (s2g *S2G) generateModelOpCode(req []EntityReq, filePath string) (err error) {

	// 加载模板文件
	tplByte, err := gen.Asset(gen.TplOpCode)
	if err != nil {
		return
	}
	tpl, err := template.New("page").Parse(string(tplByte))
	if err != nil {
		return
	}

	content := bytes.NewBuffer([]byte{})
	err = tpl.Execute(content, req)
	if err != nil {
		fmt.Println(err)
	}
	// 表信息写入文件
	con := strings.Replace(content.String(), "&#34;", `"`, -1)
	err = WriteFile(filePath, con)
	if err != nil {
		return
	}
	return
}

func (s2g *S2G) generateModelHttpBiz(req []EntityReq, filePath string) (err error) {

	// 加载模板文件
	tplByte, err := gen.Asset(gen.TplHttpBiz)
	if err != nil {
		return
	}
	tpl, err := template.New("page").Parse(string(tplByte))
	if err != nil {
		return
	}

	content := bytes.NewBuffer([]byte{})
	err = tpl.Execute(content, req)
	if err != nil {
		fmt.Println(err)
	}
	// 表信息写入文件
	con := strings.Replace(content.String(), "&#34;", `"`, -1)
	err = WriteFile(filePath, con)
	if err != nil {
		return
	}
	return
}

func (s2g *S2G) generateModelAdmCmd(req []EntityReq, filePath string) (err error) {

	// 加载模板文件
	tplByte, err := gen.Asset(gen.TplAdminCmd)
	if err != nil {
		return
	}
	tpl, err := template.New("page").Parse(string(tplByte))
	if err != nil {
		return
	}

	content := bytes.NewBuffer([]byte{})
	err = tpl.Execute(content, req)
	if err != nil {
		fmt.Println(err)
	}
	// 表信息写入文件
	con := strings.Replace(content.String(), "&#34;", `"`, -1)
	err = WriteFile(filePath, con)
	if err != nil {
		return
	}
	return
}

func (s2g *S2G) generateModelBBAdminBiz(req []EntityReq, filePath string) (err error) {

	// 加载模板文件
	tplByte, err := gen.Asset(gen.TplBBAdminBiz)
	if err != nil {
		return
	}
	tpl, err := template.New("page").Parse(string(tplByte))
	if err != nil {
		return
	}

	content := bytes.NewBuffer([]byte{})
	err = tpl.Execute(content, req)
	if err != nil {
		fmt.Println(err)
	}
	// 表信息写入文件
	con := strings.Replace(content.String(), "&#34;", `"`, -1)
	err = WriteFile(filePath, con)
	if err != nil {
		return
	}
	return
}

func (s2g *S2G) generateModelBBAdminApiBiz(req []EntityReq, filePath string) (err error) {

	// 加载模板文件
	tplByte, err := gen.Asset(gen.TplBBAdminApiBiz)
	if err != nil {
		return
	}
	tpl, err := template.New("page").Parse(string(tplByte))
	if err != nil {
		return
	}

	content := bytes.NewBuffer([]byte{})
	err = tpl.Execute(content, req)
	if err != nil {
		fmt.Println(err)
	}
	// 表信息写入文件
	con := strings.Replace(content.String(), "&#34;", `"`, -1)
	err = WriteFile(filePath, con)
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
