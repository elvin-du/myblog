package main

import (
	//"code.google.com/p/go.net/websocket"
	"myblog/controllers"
	"myblog/logger"
	"net/http"
	"strings"
	"log"
)

const (
	rootUrl     = "/"
	articlesUrl = "/articles"
	adminUrl    = "/admin"
	wsUrl       = "/ws"
	contactUrl  = "/contact"
	aboutUrl  	= "/about"
)

/*
对所有的URL进行注册
*/
func router(rw http.ResponseWriter, req *http.Request) {
	urlPath := req.URL.Path
	log.Println(urlPath)
	logger.Debugln(urlPath)
	switch {
	case rootUrl == urlPath || strings.HasPrefix(urlPath, articlesUrl):
		friends := controllers.NewArticles()
		friends.Handler(rw, req)
	case strings.HasPrefix(urlPath, adminUrl):
		admin := controllers.NewAdmin()
		admin.Handler(rw, req)
	case strings.HasPrefix(urlPath, wsUrl):
		ws := controllers.NewWS()
		ws.Handler(rw, req)
	case strings.HasPrefix(urlPath, contactUrl):
		contact := controllers.NewContact()
		contact.Handler(rw, req)
	case strings.HasPrefix(urlPath, aboutUrl):
		about := controllers.NewAbout()
		about.Handler(rw, req)
	case strings.HasPrefix(urlPath, "/bootstrap")://static files
		http.ServeFile(rw,req, "libs" + urlPath)
	case strings.HasPrefix(urlPath, "public/"): //static files
		http.ServeFile(rw, req, urlPath)
	case urlPath == "/favicon.ico": //the request which browser send automatically
		http.ServeFile(rw, req, "public/images/favicon.ico")
	default:
		controllers.NotFoundHandler(rw, req)
	}
}
