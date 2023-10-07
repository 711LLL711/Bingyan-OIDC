package utils

import (
	"errors"
	"os"
	"path"
	"strconv"
	"time"
)

func GenerateImgName(file string) (string, error) {
	//判断文件类型合法
	extname := path.Ext(file)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}
	if _, ok := allowExtMap[extname]; !ok {
		return "", errors.New("文件类型不合法")
	}
	//创建文件目录
	//uploadDir := "upload/" + time.Now().Format("2006/01/02/")
	uploadDir := "upload/"
	os.MkdirAll(uploadDir, os.ModePerm)
	//生成文件名
	filename := strconv.FormatInt(time.Now().Unix(), 10) + extname
	//返回文件名
	return uploadDir + filename, nil
}
