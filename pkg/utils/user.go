package util

import (
	"MyMall/config"
	"errors"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	pictureIndexMap = map[string]struct{}{
		".jpg": {},
		".bmp": {},
		".png": {},
		".svg": {},
	}
)

// IsPicture 判断文件类型是否为图片
func IsPicture(ext string) bool {
	_, ok := pictureIndexMap[ext]
	return ok
}

// DirExistOrNot 判断文件夹是否存在
func DirExistOrNot(fileAddr string) bool {
	status, err := os.Stat(fileAddr)
	if err != nil {
		return false
	}
	return status.IsDir()
}

// CreateDir 新建文件夹
func CreateDir(dirName string) error {
	return os.MkdirAll(dirName, 755)
}

func UploadAvatarToLocalStatic(userId uint, userName string, file multipart.File, fileName string) (string, error) {
	ext := strings.ToLower(filepath.Ext(fileName))
	if !IsPicture(ext) {
		return "", errors.New("图片格式错误")
	}
	//相对路径 要加.
	basePath := filepath.Join(".", config.AvatarPath, "user"+strconv.Itoa(int(userId)))
	newFileName := userName + ext
	filePath := filepath.Join(basePath, newFileName)
	if !DirExistOrNot(basePath) {
		_ = CreateDir(basePath)
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(filePath, content, 0666)
	if err != nil {
		return "", err
	}
	return "user" + strconv.Itoa(int(userId)) + "/" + newFileName, nil
}
