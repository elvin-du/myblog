package models

import(
	_"github.com/Go-SQL-Driver/MySQL"
	"html"
	"log"
	"errors"
	"database/sql"
)

type Model struct{
}

type Condition struct{
	con_type string //VALUE:"byDate","byTypeID","byTitle"
	date string //which year
	title string
	type_id int
}

func (m *Model)CheckNamePsw(name,psw string)error{
	db, err := sql.Open("mysql", "root:dumx@tcp(localhost:3306)/myblog?charset=utf8")
	if nil != err{
		log.Print(err)
		return err
	}
	defer db.Close()
	querySql := "select 1 from myblog.users WHERE name = ' " + html.EscapeString(name) + "'"
	rows, err := db.Query(querySql)
	if nil != err{
		log.Print(err)
		return err
	}
	if rows.Next(){
		return errors.New("user " + name + "exsited")
	}

	return nil
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
		return errors.New("user " + name + "exsited")
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

func (m *Model)EditBlogs(newBlog string)error{
	return nil
}

func (m *Model)AddBlogs(title,blog string)error{
	db, err := sql.Open("mysql", "root:dumx@tcp(localhost:3306)/myblog?charset=utf8")
	if nil != err{
		log.Print(err)
		return err
	}
	defer db.Close()
	querySql := "select name from myblog.blogs WHERE title = ' " + html.EscapeString(title) + "'"
	rows, err := db.Query(querySql)
	if nil != err{
		log.Print(err)
		return err
	}
	if rows.Next(){
		return errors.New("title exsited")
	}

	insertSql := "INSERT myblog.users SET name=?, password=?"
	stmt, err := db.Prepare(insertSql)
	if nil != err{
		log.Print(err)
		return err
	}
	defer stmt.Close()

//	_, err = stmt.Exec(username, password)
	if nil != err{
		log.Print(err)
		return err
	}

	return nil
}

//query condition perhaps change,so agument type is interface
func (m *Model)QueryBlogs(condition interface{})error{
	db, err := sql.Open("mysql", "root:dumx@tcp(localhost:3306)/myblog?charset=utf8")
	if nil != err{
		log.Print(err)
		return err
	}
	defer db.Close()
	querySql := "select name from myblog.blogs WHERE title = ' " //+ html.EscapeString(title) + "'"
	rows, err := db.Query(querySql)
	if nil != err{
		log.Print(err)
		return err
	}
	if rows.Next(){
		return errors.New("title exsited")
	}

	insertSql := "INSERT myblog.users SET name=?, password=?"
	stmt, err := db.Prepare(insertSql)
	if nil != err{
		log.Print(err)
		return err
	}
	defer stmt.Close()

	//_, err = stmt.Exec(username, password)
	if nil != err{
		log.Print(err)
		return err
	}

	return nil
}

func (m *Model)DelBlogs(title string)error{
	return nil
}

func (m *Model)AddComments(title, commtent string)error{
	return nil
}

func (m *Model)DelComments(title string)error{
	return nil
}
