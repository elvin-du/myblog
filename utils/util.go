package utils

import(
	"runtime"
	"log"
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
