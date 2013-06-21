package controllers

import(
	"net/http"
	"strings"
	"log"
)

func Register(w http.ResponseWriter, r *http.Request){
	urlPath := strings.ToLower(r.URL.Path)
	log.Println("URL: ",urlPath)
	ctrl := &Controller{}

	switch{
	case strings.HasPrefix(urlPath, "/public")://static files
		http.ServeFile(w,r, urlPath[1:])
	case strings.HasPrefix(urlPath, "/bootstrap")://static files
		http.ServeFile(w,r, "libs" + urlPath)
	case "/favicon.ico" == urlPath: //browser itself requests
		http.ServeFile(w,r,"public/images/favicon.ico")
	case "/login" == urlPath:
		ctrl.Login(w,r)
	case "/register" == urlPath:
		ctrl.Register(w,r)
	case "/" == urlPath:
		ctrl.Index(w,r)
	case "/add" == urlPath://add blog
		ctrl.AddBlog(w,r)
	case strings.HasPrefix(urlPath, "/edit")://edit blog, comment can not be edited, just can be deleted 
		ctrl.Edit(w,r)
	case strings.HasPrefix(urlPath, "/del"): //delete blog or comment
		ctrl.Del(w,r)
	case "/add_comment" == urlPath:
		ctrl.AddComment(w,r)
	case strings.HasPrefix(urlPath, "/articles"):
		ctrl.Articles(w,r)
	case strings.HasPrefix(urlPath, "/tags"):
		ctrl.Tags(w,r)
	default:
		ctrl.NotFound(w,r)
	}
}
