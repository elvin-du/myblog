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

//保存所有博客内容
var AllBlogs = []models.Blog{}

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
	case urlPath == "/" || strings.HasPrefix(urlPath, "/articles"):
		this.ArticlesHandler(rw, req)
	default:
		NotFoundHandler(rw, req)
	}
}

/*
显示所有文章
*/
func (this *Articles) ArticlesHandler(rw http.ResponseWriter, req *http.Request) {
	logger.Infoln("entered ArticlesHandler()")
	//username,err := CheckCookie(r)
	//if err != nil{
	//	log.Println(err)
	//http.Redirect(w,r,"/", http.StatusFound)
	//	}
	var blogs []models.Blog
	var err error
	switch req.Method {
	case "GET":
		switch {
		case "" != req.FormValue("title"):
			strTitle := req.FormValue("title")
			logger.Debugln(strTitle)
			titleId, _ := strconv.Atoi(strTitle)
			blogs, err = this.ArticleByTitle(titleId)
		case "" != req.FormValue("tag"):
			strTag := req.FormValue("tag")
			logger.Debugln(strTag)
			tagId, _ := strconv.Atoi(strTag)
			blogs, err = this.ArticleByTitle(tagId)
		case "" != req.FormValue("page"):
			strPage := req.FormValue("page")
			logger.Debugln(strPage)
			pageId, _ := strconv.Atoi(strPage)
			blogs, err = this.ArticleByPage(pageId)
		default:
			err = this.QueryAllBlogs()
			if nil != err {
				logger.Errorln(err)
				return
			}
			blogs = AllBlogs
		}
		//获取所有标签
		var tags []models.Tag
		tags, err = new(models.Model).QueryTags()
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
	//按照标题查找博客
	blog, err = new(models.Model).QueryByTitle(blogId)
	if nil != err {
		logger.Errorln(err)
	}
	return
}

/*
根据用户选择的tag，来显示具体哪些文章
*/
func (this *Articles) ArticlesByTag(tagId int) (blogs []models.Blog, err error) {
	//按照标签查找博客
	blogs, err = new(models.Model).QueryByTag(tagId)
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
	err = this.QueryAllBlogs()
	if nil != err {
		logger.Errorln(err)
		return
	}

	const MaxPageNum = 3 //一页最大显示博客数量
	//访问的page没有超出
	if tagId <= 0 || len(AllBlogs) < tagId*MaxPageNum+MaxPageNum {
		blogs = AllBlogs[0:MaxPageNum : MaxPageNum+1]
	} else {
		blogs = AllBlogs[tagId*MaxPageNum : tagId*MaxPageNum+MaxPageNum : MaxPageNum+1]
	}

	return
}

/*
获取所有博客内容存储在@AllBlogs
*/
func (this *Articles) QueryAllBlogs() (err error) {
	if 0 < len(AllBlogs) {
		return //已经查询过，AllBlogs已经有数据
	}
	//提取所有博客
	AllBlogs, err = new(models.Model).QueryBlogs()
	if nil != err {
		logger.Errorln(err)
	}
	return
}
