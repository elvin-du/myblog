package models

import(
	_"github.com/Go-SQL-Driver/MySQL"
)

type model struct{
}

func (db *model)CheckNamePsw(name,psw string)error{

}

func (db *model)AddUser(name,psw string)error{
}

func (db *model)EditBlogs(newBlog string)error{

}

func (db *model)AddBlogs(title,blog string)error{

}

func (db *model)DelBlogs(title string)error{
}

//query condition perhaps change,so agument type is interface
func (db *model)QueryBlogs(condition interface{})error{
}

func (db *model)AddComments(title, commtent string)error{
}

func (db *model)DelComments(title string)error{
}
