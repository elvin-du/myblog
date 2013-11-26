/*
显示对网站的描述
*/

package controllers

import (
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
所有[/contact]路由的请求，都要经过这里进行转发
*/
func (this *About) Handler(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	switch req.URL.Path {
	case "/admin/login":
			
	default:
		NotFoundHandler(rw, req)
	}
}