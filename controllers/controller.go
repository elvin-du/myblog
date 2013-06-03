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
	username,err := CheckCookie(r)
	if err != nil{
		log.Println(err)
		http.Redirect(w,r,"/", http.StatusFound)
		return
	}
	switch r.Method{
	case "GET":
		t,err := template.ParseFiles("views/index.html")
		if nil != err{
			log.Println(err)
			return
		}
		cond := models.Condition{"byName",username}
		model := models.Model{}
		err, blogsSlice := model.QueryBlogs(cond)
		if nil != err{
			log.Println(err)
			blogsSlice = append(blogsSlice, models.Blogs{Blogs:"no article"})
		}
		err, blgTpSlice := model.QueryBlogType()
		if nil != err{
			log.Println(err)
		}
		type tmp struct{
			Blg			[]models.Blogs
			BlgTp		[]models.BlogType
		}

		blg := tmp{blogsSlice, blgTpSlice}
		if err = t.Execute(w, blg); nil != err{
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
		m := models.Model{}
		err,blogTp := m.QueryBlogType()
		if nil != err{
			log.Println(err)
		}
		type tmp struct{
			BlogTp []models.BlogType
		}
		bt := tmp{blogTp}
		if err = t.Execute(w, bt); nil != err{
			log.Println(err)
			return
		}
	case "POST":
		r.ParseForm()
		title := r.FormValue("title")
		content := r.FormValue("content")
		arcticleTag := r.FormValue("arcticle_tag")
		//log.Println("title: ", title)
		//log.Println("content: ", content)
		//log.Println("arcticleTag: ", arcticleTag)

		model := models.Model{}
		model.AddBlogs(title,content,arcticleTag,name)
		http.Redirect(w,r,"/index",http.StatusFound)
	}
}

func (c *Controller)Articles(w http.ResponseWriter, r *http.Request){
	log.Println("entered Articles()")
	r.ParseForm()
	switch r.Method{
	case "GET":
		t,err := template.ParseFiles("views/index.html")
		if nil != err{
			log.Println(err)
			return
		}
		title := r.URL.String()[10:]
		log.Println("title:", title)
		cond := models.Condition{"byTitle",title}
		model := models.Model{}
		err, blogsSlice := model.QueryBlogs(cond)
		if nil != err{
			log.Println(err)
			return
		}
		err, blgTpSlice := model.QueryBlogType()
		if nil != err{
			log.Println(err)
			return
		}
		type tmp struct{
			Blg			[]models.Blogs
			BlgTp		[]models.BlogType
		}

		blg := tmp{blogsSlice, blgTpSlice}
		if err = t.Execute(w, blg); nil != err{
			log.Println(err)
			return
		}
	case "POST":
	}
}

func (c *Controller)Tags(w http.ResponseWriter, r *http.Request){
	log.Println("entered Tags()")
	r.ParseForm()
		switch r.Method{
	case "GET":
		t,err := template.ParseFiles("views/index.html")
		if nil != err{
			log.Println(err)
			return
		}
		tag := r.URL.String()[6:]
		log.Println("tag:", tag)
		cond := models.Condition{"byType",tag}
		model := models.Model{}
		err, blogsSlice := model.QueryBlogs(cond)
		if nil != err{
			log.Println(err)
			return
		}
		err, blgTpSlice := model.QueryBlogType()
		if nil != err{
			log.Println(err)
			return
		}
		type tmp struct{
			Blg			[]models.Blogs
			BlgTp		[]models.BlogType
		}

		blg := tmp{blogsSlice, blgTpSlice}
		if err = t.Execute(w, blg); nil != err{
			log.Println(err)
			return
		}
	case "POST":
	}
}
