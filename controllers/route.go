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
	default:
		ctrl.NotFound(w,r)
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
	case "/edit" == urlPath:
		ctrl.Edit(w,r)
	case "/" == urlPath:
		ctrl.Index(w,r)
	case strings.HasPrefix(urlPath, "/articles"):
		ctrl.Articles(w,r)
	case strings.HasPrefix(urlPath, "/tags"):
		ctrl.Tags(w,r)
	}
}
