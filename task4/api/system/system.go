package system

import (
	"github.com/gin-gonic/gin"
	"github.com/yhcui/web3study/task4/global"
	"github.com/yhcui/web3study/task4/model"
	"github.com/yhcui/web3study/task4/model/response"
	"golang.org/x/crypto/bcrypt"
)

//const (
//	MinCost     int = 4
//	MaxCost     int = 31
//	DefaultCost int = 10
//)

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

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(password)
	global.SDB.Create(&user)
	response.OkWithMessage("注册成功", c)
}

func Login(c *gin.Context) {
	user := model.User{}
	err := c.ShouldBind(&user)
	if err != nil {
		response.FailWithMsg("系统异常", c)
		return
	}

	userdb := model.User{}
	global.SDB.Select("name = ?", user.Name).First(&userdb)
	if userdb.ID == 0 {
		response.FailWithMsg("没有该有户", c)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(userdb.Password), []byte(user.Password))
	if err != nil {
		response.FailWithMsg("用户密码错误", c)
		return
	}
	// 生成JWT
	c.JSON(200, gin.H{
		"message": "LOGIN",
	})
}
