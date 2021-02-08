package core

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"path"
	"sql2go"
	"sql2go/database"
	"sql2go/gen"
	"strings"
)

// 表字段信息
type FieldsInfo struct {
	TableDesc    database.TableDesc
	FormatFields string
	DbOriField   string
}

// 生成实体的请求结构
type EntityReq struct {
	Index        int    // 序列号
	TableName    string // 表名称
	TableNameGo  string // 表名称
	TableComment string // 表注释
	Path         string // 文件路径
	EntityPath   string // 实体路径
	Pkg          string // 命名空间名称
	EntityPkg    string // entity实体的空间名称
	FormatList   []string
	TableDesc    []*FieldsInfo
}

// 生成结构实体文件
func (s2g *S2G) createEntity(formatList []string) (err error) {
	// 表结构文件路径
	createDir(path.Join(s2g.OutPath, GODIR_Model))
	filePath := path.Join(s2g.OutPath, GODIR_Model, GOFILE_ENTITY)
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
		req.TableComment = table.Comment
		req.TableDesc = fieldsInfos
		req.FormatList = formatList
		req.EntityPkg = PkgEntity
		reqs = append(reqs, *req)
	}

	// 生成基础信息
	err = s2g.generateDBEntity(reqs, filePath)
	if err != nil {
		log.Fatal("CreateEntityErr>>", err)
	}

	return
}

// 创建结构实体
func (s2g *S2G) generateDBEntity(req []EntityReq, filePath string) (err error) {

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
