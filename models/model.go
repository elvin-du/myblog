package models

import (
	"database/sql"
	"errors"
	_ "github.com/Go-SQL-Driver/MySQL"
	"html"
	"myblog/config"
	"myblog/logger"
	"strconv"
	"strings"
	"time"
)

type Model struct {
}

/*
博客
*/
type Blog struct {
	Id         int
	Content    string
	Title      string
	CreateDate string
	TagId      int
	Comments   []Comment
}

/*
文章评论
*/
type Comment struct {
	Id         int
	IP         string
	Content    string
	CreateDate string
	BlogId     int
}

/*
文章标签
*/
type Tag struct {
	Id  int
	Tag string
}

/*
检查用户名和密码
Return Value：nil：success
*/
func (this *Model) CheckNamePsw(name, psw string) error {
	username := html.EscapeString(name)
	password := html.EscapeString(psw)
	db, err := sql.Open("mysql", "root:dumx@tcp(localhost:3306)/myblog?charset=utf8")
	if nil != err {
		logger.Errorln(err)
		return err
	}
	defer db.Close()
	querySql := "select 1 from myblog.users WHERE name = '" + username + "' AND password = '" + password + "'"
	rows, err := db.Query(querySql)
	if nil != err {
		logger.Errorln(err)
		return err
	}
	if rows.Next() {
		return nil
	}

	return errors.New("Unkown error")
}

