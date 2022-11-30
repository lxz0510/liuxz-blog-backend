package articles

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"liuxz-blog-backend/pkg/common/models"
)

type handler struct {
	DB *gorm.DB
}

//注册路由
func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	//init
	h := &handler{
		DB: db,
	}

	//注册路由
	routers := r.Group("/articles")

	//获取文章列表
	routers.GET("/", h.GetArticles)

	////新增
	//routers.POST("/", h.AddArticle)

	////更新
	//routers.PUT("/:id", h.UpdateArticle)

	//获取一篇文章
	routers.GET("/:id", h.GetArticle)
	//TODO:删除文章
	addFirstBook(db)
}

func addFirstBook(db *gorm.DB) {
	h := &handler{
		DB: db,
	}
	var article = models.Article{}
	var count int64 = 0
	h.DB.Model(&models.Article{}).Where("id = ?", 1).Count(&count)

	if count == 0 {
		article.Title = "War and Peace"
		article.Content = "Content"
		article.Description = "War and Peace is a literary work mixed with chapters on history and philosophy by the Russian author Leo Tolstoy. It was first published serially, then published in its entirety in 1869."
		h.DB.Create(&article)
	}
}
