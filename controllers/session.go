package controllers

import(
	"fmt"
	"time"
	"math/rand"
	"errors"
	"crypto/md5"
	"strconv"
	"io"
	"net/http"
)
//map[seesionid]name
var GSessionSet = map[string]string{}

func CreateSessionID(name string)string{
    nano := time.Now().UnixNano()
    rand.Seed(nano)
    rndNum := rand.Int63()
    sessionId := Md5(Md5(strconv.FormatInt(nano, 10))+Md5(strconv.FormatInt(rndNum, 10)))
	GSessionSet[sessionId] = name
	return sessionId
}

// return "" when there is not @id in @GSessionSet
func CheckSessionID(id string)string{
	return GSessionSet[id]
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

func CheckCookie(r *http.Request)(name string, e error){
	cookie, err := r.Cookie("MYBLOG")
	if nil != err{
		return "",err
	}

	name = CheckSessionID(cookie.Value)
	if name == ""{
		return name, errors.New("cookie expired")
	}

	return name,nil
}
