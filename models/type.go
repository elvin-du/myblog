package models

/*
博客
*/
type Blog struct {
	Id         int
	Content    string
	Title      string
	CreateDate string
	TagId      int
	Comments   []Comment
}

/*
文章评论
*/
type Comment struct {
	Id         int
	IP         string
	Content    string
	CreateDate string
	BlogId     int
}

/*
文章标签
*/
type Tag struct {
	Id  int
	Tag string
}
