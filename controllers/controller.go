package controllers

import(
	"net/http"
	"log"
	"html/template"
	"myblog/models"
	"strconv"
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
	//username,err := CheckCookie(r)
	//if err != nil{
	//	log.Println(err)
		//http.Redirect(w,r,"/", http.StatusFound)
//	}
	switch r.Method{
	case "GET":
		t,err := template.ParseFiles("views/index.html")
		if nil != err{
			log.Println(err)
			return
		}
		model := models.Model{}
		err, blogs:= model.QueryBlogs()
		if nil != err{
			log.Println(err)
			blogs = append(blogs, models.Blog{Content:"no article"})
		}
		err, tags:= model.QueryTags()
		if nil != err{
			log.Println(err)
		}
		type tmp struct{
			Blgs		[]models.Blog
			Tags		[]models.Tag
		}

		tmp2 := tmp{blogs, tags}
		if err = t.Execute(w, tmp2); nil != err{
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
	//name,err :=CheckCookie(r)
	//if err != nil{
	//	log.Println(err)
	//	http.Redirect(w,r,"/", http.StatusFound)
	//	return
//	}
	switch r.Method{
	case "GET":
		t,err := template.ParseFiles("views/edit.html")
		if nil != err{
			log.Println(err)
			return
		}
		m := models.Model{}
		err, tags:= m.QueryTags()
		if nil != err{
			log.Println(err)
		}
		type tmp struct{
			Tags []models.Tag
		}
		tmp2 := tmp{tags}
		if err = t.Execute(w, tmp2); nil != err{
			log.Println(err)
			return
		}
	case "POST":
		r.ParseForm()
		title := r.FormValue("title")
		content := r.FormValue("content")
		tag := r.FormValue("arcticle_tag")
		tagId, _ := strconv.Atoi(tag)
		//log.Println("title: ", title)
		//log.Println("content: ", content)
		//log.Println("arcticleTag: ", arcticleTag)

		model := models.Model{}
		model.AddBlog(title,content,tagId)
		http.Redirect(w,r,"/index",http.StatusFound)
	}
}

func (c *Controller)Articles(w http.ResponseWriter, r *http.Request){
	log.Println("entered Articles()")
	r.ParseForm()
	switch r.Method{
	case "GET":
		strBlogId := r.URL.String()[10:]
		blogId,_ := strconv.Atoi(strBlogId)
		log.Println("title:", blogId)
		model := models.Model{}
		err, blog := model.QueryByTitle(blogId)
		if nil != err{
			log.Println(err)
		}
		err, tags := model.QueryTags()
		if nil != err{
			log.Println(err)
		}

		type tmp struct{
			Blg		models.Blog
			Tags	[]models.Tag
		}
		tmp2 := tmp{blog, tags}

		t,err := template.ParseFiles("views/index.html")
		if nil != err{
			log.Println(err)
			return
		}
		if err = t.Execute(w, tmp2); nil != err{
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
		strTagId := r.URL.String()[6:]
		tagId,_ := strconv.Atoi(strTagId)
		log.Println("tagId:", tagId)
		model := models.Model{}
		err, blgs := model.QueryByTag(tagId)
		if nil != err{
			log.Println(err)
			blgs = append(blgs, models.Blog{Title:"no article"})
		}
		err, tags:= model.QueryTags()
		if nil != err{
			log.Println(err)
		}

		type tmp struct{
			Blgs	[]models.Blog
			Tags	[]models.Tag
		}
		tmp2 := tmp{blgs, tags}
		t,err := template.ParseFiles("views/index.html")
		if nil != err{
			log.Println(err)
			return
		}
		if err = t.Execute(w, tmp2); nil != err{
			log.Println(err)
			return
		}
	case "POST":
	}
}
