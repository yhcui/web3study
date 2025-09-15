package main

import (
	"fmt"

	"gorm.io/gorm"
)

//func (c *Comment) BeforeDelete(tx *gorm.DB) (err error) {
//	tx.Statement.Set("delete_comments", *c)
//	fmt.Printf("BeforeDelete: 准备删除 ID 为 %d 的产品。\n", c.ID)
//	return nil
//}

// 在同一个事务中更新数据
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	post := Post{}
	tx.Where("id = ?", c.PostID).Take(&post)
	fmt.Println("post.CommentNum:", post.CommentNum)
	tx.Model(post).Updates(Post{CommentNum: post.CommentNum})
	if post.CommentNum == 0 {
		// 修改状态
		post.Status = "无状态"
		tx.Model(&post).Updates(Post{Status: "无状态"})
		return
	}
	return
}

func deleteComment(db *gorm.DB, id int) {
	cc := Comment{}
	db.First(&cc, id)
	db.Delete(&cc)
}
