package models

import (
	"database/sql"
	"errors"
	"myblog/logger"
)

/*
query all tags
*/
func (this *Model) QueryTags() (tags []Tag, err error) {
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
