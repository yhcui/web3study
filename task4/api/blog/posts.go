package blog

import "github.com/gin-gonic/gin"

func CreateBlog(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "list",
	})
}

func ListBlog(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "list",
	})
}

func DetailBlog(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "list",
	})
}

func UpdateBlog(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "list",
	})
}

func DeleteBlog(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "list",
	})
}

func CommentCreate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "list",
	})
}

func ListCommentByPostId(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "list",
	})
}
