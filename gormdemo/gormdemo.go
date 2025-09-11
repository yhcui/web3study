package main

import (
	"fmt"
	"log"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string
	ArticleNum int
}

type Post struct {
	gorm.Model
	Title      string
	Body       string
	CommentNum uint
	UserID     uint
	Status     string
	User       User
	Comments   []Comment `tag:"grom:-"`
}

type Comment struct {
	gorm.Model
	Comment string
	PostID  uint
	Post    Post
	UserID  uint
	User    User
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	user := User{}
	tx.Where("id = ?", p.UserID).Take(&user)
	//user.ArticleNum++
	tx.Model(&user).UpdateColumn("article_num", user.ArticleNum+1)
	fmt.Printf("userAfterCreate")
	return
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	gormdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(reflect.TypeOf(gormdb))
	if err != nil {
		log.Fatal(err)
	}
	gormdb.AutoMigrate(&User{}, &Post{}, &Comment{})
	//posts := queryByUserId(gormdb, 1)
	//post := queryMostComment(gormdb)
	//
	//fmt.Println(post.CommentNum)
	//fmt.Println("============")
	//fmt.Println(post.Title)
	//createPost(gormdb)
	deleteComment(gormdb, 2)

}
func createPost(db *gorm.DB) {
	post := Post{Title: "新的一个", Body: "新的一个", CommentNum: 0, UserID: 1}
	tx := db.Create(&post)
	fmt.Println("新增:", tx.RowsAffected)
}
func queryByUserId(db *gorm.DB, userId uint) (posts []Post) {
	//var users []User
	//db.Where("id = ?", userId).Find(&users)
	//for _, u := range users {
	//	fmt.Println(u.Name)
	//}
	//var posts []Post
	db.Where("user_id = ?", userId).Find(&posts)
	for _, post := range posts {
		//fmt.Println(post.Title)
		db.Where("post_id = ?", post.ID).Find(&post.Comments)
		//for _, comment := range post.Comments {
		//	fmt.Println(comment)
		//}
	}
	return
}

func queryMostComment(db *gorm.DB) (post Post) {
	db.Order("comment_num desc").First(&post)
	return
}
