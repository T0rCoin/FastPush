package logs

import (
	"io"
	"os"
	"time"
)

const (
	LOGPATH = "log/"
	FORMAT = "20060102"
	LineFeed = "\r\n"
)

var path = LOGPATH + time.Now().Format(FORMAT) + "/"

func WriteLog(fileName, msg string) error {
	if !IsExist(path) {
		return CreateDir(path)
	}
	var (
		err error
		f   *os.File
	)

	f, err = os.OpenFile(path+fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	_, err = io.WriteString(f, LineFeed+msg)

	defer f.Close()
	return err
}

func CreateDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	os.Chmod(path, os.ModePerm)
	return nil
}

func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}
