package main

import (
	"lwebf/web"
	"myblog/controllers"

	_ "github.com/Go-SQL-Driver/MySQL"
)

func register() {
	web.Add(&controllers.AboutController{})
}
