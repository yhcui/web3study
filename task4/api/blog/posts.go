package blog

import (
	"fmt"
	"time"

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
		qpost := model.Posts{}
		global.SDB.Where("id = ?", post.ID).First(&qpost)

		value, _ := c.Get("userID")
		userId := value.(uint)

		if qpost.UserId != userId {
			response.FailWithMsg("不是你的，你不能改", c)
			return
		}
		post.Title = qpost.Title
		post.Content = qpost.Content
		post.UpdatedAt = time.Now()
		post.UserId = userId
		tx := global.SDB.Model(&post).Updates(post)
		if tx.RowsAffected > 0 {
			response.Ok(c)
		} else {
			response.Fail(c)
		}
	}
}

func DeleteBlog(c *gin.Context) {
	//id, _ := c.GetPostFormMap("id")
	id, _ := c.GetPostForm("id")

	qpost := model.Posts{}
	global.SDB.Where("id = ?", id).First(&qpost)

	value, _ := c.Get("userID")
	userId := value.(uint)

	if qpost.UserId != userId {
		response.FailWithMsg("不是你的，你不能改", c)
		return
	}

	global.SDB.Delete(&model.Posts{}, id)
	response.OkWithMessage("删除成功", c)
}

func CommentCreate(c *gin.Context) {

	value, _ := c.Get("userID")
	userId := value.(uint)

	comment := model.Comments{}
	err := c.ShouldBind(&comment)
	if err != nil {
		response.Fail(c)
	} else {
		comment.UserId = userId
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
	global.SDB.Where("posts_id = ?", postId).Find(&comments)
	response.OkWithData(comments, c)
}
