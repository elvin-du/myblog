package controllers

import(
	"net/http"
	"log"
	"html/template"
	"myblog/models"
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
			http.Redirect(w,r,"/",http.StatusFound)
			return
		}
		SetCookie(w, CreateSessionID(name))
		http.Redirect(w,r,"/index",http.StatusFound)
	}
}

func (c *Controller)Index(w http.ResponseWriter, r *http.Request){
	log.Println("entered Index()")
	switch r.Method{
	case "GET":
		t,err := template.ParseFiles("views/index.html")
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
	switch r.Method{
	case "GET":
		t,err := template.ParseFiles("views/register.html")
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
		confirmPsw := r.FormValue("confirm_password")
		if psw != confirmPsw{
			//TBD
		}
		model := models.Model{}
		model.AddUser(name,psw)
		http.Redirect(w,r,"/index",http.StatusFound)
	}
}

func (c *Controller)Edit(w http.ResponseWriter, r *http.Request){
	log.Println("entered Edit()")
	name,err :=CheckCookie(r) 
	if err != nil{
		log.Println(err)
		http.Redirect(w,r,"/", http.StatusFound)
		return
	}
	switch r.Method{
	case "GET":
		t,err := template.ParseFiles("views/edit.html")
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
		title := r.FormValue("title")
		context := r.FormValue("context")
		log.Println("title: ", title)
		log.Println("context: ", context)

		model := models.Model{}
		model.AddBlogs(title,context,"",name)
		http.Redirect(w,r,"/index",http.StatusFound)
	}
}


