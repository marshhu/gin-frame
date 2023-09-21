package utils

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
)

// FileExists 判断文件是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// IsDir 是否是目录
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile 是否是文件
func IsFile(path string) bool {
	return !IsDir(path)
}

// RootDir 应用所在根目录
func RootDir() string {
	wd, _ := os.Getwd()
	return wd
}

// GetExt 获取文件后缀
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// ReadFile 读取文件 ReadFile
func ReadFile(filePath string) ([]byte, error) {
	fin, err := os.Open(filePath)
	defer fin.Close()
	if err != nil {
		return nil, err
	}
	return file2Bytes(fin)
}

// file2Bytes
func file2Bytes(file *os.File) ([]byte, error) {
	defer file.Close()
	stats, err := file.Stat()
	if err != nil {
		return nil, err
	}
	data := make([]byte, stats.Size())
	count, err := file.Read(data)
	if err != nil {
		return nil, err
	}
	fmt.Printf("read file len: %d \n", count)
	return data, nil
}

func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}

func MkDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// 目录不存在，创建目录
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}

	}
	return nil
}
