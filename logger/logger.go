/*
负责log的输出
*/

package logger

import (
	"io"
	"log"
	"myblog/config"
	"myblog/utils"
	"os"
)

var (
	//日志目录
	infoFileDir  = config.ROOT + "/logs/info/"
	debugFileDir = config.ROOT + "/logs/debug/"
	errorFileDir = config.ROOT + "/logs/error/"
)

func init() {
	os.MkdirAll(infoFileDir, os.ModeDir)
	os.MkdirAll(debugFileDir, os.ModeDir)
	os.MkdirAll(errorFileDir, os.ModeDir)
}

type logger struct {
	*log.Logger
}

func New(out io.Writer) *logger {
	return &logger{
		Logger: log.New(out, "", log.LstdFlags),
	}
}

func Infof(format string, args ...interface{}) {
	//create log file name by current time
	logName := utils.GetCurrentTime() + ".info"
	file, err := os.OpenFile(infoFileDir+logName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	New(file).Printf(format, args...)
}

/*
输出自动换行
*/
func Infoln(args ...interface{}) {
	logName := utils.GetCurrentTime() + ".info"
	file, err := os.OpenFile(infoFileDir+logName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	New(file).Println(args...)
}

func Errorf(format string, args ...interface{}) {
	logName := utils.GetCurrentTime() + ".error"
	file, err := os.OpenFile(errorFileDir+logName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	New(file).Printf(format, args...)
}

/*
输出自动换行
*/
func Errorln(args ...interface{}) {
	logName := utils.GetCurrentTime() + ".error"
	file, err := os.OpenFile(errorFileDir+logName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	New(file).Println(args...)
}
