package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

const (
	BaseHtmlTplFile     = "views/common/base.html"
	Error404HtmlTplFile = "views/common/404.html"
)

func CheckError(err error) {
	if nil != err {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			log.Print(file, ":", line)
		}
		log.Println(err)
	}
}

// 获得可执行程序所在目录
func ExeDir() (string, error) {
	pathAbs, err := filepath.Abs(os.Args[0])
	if err != nil {
		return "", err
	}
	//Dir returns all but the last element of path,and the trailing slashes are removed
	return filepath.Dir(pathAbs), nil
}

//得到现在时刻的年，月，日，小时，并转化为字符串形式返回
func GetCurrentTime() string {
	t := time.Now()
	y, m, d := t.Date()
	h := t.Hour()
	//月，天，小时都必须为占两位数，不够则补零
	return fmt.Sprintf("%d%02d%02d%02d", y, m, d, h)
}

/*
template模板函数：加一
*/
func Plus(args ...interface{}) int {
	ok := false
	var pageId int
	if len(args) == 1 {
		pageId, ok = args[0].(int)
	}
	if !ok {
		return 0
	}

	if pageId == 0 {
		pageId = 2
	} else {
		pageId = pageId + 1
	}

	return pageId
}
