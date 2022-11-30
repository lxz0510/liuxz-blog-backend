package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"liuxz-blog-backend/pkg/articles"
	"liuxz-blog-backend/pkg/common/db"
)

const (
	host     = "localhost:5432"
	user     = "postgres"
	password = "110805"
	dbname   = "liuxz_blog_db"
)

func main() {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s/%s",
		user, password, host, dbname)
	r := gin.Default()
	h := db.Init(psqlInfo)
	r.Use(cors.Default())
	articles.RegisterRoutes(r, h)

	//test
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run("0.0.0.0:" + "8080") // 监听并在 0.0.0.0:8080 上启动服务
}
