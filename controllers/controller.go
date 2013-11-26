package controllers

import(
	"net/http"
	"html/template"
	"myblog/logger"
	"myblog/utils"
)

type Controller struct{
}

func (this *Controller) Handler(rw http.ResponseWriter, req *http.Request) {
	NotFoundHandler(rw, req)
}

func NotFoundHandler(rw http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles(utils.BaseHtmlTplFile, utils.Error404HtmlTplFile)
	if nil != err {
		logger.Errorln(err)
		return
	}
	err = t.Execute(rw, nil)
	if nil != err {
		logger.Errorln(err)
		return
	}
}