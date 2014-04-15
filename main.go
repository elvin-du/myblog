package main

import (
	"log"
	"lwebf/web"
	"myblog/config"
	"myblog/logger"
)

func init() {
	log.SetFlags(log.Llongfile | log.LstdFlags)
}

func main() {
	log.Println("listen on port :8888")
	logger.Infoln("Server Started, listen on port: 8888")

	register()
	web.Run(config.Config["host"])
	//if err := http.ListenAndServe(config.Config["host"], http.HandlerFunc(router)); err != nil {
	//	log.Println("ListenAndServe:", err)
	//	logger.Errorln("ListenAndServe:", err)
	//	return
	//}
}
