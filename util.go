package main

import (
	"lwebf/web"
	"myblog/controllers"
)

func register() {
	web.Add(&controllers.AboutController{})
}
