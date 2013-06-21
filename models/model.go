package models

import(
	_"github.com/Go-SQL-Driver/MySQL"
	"html"
	"strings"
	"log"
	"errors"
	"database/sql"
	"time"
	"strconv"
)

type Model struct{
}

type Blog struct{
	Id				int
	Content			string
	Title			string
	CreateDate		string
	TagId			int
	Comments		[]Comment
}

type Comment struct{
	Id			int
	IP			string
	Content		string
	CreateDate	string
	BlogId		int
}

type Tag struct{
	Id			int
	Tag			string
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

func (m *Model)AddBlog(title, content string, tagId int)error{
	title = html.EscapeString(title)
	content = html.EscapeString(content)
	db, err := sql.Open("mysql", "root:dumx@tcp(localhost:3306)/myblog?charset=utf8")
	if nil != err{
		log.Print(err)
		return err
	}
	defer db.Close()

	insertSql := "INSERT myblog.blogs SET content=?, title=?, create_date=?, tag_id=?"
	stmt, err := db.Prepare(insertSql)
	if nil != err{
		log.Print(err)
		return err
	}
	defer stmt.Close()

	now := strings.Split(time.Now().String(), " ")[0]
	_, err = stmt.Exec(content, title, now, tagId)
	if nil != err{
		log.Print(err)
		return err
	}

	return nil
}

//query all blogs
func (m *Model)QueryBlogs()(err error,blogs []Blog){
	db, err := sql.Open("mysql", "root:dumx@tcp(localhost:3306)/myblog?charset=utf8")
	if nil != err{
		log.Print(err)
		return
	}
	defer db.Close()
	sql := `SELECT * FROM myblog.blogs`
	rows, err := db.Query(sql)
	if nil != err{
		log.Print(err)
	}

	for rows.Next(){
		var id, tagId int
		var content, title, createDate string
		rows.Scan(&id, &content, &title, &createDate, &tagId)
		blogs = append(blogs,Blog{Id:id, Content:content,Title:title, CreateDate:createDate,TagId:tagId})
	}
	log.Println("Blogs table :", blogs)
	if 0 ==  len(blogs){
		err = errors.New("not found")
	}
	return
}

//query all tags
func (m *Model)QueryTags()(err error,tags []Tag){
	db, err := sql.Open("mysql", "root:dumx@tcp(localhost:3306)/myblog?charset=utf8")
	if nil != err{
		log.Print(err)
		return
	}
	defer db.Close()
	querySql := "select * from myblog.tags"
	rows, err := db.Query(querySql)
	if nil != err{
		log.Print(err)
	}

	for rows.Next(){
		var id int
		var tag string
		rows.Scan(&id,&tag)
		log.Println("tag", tag)
		tags = append(tags, Tag{id,tag})
	}
	log.Println("Tags table :", tags)
	if 0 == len(tags){
		err = errors.New("not found")
	}
	return
}

func (m *Model)QueryByTitle(blogId int)(err error, blogs []Blog){
	db, err := sql.Open("mysql", "root:dumx@tcp(localhost:3306)/myblog?charset=utf8")
	if nil != err{
		log.Print(err)
		return
	}
	defer db.Close()
	querySql := `select * from myblog.blogs WHERE id=`
	tmp := strconv.FormatInt(int64(blogId), 10)
	querySql += tmp
	rows, err := db.Query(querySql)
	if nil != err{
		log.Print(err)
	}

	for rows.Next(){
		var id, tagId int
		var content, title, createDate string
		rows.Scan(&id, &content, &title, &createDate, &tagId)
		blogs = append(blogs,Blog{Id:id, Content:content,Title:title, CreateDate:createDate,TagId:tagId})
	}
	log.Println("Blogs table :", blogs)
	if 0 ==  len(blogs){
		err = errors.New("not found")
	}
	return
}

func (m *Model)QueryByTag(tagId int)(err error, blogs []Blog){
	//TODO
	db, err := sql.Open("mysql", "root:dumx@tcp(localhost:3306)/myblog?charset=utf8")
	if nil != err{
		log.Print(err)
		return
	}
	defer db.Close()
	querySql := "select * from myblog.blogs WHERE tag_id="
	tmp := strconv.FormatInt(int64(tagId), 10)
	querySql += tmp
	rows, err := db.Query(querySql)
	if nil != err{
		log.Print(err)
	}

	for rows.Next(){
		var id, tagId int
		var content, title, createDate string
		rows.Scan(&id, &content, &title, &createDate, &tagId)
		blogs = append(blogs,Blog{Id:id, Content:content,Title:title, CreateDate:createDate,TagId:tagId})
	}
	log.Println("Blogs table :", blogs)
	if 0 ==  len(blogs){
		err = errors.New("not found")
	}
	return
}

func (m *Model)EditBlog(title string)error{
	//TODO
	return nil
}

func (m *Model)DelBlog(title string)error{
	//TODO
	return nil
}

func (m *Model)AddComment(title, commtent string)error{
	//TODO
	return nil
}

func (m *Model)DelComment(title, comment string)error{
	//TODO
	return nil
}
