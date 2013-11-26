/*
显示对网站的描述
*/

package controllers

import (
	"html/template"
	"myblog/logger"
	"net/http"
)

type About struct {
	*Controller
}

func NewAbout() *About {
	return &About{
		Controller: &Controller{},
	}
}

/*
所有[/about]路由的请求，都要经过这里进行转发
*/
func (this *About) Handler(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	switch req.URL.Path {
	case "/about":
		this.AboutHandler(rw, req)
	default:
		NotFoundHandler(rw, req)
	}
}

/*
本站介绍
*/
func (this *About) AboutHandler(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		t, err := template.ParseFiles("views/about/about.html")
		if nil != err {
			logger.Errorln(err)
			return
		}
		if err = t.Execute(rw, nil); nil != err {
			logger.Errorln(err)
			return
		}
	case "POST":
	}
}
