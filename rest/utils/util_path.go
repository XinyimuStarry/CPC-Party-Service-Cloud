package utils

import (
	"os"
	"path/filepath"
)

func GetHomePath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}

// 判断文件或文件夹是否已经存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetLogPath() string {
	homePath := GetHomePath()
	filePath := homePath + "/log"
	exists, err := PathExists(filePath)
	if err != nil {
		panic("判断日志文件夹是否存在时出错！")
	}
	if !exists {
		err = os.MkdirAll(filePath, 0777)
		if err != nil {
			panic("创建日志文件夹时出错！")
		}
	}
	return filePath
}
