package service

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

// UploadToLocalStatic (avatar, userid, item_name, xxx.jpg, file)
// path = static/img/avatar/userid/username.ext
func UploadToLocalStatic(category string, id uint, itemName string, fileName string, file multipart.File) (path string, err error) {
	ext := strings.ToLower(filepath.Ext(fileName))
	if !IsPicture(ext) {
		return "", errors.New("图片格式有误")
	}
	selectCategory := func(name string) string {
		if name == "avatar" {
			return config.AvatarPath
		} else if name == "product" {
			return config.ProductPath
		}
		return ""
	}
	//相对路径 要加.
	basePath := filepath.Join(".", selectCategory(category), strconv.Itoa(int(id)))
	filePath := filepath.Join(basePath, itemName+ext)
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
	return strconv.Itoa(int(id)) + "/" + itemName + ext, nil
}
