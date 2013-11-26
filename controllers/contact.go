/*
显示对联系方式
*/

package controllers

import (
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
	case "/admin/login":
			
	default:
		NotFoundHandler(rw, req)
	}
}