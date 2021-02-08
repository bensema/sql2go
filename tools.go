package sql2go

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

const (
	Unknown = iota
	Darwin
	Window
	Linux
)

//清屏
func Clean() {
	switch GetOs() {
	case Darwin, Linux:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case Window:
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

//获取操作系统
func GetOs() int {
	switch runtime.GOOS {
	case "darwin":
		return Darwin
	case "windows":
		return Window
	case "linux":
		return Linux
	default:
		return Unknown
	}
}

func Capitalize(s string) string {
	var upperStr string
	chars := strings.Split(s, "_")
	for _, val := range chars {
		upperStr += strings.ToUpper(val[0:1]) + val[1:]
	}
	return upperStr
}

func CheckCharDoSpecialArr(s string, char byte, reg string) []string {
	s = CheckCharDoSpecial(s, char, reg)
	return strings.Split(s, string(char))
}

//检查字符串,去掉特殊字符
func CheckCharDoSpecial(s string, char byte, regs string) string {
	reg := regexp.MustCompile(regs)
	var result []string
	if arr := reg.FindAllString(s, -1); len(arr) > 0 {
		buf := bytes.Buffer{}
		for key, val := range arr {
			if val != string(char) {
				buf.WriteString(val)
			}
			if val == string(char) && buf.Len() > 0 {
				result = append(result, buf.String())
				buf.Reset()
			}
			//处理最后一批数据
			if buf.Len() > 0 && key == len(arr)-1 {
				result = append(result, buf.String())
			}
		}
	}
	return strings.Join(result, string(char))
}

//FMT 格式代码
func Gofmt(path string) bool {
	if IsDirOrFileExist(path) {
		if !ExecCommand("goimports", "-l", "-w", path) {
			if !ExecCommand("gofmt", "-l", "-w", path) {
				return ExecCommand("go", "fmt", path)
			}
		}
		return true
	}
	return false
}

//判断文件 或 目录是否存在
func IsDirOrFileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

//GetRootDir 获取执行路径
func GetExeRootDir() string {
	// 文件不存在获取执行路径
	file, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		file = fmt.Sprintf(".%s", string(os.PathSeparator))
	} else {
		file = fmt.Sprintf("%s%s", file, string(os.PathSeparator))
	}
	return file
}

//添加注释 //
func AddToComment(s string, suff string) string {
	if strings.EqualFold(s, "") {
		return ""
	}
	return "// " + s + suff
}

//拼接特殊字符串
func FormatField(field string, formats []string) string {
	if len(formats) == 0 {
		return ""
	}
	buf := bytes.Buffer{}
	for key := range formats {
		buf.WriteString(fmt.Sprintf(`%s:"%s" `, formats[key], field))
	}
	return "`" + strings.TrimRight(buf.String(), " ") + "`"
}

func ExecCommand(name string, args ...string) bool {
	cmd := exec.Command(name, args...)
	_, err := cmd.Output()
	if err != nil {
		return false
	}
	return true
}
