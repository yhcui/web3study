package system

import "github.com/gin-gonic/gin"

func Register(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "REGISTER",
	})
}

func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "LOGIN",
	})
}
