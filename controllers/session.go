package controllers

import(
	"fmt"
	"time"
	"math/rand"
	"crypto/md5"
	"strconv"
	"io"
	"net/http"
)

var GSessionSet []string

func CreateSessionID()string{
    nano := time.Now().UnixNano()
    rand.Seed(nano)
    rndNum := rand.Int63()
    sessionId := Md5(Md5(strconv.FormatInt(nano, 10))+Md5(strconv.FormatInt(rndNum, 10)))
	GSessionSet = append(GSessionSet, sessionId)
	return sessionId
}

func CheckSessionID(id string)bool{
	for _, v := range GSessionSet{
		if id == v{
			return true
		}
	}

	return false;
}

func Md5(text string) string {
    hashMd5 := md5.New()
    io.WriteString(hashMd5, text)
	return fmt.Sprintf("%x", hashMd5.Sum(nil))
}

func SetCookie(w http.ResponseWriter,sessionid string) {
	cookie := http.Cookie{Name:"MYBLOG", Value: sessionid}
	http.SetCookie(w, &cookie)
}


