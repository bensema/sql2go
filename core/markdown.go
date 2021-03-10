package core

import (
	"bytes"
	"fmt"
	"github.com/bensema/sql2go/database"
	"github.com/bensema/sql2go/gen"
	"log"
	"strings"
	"text/template"
	"time"
)

// markdown
type MarkDownData struct {
	TableList []*TableList
	DescList  []*MarkDownDataChild
}

// 表名和表注释
type TableList struct {
	Index          int
	UpperTableName string
	TableName      string
	Comment        string
}

type MarkDownDataChild struct {
	Index     int    // 自增
	TableName string // 表名
	Comment   string // 表备注
	List      []*database.TableDesc
}

// 生成mysql markdown文档
func (s2g *S2G) CreateMarkdown() (err error) {
	data := new(MarkDownData)
	// 将表结构写入文件
	i := 1
	tables, err := s2g.Db.FindTables()
	if err != nil {
		fmt.Println(err)
	}
	for _, table := range tables {
		fmt.Println("Doing table:" + table.Name)
		data.TableList = append(data.TableList, &TableList{
			Index:          i,
			UpperTableName: strings.ToUpper(table.Name),
			TableName:      table.Name,
			Comment:        table.Comment,
		})
		// 查询表结构信息
		desc := new(MarkDownDataChild)
		desc.List, err = s2g.Db.GetTableColumns(table.Name)
		if err != nil {
			log.Fatal("markdown>>", err)
			continue
		}
		desc.Index = i
		desc.TableName = table.Name
		desc.Comment = table.Comment
		data.DescList = append(data.DescList, desc)
		i++
	}

	// 生成所有表的文件
	err = s2g.generateMarkdown(data)
	if err != nil {
		return
	}
	return
}

// 生成表列表
func (s2g *S2G) generateMarkdown(data *MarkDownData) (err error) {
	// 写入markdown
	file := s2g.OutPath + fmt.Sprintf("markdown%s.md", time.Now().Format("2006-01-02_150405"))
	tplByte, err := gen.Asset(gen.TplMarkdown)
	if err != nil {
		return
	}
	// 解析
	content := bytes.NewBuffer([]byte{})
	tpl, err := template.New("markdown").Parse(string(tplByte))
	err = tpl.Execute(content, data)
	if err != nil {
		return
	}
	// 表信息写入文件
	err = WriteAppendFile(file, content.String())
	if err != nil {
		return
	}
	return
}
