/*
处理管理员相关业务
*/

package controllers

import (
	"net/http"
	"myblog/logger"
	"html/template"
	"myblog/models"
	"strconv"
)

type Admin struct {
	*Controller
}

func NewAdmin() *Admin {
	return &Admin{
		Controller: &Controller{},
	}
}

/*
所有[/admin/*]路由的请求，都要经过这里进行转发
*/
func (this *Admin) Handler(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	urlPath := req.URL.Path
	switch {
	case urlPath == "/admin/login":
		this.LoginHandler(rw, req)
	case urlPath == "/admin/register":
		this.RegisterHandler(rw, req)
	case urlPath == "/admin/addblog":
		this.AddBlogHandler(rw, req)
	case urlPath == "/admin/addcomment":
		this.AddCommentHandler(rw, req)
	case urlPath == "/admin/delcomment":
		this.DelHandler(rw,req)
	case urlPath == "/admin/editblog":
		this.EditBlogHandler(rw, req)
	default:
		NotFoundHandler(rw, req)
	}
}

/*
管理员登录
*/
func (this *Admin) LoginHandler(rw http.ResponseWriter, req *http.Request) {
	logger.Infoln("entered Login()")
	switch req.Method{
	case "GET":
		t,err := template.ParseFiles("views/login.html")
		if nil != err{
			logger.Errorln(err)
			return
		}
		if err = t.Execute(rw, nil); nil != err{
			logger.Errorln(err)
			return
		}
	case "POST":
		req.ParseForm()
		name := req.FormValue("username")
		psw := req.FormValue("password")
		if err := CheckNamePsw(name,psw);nil != err{
			logger.Errorln(err)
			http.Redirect(rw,req,"/login",http.StatusFound)
			return
		}
		SetCookie(rw, CreateSessionID(name))
		http.Redirect(rw,req,"/add",http.StatusFound)
	}
}

/*
管理员注册
*/
func (this *Admin) RegisterHandler(rw http.ResponseWriter, req *http.Request) {
	logger.Infoln("entered Register()")
	switch req.Method{
	case "GET":
		t,err := template.ParseFiles("views/register.html")
		if nil != err{
			logger.Errorln(err)
			return
		}
		if err = t.Execute(rw, nil); nil != err{
			logger.Errorln(err)
			return
		}
	case "POST":
		req.ParseForm()
		name := req.FormValue("username")
		psw := req.FormValue("password")
		confirmPsw := req.FormValue("confirm_password")
		if psw != confirmPsw{
			//TBD
		}
		model := models.Model{}
		model.AddUser(name,psw)
		http.Redirect(rw,req,"/",http.StatusFound)
	}
}


func (c *Admin)AddBlogHandler(rw http.ResponseWriter, req *http.Request){
	logger.Infoln("entered Add()")
	req.ParseForm()
	//name,err :=CheckCookie(r)
	//if err != nil{
	//	log.Println(err)
	//	http.Redirect(w,r,"/", http.StatusFound)
	//	return
//	}
	switch req.Method{
	case "GET":
		t,err := template.ParseFiles("views/edit.html")
		if nil != err{
			logger.Errorln(err)
			return
		}
		m := models.Model{}
		err, tags:= m.QueryTags()
		if nil != err{
			logger.Errorln(err)
		}
		type tmp struct{
			Tags []models.Tag
		}
		tmp2 := tmp{tags}
		if err = t.Execute(rw, tmp2); nil != err{
			logger.Errorln(err)
			return
		}
	case "POST":
		title := req.FormValue("title")
		content := req.FormValue("content")
		tag := req.FormValue("tag")
		tagId, _ := strconv.Atoi(tag)
		//log.Println("title: ", title)
		//log.Println("content: ", content)
		//log.Println("arcticleTag: ", arcticleTag)

		model := models.Model{}
		model.AddBlog(title,content,tagId)
		http.Redirect(rw,req,"/",http.StatusFound)
	}
}

//edit blog
func (c *Admin)EditBlogHandler(rw http.ResponseWriter, req *http.Request){
	//TODO
}

/*
添加评论的处理函数
*/
func (c *Admin)AddCommentHandler(rw http.ResponseWriter, req *http.Request){
	//TODO
}

//del blog or comment
func (c *Admin)DelHandler(rw http.ResponseWriter, req *http.Request){
	//TODO
}