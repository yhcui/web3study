package system

import (
	"github.com/gin-gonic/gin"
	"github.com/yhcui/web3study/task4/global"
	"github.com/yhcui/web3study/task4/model"
	"github.com/yhcui/web3study/task4/model/response"
)

func Register(c *gin.Context) {
	user := model.User{}
	err := c.ShouldBind(&user)
	if err != nil {
		response.Fail(c)
		return
	}

	users := []model.User{}
	global.SDB.Select("name = ?", user.Name).Find(&users)
	if len(users) > 0 {
		response.FailWithMsg("用户名已被注册", c)
		return
	}

	global.SDB.Create(&user)
	c.JSON(200, gin.H{
		"message": "REGISTER",
	})
}

func Login(c *gin.Context) {
	user := model.User{}
	err := c.ShouldBind(&user)
	if err != nil {
		response.FailWithMsg("系统异常", c)
		return
	}

	c.JSON(200, gin.H{
		"message": "LOGIN",
	})
}
