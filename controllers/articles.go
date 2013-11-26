/*
文章显示
*/

package controllers

import (
	"html/template"
	"myblog/logger"
	"myblog/models"
	"net/http"
	"strconv"
	"strings"
)

type Articles struct {
	*Controller
}

/*
create a  new Arcticle
*/
func NewArticles() *Articles {
	return &Articles{
		Controller: &Controller{},
	}
}

/*
所有[/Articles]路由的请求，都要经过这里进行转发
*/
func (this *Articles) Handler(rw http.ResponseWriter, req *http.Request) {
	urlPath := req.URL.Path
	switch {
	case urlPath == "/":
		this.IndexHandler(rw, req)
	case strings.HasPrefix(urlPath, "/articles/tags"): //case顺序不能改变
		this.ArticlesByTagHandler(rw, req)
	case strings.HasPrefix(urlPath, "/articles"):
		this.ArticleByTitleHandler(rw, req)
	default:
		NotFoundHandler(rw, req)
	}
}

/*
显示所有文章
*/
func (this *Articles) IndexHandler(rw http.ResponseWriter, req *http.Request) {
	logger.Infoln("entered IndexHandler()")
	//username,err := CheckCookie(r)
	//if err != nil{
	//	log.Println(err)
	//http.Redirect(w,r,"/", http.StatusFound)
	//	}
	switch req.Method {
	case "GET":
		t, err := template.ParseFiles("views/articles/index.html")
		if nil != err {
			logger.Errorln(err)
			return
		}
		model := models.Model{}
		err, blogs := model.QueryBlogs()
		if nil != err {
			logger.Errorln(err)
			blogs = append(blogs, models.Blog{Title: "找不到文章！"})
		}
		err, tags := model.QueryTags()
		if nil != err {
			logger.Errorln(err)
		}
		type tmp struct {
			Blgs []models.Blog
			Tags []models.Tag
		}

		tmp2 := tmp{blogs, tags}
		if err = t.Execute(rw, tmp2); nil != err {
			logger.Errorln(err)
			return
		}
	case "POST":
	}
}

/*
根据用户选择的标题，来显示具体那篇文章
*/
func (this *Articles) ArticleByTitleHandler(w http.ResponseWriter, r *http.Request) {
	logger.Infoln("entered Articles()")
	r.ParseForm()
	switch r.Method {
	case "GET":
		strBlogId := r.URL.String()[10:]
		blogId, _ := strconv.Atoi(strBlogId)
		logger.Infoln("title:", blogId)
		model := models.Model{}
		err, blog := model.QueryByTitle(blogId)
		if nil != err {
			logger.Errorln(err)
		}
		err, tags := model.QueryTags()
		if nil != err {
			logger.Errorln(err)
		}

		type tmp struct {
			Blgs []models.Blog
			Tags []models.Tag
		}
		tmp2 := tmp{blog, tags}

		t, err := template.ParseFiles("views/articles/index.html")
		if nil != err {
			logger.Errorln(err)
			return
		}
		if err = t.Execute(w, tmp2); nil != err {
			logger.Errorln(err)
			return
		}
	case "POST":
	}
}

/*
根据用户选择的tag，来显示具体哪些文章
*/
func (this *Articles) ArticlesByTagHandler(w http.ResponseWriter, r *http.Request) {
	logger.Infoln("entered ArticlesByTagHandler()")
	r.ParseForm()
	switch r.Method {
	case "GET":
		strTagId := r.URL.String()[15:]
		tagId, _ := strconv.Atoi(strTagId)
		logger.Infoln("tagId:", tagId)
		model := models.Model{}
		err, blgs := model.QueryByTag(tagId)
		if nil != err {
			logger.Errorln(err)
			blgs = append(blgs, models.Blog{Title: "no article"})
		}
		err, tags := model.QueryTags()
		if nil != err {
			logger.Errorln(err)
		}

		type tmp struct {
			Blgs []models.Blog
			Tags []models.Tag
		}
		tmp2 := tmp{blgs, tags}
		t, err := template.ParseFiles("views/articles/index.html")
		if nil != err {
			logger.Errorln(err)
			return
		}
		if err = t.Execute(w, tmp2); nil != err {
			logger.Errorln(err)
			return
		}
	case "POST":
	}
}
