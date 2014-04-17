package models

import (
	"database/sql"
	"html"
	"myblog/logger"
)

type Model struct {
}

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
