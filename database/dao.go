package database

import (
	"database/sql"
	"myblog/logger"
	"strings"

	_ "github.com/Go-SQL-Driver/MySQL"
)

const (
	C_DATA_SOURCE_NAME = `root:JTabc.123@tcp(localhost:3306)/myblog?charset=utf8`
	C_DRIVER_NAME      = "mysql"
)

type Dao struct {
	*sql.DB
	// 构造sql语句相关
	tableName string
	//where     string
	//whereVal  []interface{} // where条件对应中字段对应的值
	//limit     string
	//order     string
	//// insert
	//columns   []string      // 需要插入数据的字段
	//colValues []interface{} // 需要插入字段对应的值
	//// query
	//selectCols []string // default:"*"
}

func NewDao(tablename string) *Dao {
	return &Dao{tableName: tablename}
}

func (this *Dao) Open() (err error) {
	this.DB, err = sql.Open(C_DRIVER_NAME, C_DATA_SOURCE_NAME)
	return
}

//e.g:["name"]"oliver"
func (this *Dao) Insert(data map[string]interface{}) (sql.Result, error) {
	cols, colsVal := Map2Struct(data)
	strSql := this.insertSql(cols)
	logger.Debugln("Insert sql:", strSql)
	err := this.Open()
	if err != nil {
		return nil, err
	}
	defer this.Close()
	stmt, err := this.Prepare(strSql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(colsVal...)
}

func (this *Dao) insertSql(cols []string) string {
	sqlStr := "INSERT INTO " + this.tableName + "("
	colsStr := strings.Join(cols, ",")
	colsVal := strings.Join(cols, "=?,")
	sqlStr = colsStr + ") VALUES(" + colsVal + ")"
	return sqlStr
}

//e.g:["name"]"oliver"
func (this *Dao) Update(data map[string]interface{}, where string) (sql.Result, error) {
	cols, colsVal := Map2Struct(data)
	strSql := this.updateSql(cols, where)
	logger.Debugln("Update sql:", strSql)
	err := this.Open()
	if err != nil {
		return nil, err
	}
	defer this.Close()
	stmt, err := this.Prepare(strSql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(colsVal...)
}

func (this *Dao) updateSql(cols []string, where string) string {
	updateStr := "UPDATE " + this.tableName + "SET "
	colsStr := strings.Join(cols, "=?,")
	updateStr += colsStr + where

	return updateStr
}

func (this *Dao) Find(data map[string]interface{}, where, order, limit string) error {
	cols, colsVal := Map2Struct(data)
	strSql := this.findSql(cols, where, order, limit)
	logger.Debugln("Find sql:", strSql)
	err := this.Open()
	if err != nil {
		return err
	}
	defer this.Close()

	rows, err := this.Query(strSql)
	if nil != err {
		logger.Errorln(err)
		return err
	}

	if rows.Next() {
		return rows.Scan(colsVal...)
	}

	return rows.Err()
}

func (this *Dao) findSql(cols []string, where, order, limit string) string {
	selectStr := "SELECT " + strings.Join(cols, ",") + " FROM " + this.tableName
	selectStr += " " + where + " ORDER BY " + order + " LIMIT " + limit
	return selectStr
}
