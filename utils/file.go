package utils

import (
	"io"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
)

func GetSize(f multipart.File) (int, error) {
	content, err := io.ReadAll(f)

	return len(content), err
}

func GetExt(fileName string) string {
	return path.Ext(fileName)
}

func CheckNotExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

func IsNotExistMkDir(src string) error {
	if notExist := CheckNotExist(src); notExist == true {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

func MkDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// 目录不存在，创建目录
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}

	}
	return nil
}

func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// RootDir 应用所在根目录
func RootDir() string {
	wd, _ := os.Getwd()
	return wd
}

func CopyFile(sourcePath, destPath, fileName string) error {
	// 打开源文件
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()
	if err = MkDir(destPath); err != nil {
		return err
	}
	// 创建目标文件
	destFullPath := filepath.Join(destPath, fileName)
	destFile, err := os.Create(destFullPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// 将源文件拷贝到目标文件
	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}
	return nil
}
