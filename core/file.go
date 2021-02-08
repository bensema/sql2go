package core

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func CheckOrCreateDir(path string) (string, error) {
	if len(path) == 0 {
		return "", errors.New("directory is null")
	}
	last := path[len(path)-1:]
	if !strings.EqualFold(last, string(os.PathSeparator)) {
		path = path + string(os.PathSeparator)
	}
	if !isDir(path) {
		if createDir(path) {
			return path, nil
		}
		return "", errors.New(path + "Failed to create or insufficient permissions")
	}
	return path, nil
}

func isDir(filename string) bool {
	return isFileOrDir(filename, true)
}

func isFileOrDir(filename string, decideDir bool) bool {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return false
	}
	isDir := fileInfo.IsDir()
	if decideDir {
		return isDir
	}
	return !isDir
}

//创建目录
func createDir(path string) bool {
	if isDirOrFileExist(path) == false {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return false
		}
	}
	return true
}

//判断文件 或 目录是否存在
func isDirOrFileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

//追加写文件
func WriteAppendFile(path, data string) (err error) {
	if _, err := WriteFileAppend(path, data); err == nil {
		fmt.Printf("Generate success:%s\n", path)
		return nil
	} else {
		return err
	}
}

//追加写文件
func WriteFileAppend(filename string, data string) (count int, err error) {
	var f *os.File
	if isDirOrFileExist(filename) == false {
		f, err = os.Create(filename)
		if err != nil {
			return
		}
	} else {
		f, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0666)
	}
	defer f.Close()
	count, err = io.WriteString(f, data)
	if err != nil {
		return
	}
	return
}

// 写文件
func WriteFile(path, data string) (err error) {
	if _, err := writeFile(path, data); err == nil {
		fmt.Printf("Generate success:%s\n", path)
		return nil
	} else {
		return err
	}
}

func writeFile(filename string, data string) (count int, err error) {
	var f *os.File
	if isDirOrFileExist(filename) == false {
		f, err = os.Create(filename)
		if err != nil {
			return
		}
	} else {
		f, err = os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0600)
	}
	defer f.Close()
	count, err = io.WriteString(f, data)
	if err != nil {
		return
	}
	return
}
