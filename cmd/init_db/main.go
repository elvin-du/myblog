package main

import (
	"log"
	"myblog/database"
	"myblog/logger"
	"strings"
	"time"
)

func main() {
	InitDB()
}

func InitDB() {
	db, err := database.Open()
	if nil != err {
		log.Fatal(err)
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

}
