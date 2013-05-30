package models

import(
	_"github.com/Go-SQL-Driver/MySQL"
	"html"
	"strings"
	"log"
	"errors"
	"database/sql"
	"time"
)

type Model struct{
}

type Condition struct{
	CondType string //VALUE:"byDate","byType","byTitle","byName"
	Content interface{}//example of date:"2012-11-1"
}

type Blogs struct{
	Id				int
	Username		string
	Blogs			string
	CreateDate		string
	TypeName		string
	Title			string
}

func (m *Model)CheckNamePsw(name,psw string)error{
	username := html.EscapeString(name)
	password := html.EscapeString(psw)
	db, err := sql.Open("mysql", "root:dumx@tcp(localhost:3306)/myblog?charset=utf8")
	if nil != err{
		log.Print(err)
		return err
	}
	defer db.Close()
	querySql := "select 1 from myblog.users WHERE name = '" + username + "' AND password = '" + password + "'"
	rows, err := db.Query(querySql)
	if nil != err{
		log.Print(err)
		return err
	}
	if rows.Next(){
		return nil
	}

	return errors.New("Unkown error") 
}

func (m *Model)AddUser(name,psw string)error{
	username := html.EscapeString(name)
	password := html.EscapeString(psw)
	db, err := sql.Open("mysql", "root:dumx@tcp(localhost:3306)/myblog?charset=utf8")
	if nil != err{
		log.Print(err)
		return err
	}
	defer db.Close()
	querySql := "select 1 from myblog.users WHERE name = ' " + username + "'"
	rows, err := db.Query(querySql)
	if nil != err{
		log.Print(err)
		return err
	}
	if rows.Next(){
		return errors.New("user " + username + "exsited")
	}

	insertSql := "INSERT myblog.users SET name=?, password=?"
	stmt, err := db.Prepare(insertSql)
	if nil != err{
		log.Print(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, password)
	if nil != err{
		log.Print(err)
		return err
	}

	return nil
}

func (m *Model)EditBlogs(newBlog string)error{
	return nil
}

func (m *Model)AddBlogs(title,blog,blogType,username string)error{
	title = html.EscapeString(title)
	blog = html.EscapeString(blog)
	blogType = html.EscapeString(blogType)
	username = html.EscapeString(username)
	db, err := sql.Open("mysql", "root:dumx@tcp(localhost:3306)/myblog?charset=utf8")
	if nil != err{
		log.Print(err)
		return err
	}
	defer db.Close()
	querySql := "select 1 from myblog.blogs WHERE title = ' " + title + "'"
	rows, err := db.Query(querySql)
	if nil != err{
		log.Print(err)
		return err
	}
	if rows.Next(){
		return errors.New("title exsited")
	}

	insertSql := "INSERT myblog.blogs SET username =?, blogs=? , create_date=?,type_name=?,title=?"
	stmt, err := db.Prepare(insertSql)
	if nil != err{
		log.Print(err)
		return err
	}
	defer stmt.Close()

	now := strings.Split(time.Now().String(), " ")[0]
	_, err = stmt.Exec(username, blog,now,blogType,title)
	if nil != err{
		log.Print(err)
		return err
	}

	return nil
}

//query condition perhaps change,so agument type is interface
func (m *Model)QueryBlogs(cond Condition)(err error,blogsSlice []Blogs){
	db, err := sql.Open("mysql", "root:dumx@tcp(localhost:3306)/myblog?charset=utf8")
	if nil != err{
		log.Print(err)
		return
	}
	defer db.Close()

	querySql := "select * from myblog.blogs WHERE "//title = ' " //+ html.EscapeString(title) + "'"
	switch cond.CondType{
	case "byDate":
		date := cond.Content.(string)
		querySql += "create_date = '" + date + "'"
	case "byType":
		typeName := cond.Content.(string)
		querySql += "type_name = '" + typeName + "'"
	case "byTitle":
		title := cond.Content.(string)
		querySql += "title = '" + title+ "'"
	case "byName":
		name := cond.Content.(string)
		querySql += "username= '" + name + "'"
	}

	rows, err := db.Query(querySql)
	if nil != err{
		log.Print(err)
		return
	}

	flag := false
	for rows.Next(){
		flag = true
		var id int
		var username,blogs,createDate,typeName,title string
		rows.Scan(&id,&username,&blogs,&createDate,&typeName,&title)
		blogsSlice = append(blogsSlice,Blogs{id,username,blogs,createDate,typeName,title})
	}

	if !flag{
		err = errors.New("not found")
	}
	return
}

func (m *Model)DelBlogs(title string)error{
	return nil
}

func (m *Model)AddComments(title, commtent string)error{
	return nil
}

func (m *Model)DelComments(title string)error{
	return nil
}
