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
	case strings.HasPrefix(urlPath, "/articles"):
		this.ArticlesHandler(rw, req)
	default:
		NotFoundHandler(rw, req)
	}
}

/*
显示所有文章
*/
func (this *Articles) ArticlesHandler(rw http.ResponseWriter, req *http.Request) {
	logger.Infoln("entered IndexHandler()")
	//username,err := CheckCookie(r)
	//if err != nil{
	//	log.Println(err)
	//http.Redirect(w,r,"/", http.StatusFound)
	//	}
	var blogs []models.Blog
	var err error
	model := models.Model{}
	switch req.Method {
	case "GET":
		switch {
		case "" != req.FormValue("title"):
			strTitle := req.FormValue("title")
			titleId, _ := strconv.Atoi(strTitle)
			blogs, err = this.ArticleByTitle(titleId)
		case "" != req.FormValue("tag"):
			strTag := req.FormValue("tag")
			tagId, _ := strconv.Atoi(strTag)
			blogs, err = this.ArticleByTitle(tagId)
		case "" != req.FormValue("page"):
			strPage := req.FormValue("page")
			pageId, _ := strconv.Atoi(strPage)
			blogs, err = this.ArticleByPage(pageId)
		default:
			//提取所有博客
			blogs, err = model.QueryBlogs()
			if nil != err {
				logger.Errorln(err)
				blogs = append(blogs, models.Blog{Title: "找不到文章！"})
				return
			}
		}

		//获取所有标签
		var tags []models.Tag
		err, tags = model.QueryTags()
		if nil != err {
			logger.Errorln(err)
			return
		}

		//格式化所有博客和标签，以便template包使用
		type tmp struct {
			Blgs []models.Blog
			Tags []models.Tag
		}
		tmp2 := tmp{blogs, tags}
		t, err := template.ParseFiles("views/articles/index.html")
		if nil != err {
			logger.Errorln(err)
			return
		}

		//显示内容
		if err = t.Execute(rw, tmp2); nil != err {
			logger.Errorln(err)
			return
		}
	case "POST":
		//TODO
	}
}

/*
根据用户选择的标题，获取具体博客
*/
func (this *Articles) ArticleByTitle(blogId int) (blog []models.Blog, err error) {
	model := models.Model{}
	//按照标题查找博客
	err, blog = model.QueryByTitle(blogId)
	if nil != err {
		logger.Errorln(err)
	}
	return
}

/*
根据用户选择的tag，来显示具体哪些文章
*/
func (this *Articles) ArticlesByTag(tagId int) (blogs []models.Blog, err error) {
	model := models.Model{}
	//按照标签查找博客
	err, blogs = model.QueryByTag(tagId)
	if nil != err {
		logger.Errorln(err)
		blogs = append(blogs, models.Blog{Title: "no article"})
	}

	return
}

/*
根据用户选择的page，来显示具体哪些文章
*/
func (this *Articles) ArticleByPage(tagId int) (blogs []models.Blog, err error) {
	return
}
