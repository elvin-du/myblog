/*
显示对联系方式
*/

package controllers

import (
	"html/template"
	"myblog/logger"
	"net/http"
)

type Contact struct {
	*Controller
}

func NewContact() *Contact {
	return &Contact{
		Controller: &Controller{},
	}
}

/*
所有[/contact]路由的请求，都要经过这里进行转发
*/
func (this *Contact) Handler(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	switch req.URL.Path {
	case "/contact":
		this.ContactHandler(rw, req)
	default:
		NotFoundHandler(rw, req)
	}
}

/*
联系界面
*/
func (this *Contact) ContactHandler(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		t, err := template.ParseFiles("views/contact/contact.html")
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
