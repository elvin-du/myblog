package controllers

import(
	"net/http"
	"log"
	"html/template"
)

type Controller struct{
}

func (c *Controller)NotFound(w http.ResponseWriter, r *http.Request){
	log.Println("entered NotFound()")
}

func (c *Controller)Login(w http.ResponseWriter, r *http.Request){
	log.Println("entered Login()")
	switch r.Method{
	case "GET":
		t,err := template.ParseFiles("views/login.html")
		if nil != err{
			log.Println(err)
			return
		}
		if err = t.Execute(w, nil); nil != err{
			log.Println(err)
			return
		}
	case "POST":
		r.ParseForm()
		name := r.FormValue("username")
		psw := r.FormValue("password")
		if err := CheckNamePsw(name,psw);nil != err{
			log.Println(err)
			return
		}
		http.Redirect(w,r,"/blog",http.StatusFound)
	}
}

func (c *Controller)Blog(w http.ResponseWriter, r *http.Request){
	log.Println("entered Blog()")
	switch r.Method{
	case "GET":
		t,err := template.ParseFiles("views/blog.html")
		if nil != err{
			log.Println(err)
			return
		}
		if err = t.Execute(w, nil); nil != err{
			log.Println(err)
			return
		}
	case "POST":
	}
}

func (c *Controller)Register(w http.ResponseWriter, r *http.Request){
	log.Println("entered Register()")
}

func (c *Controller)Edit(w http.ResponseWriter, r *http.Request){
	log.Println("entered Edit()")
}
