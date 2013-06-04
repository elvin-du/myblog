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
	CondType string //VALUE:"byDate","byType","byTitle","byUser"
	Content interface{}//example of date:"2012-11-1"
}

type Blogs struct{
	Id				int
	Blogs			string
	CreateDate		string
	TypeName		string
	TypeId			int
	Title			string
	TitleId			int
}

type BlogType struct{
	Id			int
	BlogType	string
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

	sql := "INSERT myblog.title SET title=?"
	stmt,err := db.Prepare(sql)
	if nil != err{
		log.Println(err)
	}
	defer stmt.Close()
	_,err = stmt.Exec(title)
	if nil != err{
		log.Println(err)
	}

	insertSql := "INSERT myblog.blogs SET user_id =?, blogs=? , create_date=?,type_id=?,title_id=?"
	stmt, err = db.Prepare(insertSql)
	if nil != err{
		log.Print(err)
		return err
	}
	defer stmt.Close()

	now := strings.Split(time.Now().String(), " ")[0]
	titleId, typeId,userId := GetId(title,blogType,username)
	_, err = stmt.Exec(userId, blog,now, typeId, titleId)
	if nil != err{
		log.Print(err)
		return err
	}

	return nil
}

func GetId(title,blogType,username string)(titleId,typeId,userId int64){
	title = html.EscapeString(title)
	blogType = html.EscapeString(blogType)
	username = html.EscapeString(username)
	db, err := sql.Open("mysql", "root:dumx@tcp(localhost:3306)/myblog?charset=utf8")
	if nil != err{
		log.Print(err)
	}
	defer db.Close()
	sql := `SELECT id FROM myblog.title WHERE title = '` + title + `'`
	rows , err := db.Query(sql)
	if nil != err{
		log.Print(err)
	}
	for rows.Next(){
		rows.Scan(&titleId)
	}

	sql = `SELECT id FROM myblog.blog_type WHERE blog_type = '` + blogType + `'`
	rows , err = db.Query(sql)
	if nil != err{
		log.Print(err)
	}
	for rows.Next(){
		rows.Scan(&typeId)
	}

	sql = `SELECT userid FROM myblog.users WHERE name = '` + username+ `'`
	rows , err = db.Query(sql)
	if nil != err{
		log.Print(err)
	}
	for rows.Next(){
		rows.Scan(&userId)
	}
	return
}

func (m *Model)QueryBlogType()(err error,blgType []BlogType){
	db, err := sql.Open("mysql", "root:dumx@tcp(localhost:3306)/myblog?charset=utf8")
	if nil != err{
		log.Print(err)
		return
	}
	defer db.Close()
	querySql := "select * from myblog.blog_type"
	rows, err := db.Query(querySql)
	if nil != err{
		log.Print(err)
		return
	}

	flag := false
	for rows.Next(){
		flag = true
		var id int
		var blgTp string
		rows.Scan(&id,&blgTp)
		blgType = append(blgType,BlogType{id,blgTp})
	}
	if !flag{
		err = errors.New("not found")
	}
	return
}

//query condition perhaps change,so agument type is interface
func (m *Model)QueryBlogs(cond Condition)(err error,blogsSlice []Blogs){
	db, err := sql.Open("mysql", "root:dumx@tcp(localhost:3306)/myblog?charset=utf8")
	if nil != err{
		log.Print(err)
		return
	}
	defer db.Close()

	sql := `SELECT a.id,a.blogs,a.create_date, a.title_id,b.title, a.type_id, d.blog_type FROM myblog.blogs a, myblog.title b,myblog.users c, myblog.blog_type d WHERE `
	switch cond.CondType{
	case "byDate":
		sql += `a.user_id = c.userid and a.type_id = d.id AND a.title_id = b.id AND a.create_date= '` + cond.Content.(string) + `'`
	case "byTitle":
		sql += `a.user_id = c.userid and a.type_id = d.id AND a.title_id = b.id AND b.id= '` + cond.Content.(string) + `'`
	case "byType":
		sql += `a.user_id = c.userid and a.type_id = d.id AND a.title_id = b.id AND d.id = '` + cond.Content.(string) + `'`
	case "byName":
		sql += `a.user_id = c.userid and a.type_id = d.id AND a.title_id = b.id AND c.name = '` + cond.Content.(string) + `'`
	}
	log.Println("sql: ", sql)
	rows, err := db.Query(sql)
	if nil != err{
		log.Print(err)
	}

	for rows.Next(){
		var id ,titleId, typeId int
		var blogs,createDate, title, blogType string
		rows.Scan(&id, &blogs, &createDate, &titleId,&title, &typeId, &blogType)
		blogsSlice = append(blogsSlice,Blogs{id,blogs,createDate,blogType,typeId,title,titleId})
	}
	log.Println("blogsSlice:", blogsSlice)
	if 0 ==  len(blogsSlice){
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
