package main

import(
	"net/http"
	"log"
	"myblog/controllers"
)

func init(){
	log.SetFlags(log.Llongfile | log.LstdFlags)
}

func main(){
	println("listen on port :8888")
	if err := http.ListenAndServe(":8888",http.HandlerFunc(controllers.Register)); err != nil{
		log.Fatal(err)
	}
}
