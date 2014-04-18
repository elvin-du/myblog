package database

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/Go-SQL-Driver/MySQL"
)

const (
	C_DATA_SOURCE_NAME = `root:JTabc.123@tcp(localhost:3306)/myblog?charset=utf8`
	C_DRIVER_NAME      = "mysql"
)

type Dao struct {
	*sql.DB
	tableName string
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
	log.Println("Insert sql:", strSql)
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
	colsStr := strings.Join(cols, ",")
	placeHolder := strings.Repeat("?,", len(cols))
	sqlStr := fmt.Sprintf("INSERT INTO  `%s`(%s)  VALUES(%s)", this.tableName, colsStr, placeHolder[:len(placeHolder)-1])
	return strings.TrimSpace(sqlStr)
}

//e.g:["name"]"oliver"
func (this *Dao) Update(data map[string]interface{}, where map[string]interface{}) error {
	cols, colsVal := Map2Struct(data)
	ws, wsVal := Map2Struct(where)
	strSql := this.updateSql(cols, ws)
	log.Println("Update sql:", strSql)
	err := this.Open()
	if err != nil {
		return err
	}
	defer this.Close()

	result, err := this.Exec(strSql, append(colsVal, wsVal...)...)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	log.Printf("table:`%s` %d columns update success", this.tableName, affected)
	return nil
}

func (this *Dao) updateSql(cols []string, where []string) string {
	colsStr := strings.Join(cols, "=?,") + "=?"
	whereStr := strings.Join(where, "=?,")
	if "" != whereStr {
		whereStr = "WHERE " + whereStr + "=?"
	}
	sqlStr := fmt.Sprintf("UPDATE `%s` SET %s %s", this.tableName, colsStr, whereStr)
	return strings.TrimSpace(sqlStr)
}

func (this *Dao) FindOne(data map[string]interface{}, where map[string]interface{}, selectCol ...string) error {
	cols, colsVal := Map2Struct(data)
	whereS, whereVal := Map2Struct(where)
	colNum := len(selectCol)
	if colNum == 0 || (colNum == 1 && selectCol[0] == "*") {
		selectCol = cols
	} else {
		tmp := []interface{}{}
		for _, v := range selectCol {
			tmp = append(tmp, data[v])
		}
		colsVal = tmp
	}

	strSql := this.findSql(selectCol, whereS, "", "")
	log.Println("Find sql:", strSql)
	err := this.Open()
	if err != nil {
		return err
	}
	defer this.Close()

	row := this.QueryRow(strSql, whereVal...)
	if nil != err {
		log.Println(err)
		return err
	}

	return row.Scan(colsVal...)
}

func (this *Dao) findSql(cols []string, where []string, order, limit string) string {
	colsStr := strings.Join(cols, ",")
	whereStr := strings.Join(where, "=?,")
	if "" != whereStr {
		whereStr = "WHERE " + whereStr + "=?"
	}
	orderStr := ""
	if "" != order {
		orderStr = "ORDER BY " + order
	}
	limitStr := ""
	if "" != whereStr {
		limitStr = "LIMIT " + limit
	}

	sqlStr := fmt.Sprintf("SELECT %s FROM %s %s %s %s", colsStr, this.tableName, whereStr, orderStr, limitStr)
	return strings.TrimSpace(sqlStr)
}

func (this *Dao) Find(where map[string]interface{}, order, limit string, selectCol ...string) (*sql.Rows, error) {
	whereS, whereVal := Map2Struct(where)
	if 0 == len(selectCol) {
		selectCol = append(selectCol, "*")
	}
	strSql := this.findSql(selectCol, whereS, order, limit)
	log.Println("Find sql:", strSql)
	err := this.Open()
	if err != nil {
		return nil, err
	}
	defer this.Close()

	return this.Query(strSql, whereVal...)
}

func (this *Dao) Scan(rows *sql.Rows, data map[string]interface{}, selectCol ...string) error {
	_, colsVal := Map2Struct(data)
	colNum := len(selectCol)
	if colNum > 1 && (colNum == 1 && selectCol[0] != "*") {
		tmp := []interface{}{}
		for _, v := range selectCol {
			tmp = append(tmp, data[v])
		}
		colsVal = tmp
	}

	if rows.Next() {
		rows.Scan(colsVal...)
	}

	return rows.Err()
}
