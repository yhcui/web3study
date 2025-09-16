package blog

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yhcui/web3study/task4/global"
	"github.com/yhcui/web3study/task4/model"
	"github.com/yhcui/web3study/task4/model/response"
)

func CreateBlog(c *gin.Context) {
	post := model.Posts{}
	value, _ := c.Get("userID")
	post.UserId = value.(uint)
	if err := c.ShouldBind(&post); err == nil {
		tx := global.SDB.Create(&post)
		if tx.RowsAffected > 0 {
			response.Ok(c)
		} else {
			response.Fail(c)
		}

	} else {
		response.Fail(c)
	}
}

func ListBlog(c *gin.Context) {
	posts := []model.Posts{}
	tx := global.SDB.Find(&posts)
	if tx.RowsAffected > 0 {
		fmt.Println("不为0")
	}
	response.OkWithData(posts, c)
}

func DetailBlog(c *gin.Context) {
	id, _ := c.GetQuery("id")
	post := model.Posts{}
	global.SDB.Where("id = ?", id).First(&post)
	response.OkWithData(post, c)
}

func UpdateBlog(c *gin.Context) {
	post := model.Posts{}
	err := c.ShouldBind(&post)
	if err != nil {
		response.Fail(c)
	} else {
		tx := global.SDB.Model(&post).Updates(post)
		if tx.RowsAffected > 0 {
			response.Ok(c)
		} else {
			response.Fail(c)
		}
	}
	c.JSON(200, gin.H{
		"message": "list",
	})
}

func DeleteBlog(c *gin.Context) {
	id, _ := c.GetPostFormMap("id")
	global.SDB.Delete(&model.Posts{}, id)
	response.OkWithMessage("删除成功", c)
}

func CommentCreate(c *gin.Context) {
	comment := model.Comments{}
	err := c.ShouldBind(&comment)
	if err != nil {
		response.Fail(c)
	} else {
		tx := global.SDB.Create(&comment)
		if tx.RowsAffected > 0 {
			response.Ok(c)
		} else {
			response.Fail(c)
		}
	}
}

func ListCommentByPostId(c *gin.Context) {
	postId, _ := c.GetQuery("postid")
	comments := []model.Comments{}
	global.SDB.Where("post_id = ?", postId).Find(&comments)
	response.OkWithData(comments, c)
}
