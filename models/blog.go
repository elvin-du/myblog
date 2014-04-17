package models

import (
	"database/sql"
	"errors"
	"html"
	"myblog/config"
	"myblog/logger"
	"strconv"
	"strings"
	"time"
)

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

	//sql := `SELECT * FROM myblog.blogs LIMIT `
	//因为博客文章不多，所以一次性的把所有博客都读取到内存中。
	//数据库更新频率很低，所以也没有内存和数据之间的同步
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
根据标签检索博客
*/
func (this *Model) QueryByTag(tagId int) (blogs []Blog, err error) {
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
根据博客标题检索博客
*/
func (this *Model) QueryByTitle(blogId int) (blogs []Blog, err error) {
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

func (this *Model) EditBlog(title string) error {
	//TODO
	return nil
}

func (this *Model) DelBlog(title string) error {
	//TODO
	return nil
}
