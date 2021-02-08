package core

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sql2go"
	"strings"
)

type CmdEntity struct {
	No  string
	Msg string
}

var CmdHelp = []CmdEntity{
	{"0", "Set build directory."},
	{"1", "Generate the table markdown document."},
	{"2", "Generate table structure model."},
	{"3", "Generate table CURD."},
	{"c, clear", "Clear the screen"},
	{"h, help", "Show help list"},
	{"q, quit", "Quit"},
}

type Cmd struct {
}

func (c *Cmd) Handlers() map[string]func(s2g *S2G, args ...interface{}) int {
	return map[string]func(s2g *S2G, args ...interface{}) int{
		"0": c.customDir,
		"1": c.markDown,
		"2": c.model,
		"c": c.clean,
		"h": c.help,
		"q": c.quit,
	}
}

func (c *Cmd) customDir(s2g *S2G, args ...interface{}) int {
	fmt.Print("Please set the build directory>")
	line, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	if string(line) != "" {
		path, err := CheckOrCreateDir(string(line))
		if err == nil {
			s2g.OutPath = path
			fmt.Println("Directory success:", path)
		} else {
			log.Println("Set directory failed>>", err)
		}
	}
	return 0
}

//生成数据库表的markdown文档
func (c *Cmd) markDown(s2g *S2G, args ...interface{}) int {
	fmt.Println("Preparing to generate the markdown document...")
	//检查目录是否存在
	createDir(s2g.OutPath)
	err := s2g.CreateMarkdown()
	if err != nil {
		log.Println("MarkDown>>", err)
	}
	return 0
}

//set struct format
func (c *Cmd) _setFormat() []string {
	fmt.Print("Set the mapping name of the structure, separated by a comma (example :json,gorm)>")
	input, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	if string(input) != "" {
		formatList := sql2go.CheckCharDoSpecialArr(string(input), ',', `[\w\,\-]+`)
		if len(formatList) > 0 {
			fmt.Printf("Set format success: %v\n", formatList)
			return formatList
		}
	}
	fmt.Println("Set failed")
	return nil
}

//model
func (c *Cmd) model(s2g *S2G, args ...interface{}) int {
	var formats []string
	fmt.Print("Do you need to set the format of the structure?(yes/y/Yes|No/n)>")
	line, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	switch strings.ToLower(string(line)) {
	case "yes", "y":
		formats = c._setFormat()
	}
	err := s2g.createEntity(formats)
	if err != nil {
		log.Println("GenerateEntry>>", err)
	}
	go sql2go.Gofmt(sql2go.GetExeRootDir())
	return 0
}

func (c *Cmd) Help() {
	c.help(nil, nil)
}

// help
func (c *Cmd) help(s2g *S2G, args ...interface{}) int {
	for _, row := range CmdHelp {
		s := fmt.Sprintf("%s %s\n", "NO:"+row.No, row.Msg)
		fmt.Print(s)
	}
	return 0
}

// 清屏
func (c *Cmd) clean(s2g *S2G, args ...interface{}) int {
	sql2go.Clean()
	return 0
}

//退出
func (c *Cmd) quit(s2g *S2G, args ...interface{}) int {
	return 1
}
