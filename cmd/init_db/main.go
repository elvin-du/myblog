package main

import (
	"log"
	"myblog/database"
)

func main() {
	InitDB()
}

type blogs struct {
	Id          int
	Title       string
	Content     string
	CreatedDate string
	TagId       int
}

func (this *blogs) fieldMap() map[string]interface{} {
	return map[string]interface{}{
		"id":           &this.Id,
		"title":        &this.Title,
		"content":      &this.Content,
		"created_date": &this.CreatedDate,
		"tag_id":       &this.TagId,
	}
}

func InitDB() {
	dao := database.NewDao("blogs")
	blog := &blogs{}
	rows, err := dao.Find(nil, "tag_id", "")
	dao.Scan(rows, blog.fieldMap())
	log.Println(err)
	log.Println(*blog)

	//insertData := make(map[string]interface{})
	//insertData["title"] = "japan"
	//insertData["content"] = "amarican"
	//insertData["created_date"] = time.Now().String()
	//insertData["tag_id"] = 996

	//where := make(map[string]interface{})
	//where["tag_id"] = 12345

	//_, err := dao.Insert(insertData)
	//if nil != err {
	//	log.Fatal(err)
	//}
}
