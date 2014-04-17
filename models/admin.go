package models

import (
	"database/sql"
	"errors"
	"html"
	"myblog/logger"
)

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
