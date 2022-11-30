package articles

import (
	"github.com/gin-gonic/gin"
	"liuxz-blog-backend/pkg/common/models"
	"net/http"
)

func (h handler) GetArticle(c *gin.Context) {
	id := c.Param("id")
	var article models.Article
	if result := h.DB.First(&article, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)

		return
	}
	c.JSON(http.StatusOK, &article)
}
