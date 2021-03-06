package core

import (
	"bytes"
	"fmt"
	"github.com/bensema/sql2go"
	"log"
	"path"
	"strings"
	"text/template"
)

// 生成结构实体文件
func (s2g *S2G) createCurd(formatList []string) (err error) {
	// 表结构文件路径
	createDir(path.Join(s2g.OutPath, ProjectBB, GODIRDao))

	createDir(path.Join(s2g.OutPath, ProjectBB, GODIRDao, GODIR_Internal))

	curdCommomPath := path.Join(s2g.OutPath, ProjectBB, GODIRDao, GODIR_Internal, "curd_common_biz.go")
	req := []EntityReq{}
	s2g.GenCommon(req, curdCommomPath, "content", TplDaoCurdCommonBiz)
	// 将表结构写入文件
	tables, err := s2g.Db.FindTables()
	if err != nil {
		fmt.Println(err)
	}
	reqs := []EntityReq{}
	for idx, table := range tables {
		idx++
		filePath := path.Join(s2g.OutPath, ProjectBB, GODIRDao, GODIR_Internal, table.Name+"_curd_biz.go")
		// 查询表结构信息
		tableDesc, err := s2g.Db.GetTableColumns(table.Name)
		if err != nil {
			log.Fatal("Create Curd error >>", err)
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
		req.EntityPkg = "internal"
		reqs = append(reqs, *req)

		// 生成基础信息
		s2g.generateDBCurd(req, filePath, "content", TplDaoCurdBiz, TplProject)
		if err != nil {
			log.Fatal("Create Curd error >>", err)
		}

	}

	return
}

// 创建结构实体
func (s2g *S2G) generateDBCurd(req *EntityReq, savePath string, templateName string, filenames ...string) (err error) {

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
