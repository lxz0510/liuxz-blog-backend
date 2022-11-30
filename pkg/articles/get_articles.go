package articles

import (
	"github.com/gin-gonic/gin"
	"liuxz-blog-backend/pkg/common/models"
	"net/http"
)

func (h handler) GetArticles(c *gin.Context) {
	var articles []models.Article
	if result := h.DB.Find(&articles); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)

		return
	}
	c.JSON(http.StatusOK, &articles)
}