/*
添加管理员
*/
func (this *Model) AddUser(name, psw string) error {
	username := html.EscapeString(name)
	password := html.EscapeString(psw)
	db, err := sql.Open("mysql", "root:dumx@tcp(localhost:3306)/myblog?charset=utf8")
	if nil != err {
		logger.Errorln(err)
		return err
	}
	defer db.Close()
	querySql := "select 1 from myblog.users WHERE name = ' " + username + "'"
	rows, err := db.Query(querySql)
	if nil != err {
		logger.Errorln(err)
		return err
	}
	if rows.Next() {
		return errors.New("user " + username + "exsited")
	}

	insertSql := "INSERT myblog.users SET name=?, password=?"
	stmt, err := db.Prepare(insertSql)
	if nil != err {
		logger.Errorln(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, password)
	if nil != err {
		logger.Errorln(err)
		return err
	}

	return nil
}

/*

*/
func GetId(title, blogType, username string) (titleId, typeId, userId int64) {
	title = html.EscapeString(title)
	blogType = html.EscapeString(blogType)
	username = html.EscapeString(username)
	db, err := sql.Open("mysql", "root:dumx@tcp(localhost:3306)/myblog?charset=utf8")
	if nil != err {
		logger.Errorln(err)
	}
	defer db.Close()
	sql := `SELECT id FROM myblog.title WHERE title = '` + title + `'`
	rows, err := db.Query(sql)
	if nil != err {
		logger.Errorln(err)
	}
	for rows.Next() {
		rows.Scan(&titleId)
	}

	sql = `SELECT id FROM myblog.blog_type WHERE blog_type = '` + blogType + `'`
	rows, err = db.Query(sql)
	if nil != err {
		logger.Errorln(err)
	}
	for rows.Next() {
		rows.Scan(&typeId)
	}

	sql = `SELECT userid FROM myblog.users WHERE name = '` + username + `'`
	rows, err = db.Query(sql)
	if nil != err {
		logger.Errorln(err)
	}
	for rows.Next() {
		rows.Scan(&userId)
	}
	return
}

/*
添加博客
*/
func (this *Model) AddBlog(title, content string, tagId int) error {
	//防止SQL注入
	title = html.EscapeString(title)
	content = html.EscapeString(content)

	//连接MYSQL数据库
	db, err := sql.Open("mysql", "root:dumx@tcp(localhost:3306)/myblog?charset=utf8")
	if nil != err {
		logger.Errorln(err)
		return err
	}
	defer db.Close()

	//构建Sql语句
	insertSql := "INSERT myblog.blogs SET content=?, title=?, create_date=?, tag_id=?"
	stmt, err := db.Prepare(insertSql)
	if nil != err {
		logger.Errorln(err)
		return err
	}
	defer stmt.Close()

	//replacer := strings.NewReplacer(" ", "&nbsp", "\r", "<br/>")
	//content = replacer.Replace(content)
	//获取插入时间
	now := strings.Split(time.Now().String(), " ")[0]
	//执行SQL语句
	_, err = stmt.Exec(content, title, now, tagId)
	if nil != err {
		logger.Errorln(err)
		return err
	}

	return nil
}

/*
query all blogs
*/
func (this *Model) QueryBlogs() (blogs []Blog, err error) {
	//连接数据库
	db, err := sql.Open(config.Config["driver_name"], config.Config["dsn"])
	if nil != err {
		logger.Errorln(err)
		return
	}
	defer db.Close()

	sql := `SELECT * FROM myblog.blogs`

	//执行SQL语句
	rows, err := db.Query(sql)
	if nil != err {
		logger.Errorln(err)
		return
	}

	//把查询到的数据格式化
	for rows.Next() {
		var id, tagId int
		var content, title, createDate string
		rows.Scan(&id, &content, &title, &createDate, &tagId)
		blogs = append(blogs, Blog{Id: id, Content: content, Title: title, CreateDate: createDate, TagId: tagId})
	}
	logger.Debugln("Blogs table :", blogs)

	//没有查询到博客，设置err值
	if 0 == len(blogs) {
		err = errors.New("not found")
	}
	return
}

/*
query all tags
*/
func (this *Model) QueryTags() (err error, tags []Tag) {
	db, err := sql.Open("mysql", "root:dumx@tcp(localhost:3306)/myblog?charset=utf8")
	if nil != err {
		logger.Errorln(err)
		return
	}
	defer db.Close()
	querySql := "select * from myblog.tags"
	rows, err := db.Query(querySql)
	if nil != err {
		logger.Errorln(err)
	}

	for rows.Next() {
		var id int
		var tag string
		rows.Scan(&id, &tag)
		tags = append(tags, Tag{id, tag})
	}
	logger.Debugln("Tags table :", tags)
	if 0 == len(tags) {
		err = errors.New("not found")
	}
	return
}

/*
根据博客标题检索博客
*/
func (this *Model) QueryByTitle(blogId int) (err error, blogs []Blog) {
	db, err := sql.Open("mysql", "root:dumx@tcp(localhost:3306)/myblog?charset=utf8")
	if nil != err {
		logger.Errorln(err)
		return
	}
	defer db.Close()
	querySql := `select * from myblog.blogs WHERE id=`
	tmp := strconv.FormatInt(int64(blogId), 10)
	querySql += tmp
	rows, err := db.Query(querySql)
	if nil != err {
		logger.Errorln(err)
	}

	for rows.Next() {
		var id, tagId int
		var content, title, createDate string
		rows.Scan(&id, &content, &title, &createDate, &tagId)
		blogs = append(blogs, Blog{Id: id, Content: content, Title: title, CreateDate: createDate, TagId: tagId})
	}
	logger.Debugln("Blogs table :", blogs)
	if 0 == len(blogs) {
		err = errors.New("not found")
	}
	return
}

/*
根据标签检索博客
*/
func (this *Model) QueryByTag(tagId int) (err error, blogs []Blog) {
	//TODO
	db, err := sql.Open("mysql", "root:dumx@tcp(localhost:3306)/myblog?charset=utf8")
	if nil != err {
		logger.Errorln(err)
		return
	}
	defer db.Close()
	querySql := "select * from myblog.blogs WHERE tag_id="
	tmp := strconv.FormatInt(int64(tagId), 10)
	querySql += tmp
	rows, err := db.Query(querySql)
	if nil != err {
		logger.Errorln(err)
	}

	for rows.Next() {
		var id, tagId int
		var content, title, createDate string
		rows.Scan(&id, &content, &title, &createDate, &tagId)
		blogs = append(blogs, Blog{Id: id, Content: content, Title: title, CreateDate: createDate, TagId: tagId})
	}
	logger.Debugln("Blogs table :", blogs)
	if 0 == len(blogs) {
		err = errors.New("not found")
	}
	return
}

/*
编辑博客
*/
func (this *Model) EditBlog(title string) error {
	//TODO
	return nil
}

/*
删除博客
*/
func (this *Model) DelBlog(title string) error {
	//TODO
	return nil
}

/*
添加评论
*/
func (this *Model) AddComment(title, commtent string) error {
	//TODO
	return nil
}

/*
删除评论
*/
func (this *Model) DelComment(title, comment string) error {
	//TODO
	return nil
}
