package initialize

import (
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/yhcui/web3study/task4/api/blog"
	"github.com/yhcui/web3study/task4/api/system"
	"github.com/yhcui/web3study/task4/global"
	"github.com/yhcui/web3study/task4/model/response"
)

func Routers() *gin.Engine {
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
		blogRouter.GET("/comment/listbypostid", blog.ListCommentByPostId)

	}

	return router
}
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			global.Logger.Info("error log",
				slog.String("err", err.Error()))
			response.FailWithMsg(err.Error(), c)
		}
	}
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		ignoreUrl := []string{"/blog/blog/list", "/blog/blog/detail", "/blog/comment/listbypostid"}
		path := c.Request.URL.Path
		slog.Info("url", slog.String("path", path))
		for _, s := range ignoreUrl {
			if s == path {
				c.Next()
				return
			}
		}

		auth := c.GetHeader("Authorization")

		if strings.TrimSpace(auth) == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		tokenString := strings.TrimPrefix(auth, "Bearer ")

		jwtCustomClaims := system.JwtCustomClaims{}
		jwt.ParseWithClaims(tokenString, &jwtCustomClaims, func(token *jwt.Token) (interface{}, error) {
			return "$#@$#54$2qrweqrew", nil
		})

		if jwtCustomClaims.ID > 0 {
			c.Set("userID", jwtCustomClaims.ID)
			c.Set("name", jwtCustomClaims.Name)
			c.Next()
		} else {
			response.FailWithMsg("invalid token", c)
			c.Abort()

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
