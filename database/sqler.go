package database

//import (
//	"fmt"
//	"strings"
//)

//type Sqler interface {
//	TableName() string
//	Columns() []string
//	SelectCols() []string // 需要查询哪些字段
//	Where() string
//	Order() string
//	Limit() string
//}

//func InsertSql(sqler Sqler) string {
//	columns := sqler.Columns()
//	columnStr := "`" + strings.Join(columns, "`,`") + "`"
//	placeHolder := strings.Repeat("?,", len(columns))
//	sql := fmt.Sprintf("INSERT INTO `%s`(%s) VALUES(%s)", sqler.TableName(), columnStr, placeHolder[:len(placeHolder)-1])
//	return strings.TrimSpace(sql)
//}

//func SelectSql(sqler Sqler) string {
//	where := sqler.Where()
//	if where != "" {
//		where = "WHERE " + where
//	}
//	order := sqler.Order()
//	if order != "" {
//		order = "ORDER BY " + order
//	}
//	limit := sqler.Limit()
//	if limit != "" {
//		limit = "LIMIT " + limit
//	}

//	ss := sqler.SelectCols()
//	selectCols := "`" + strings.Join(ss, "`,`") + "`"
//	sql := fmt.Sprintf("SELECT %s FROM `%s` %s %s %s", selectCols, sqler.TableName(), where, order, limit)
//	return strings.TrimSpace(sql)
//}
