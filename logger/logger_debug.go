/*
负责debug log的输出
*/

package logger

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"myblog/utils"
)

func Debugf(format string, args ...interface{}) {
	logName := utils.GetCurrentTime() + ".debug"
	file, err := os.OpenFile(debugFileDir+logName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
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
func Debugln(args ...interface{}) {
	logName := utils.GetCurrentTime() + ".debug"
	file, err := os.OpenFile(debugFileDir+logName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	// 加上文件调用和行号
	_, callerFile, line, ok := runtime.Caller(1)
	if ok {
		args = append([]interface{}{filepath.Base(callerFile), ":", line}, args...)
	}
	New(file).Println(args...)
}
