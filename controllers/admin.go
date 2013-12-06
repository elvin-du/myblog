/*
处理管理员相关业务
管理员用户注册，登录，删除博客，添加博客，修改博客，添加博客评论，删除博客评论
*/

package controllers

import (
	"html/template"
	"myblog/logger"
	"myblog/models"
	"net/http"
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
	case urlPath == "/admin/editblog":
		this.EditBlogHandler(rw, req)
	case urlPath == "/admin/delblog":
		this.DelBlogHandler(rw, req)
	case urlPath == "/admin/addcomment":
		this.AddCommentHandler(rw, req)
	case urlPath == "/admin/delcomment":
		this.DelCommentHandler(rw, req)
	default:
		NotFoundHandler(rw, req)
	}
}

/*
管理员登录
*/
func (this *Admin) LoginHandler(rw http.ResponseWriter, req *http.Request) {
	logger.Infoln("entered Login()")
	switch req.Method {
	case "GET":
		t, err := template.ParseFiles("views/admin/login.html")
		if nil != err {
			logger.Errorln(err)
			return
		}
		if err = t.Execute(rw, nil); nil != err {
			logger.Errorln(err)
			return
		}
	case "POST":
		req.ParseForm()
		name := req.FormValue("username")
		psw := req.FormValue("password")
		if err := CheckNamePsw(name, psw); nil != err {
			logger.Errorln(err)
			http.Redirect(rw, req, "/login", http.StatusFound)
			return
		}
		SetCookie(rw, CreateSessionID(name))
		http.Redirect(rw, req, "/add", http.StatusFound)
	}
}

/*
管理员注册
*/
func (this *Admin) RegisterHandler(rw http.ResponseWriter, req *http.Request) {
	logger.Infoln("entered Register()")
	switch req.Method {
	case "GET":
		t, err := template.ParseFiles("views/admin/register.html")
		if nil != err {
			logger.Errorln(err)
			return
		}
		if err = t.Execute(rw, nil); nil != err {
			logger.Errorln(err)
			return
		}
	case "POST":
		req.ParseForm()
		name := req.FormValue("username")
		psw := req.FormValue("password")
		confirmPsw := req.FormValue("confirm_password")
		if psw != confirmPsw {
			//TBD
		}
		model := models.Model{}
		model.AddUser(name, psw)
		http.Redirect(rw, req, "/", http.StatusFound)
	}
}

func (this *Admin) AddBlogHandler(rw http.ResponseWriter, req *http.Request) {
	logger.Infoln("entered AddBlogHandler()")
	req.ParseForm()
	//name,err :=CheckCookie(r)
	//if err != nil{
	//	log.Println(err)
	//	http.Redirect(w,r,"/", http.StatusFound)
	//	return
	//	}
	switch req.Method {
	case "GET":
		t, err := template.ParseFiles("views/admin/add_blog.html")
		if nil != err {
			logger.Errorln(err)
			return
		}
		//查询博客标签
		m := models.Model{}
		tags, err := m.QueryTags()
		if nil != err {
			logger.Errorln(err)
		}
		//为了使用查询出来的博客标签，格式化查询结果
		type tmp struct {
			Tags []models.Tag
		}
		tmp2 := tmp{tags}
		if err = t.Execute(rw, tmp2); nil != err {
			logger.Errorln(err)
			return
		}
	case "POST":
		title := req.FormValue("title")
		content := req.FormValue("content")
		tag := req.FormValue("tag")
		tagId, _ := strconv.Atoi(tag) //从前端返回的是tag的ID
		logger.Debugln("title: ", title)
		logger.Debugln("content: ", content)
		logger.Debugln("arcticleTag: ", tag)

		model := models.Model{}
		model.AddBlog(title, content, tagId)
		http.Redirect(rw, req, "/", http.StatusFound)
	}
}

/*
edit blog
*/
func (this *Admin) EditBlogHandler(rw http.ResponseWriter, req *http.Request) {
	//TODO
}

/*
del blog
*/
func (this *Admin) DelBlogHandler(rw http.ResponseWriter, req *http.Request) {
	//TODO
}

/*
添加评论的处理函数
*/
func (this *Admin) AddCommentHandler(rw http.ResponseWriter, req *http.Request) {
	//TODO
}

/*
del comment
*/
func (this *Admin) DelCommentHandler(rw http.ResponseWriter, req *http.Request) {
	//TODO
}
