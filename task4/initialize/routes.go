package initialize

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/yhcui/web3study/task4/api/blog"
	"github.com/yhcui/web3study/task4/api/system"
	"github.com/yhcui/web3study/task4/global"
	"github.com/yhcui/web3study/task4/model/response"
)

func Routes() *gin.Engine {

	router := gin.Default()
	router.Use(ErrorHandler(), LoggerMiddleware())
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	sysRouter := router.Group("system").Use(ErrorHandler(), LoggerMiddleware())
	{
		sysRouter.POST("/register", system.Register)
		sysRouter.POST("/login", system.Login)
	}

	blogRouter := router.Group("blog").Use(ErrorHandler(), LoggerMiddleware(), JWTAuth())
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

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			response.FailWithMsg(err.Error(), c)
		}
	}
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		global.Logger.Info("request log",
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.URL.Path),
			slog.String("ip", c.ClientIP()),
			slog.Int("status", c.Writer.Status()),
			slog.Duration("latency", latency))

	}
}
