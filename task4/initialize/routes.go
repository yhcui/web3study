package initialize

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/yhcui/web3study/task4/api/blog"
	"github.com/yhcui/web3study/task4/api/system"
)

func Routes() *gin.Engine {

	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	sysRouter := router.Group("system")
	{
		sysRouter.POST("/register", system.Register)
		sysRouter.POST("/login", system.Login)
	}

	blogRouter := router.Group("blog").Use(JWTAuth())
	{
		blogRouter.POST("/blog/create", blog.CreateBlog)
		blogRouter.GET("/blog/list", blog.ListBlog)
		blogRouter.GET("/blog/detail", blog.DetailBlog)
		blogRouter.POST("/blog/update", blog.UpdateBlog)
		blogRouter.POST("/blog/delete", blog.DeleteBlog)
		blogRouter.POST("/comment/create", blog.CommentCreate)
		blogRouter.POST("/comment/listbypostid", blog.ListCommentByPostId)

	}

	return router
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")

		if strings.TrimSpace(auth) == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		tokenString := strings.TrimPrefix(auth, "Bearer ")

		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userID", claims["userID"])
			c.Set("roles", claims["roles"])
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		}
	}
}
