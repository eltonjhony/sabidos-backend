package entrypoint

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sabidos/core/entity"
	"go.mongodb.org/mongo-driver/bson"
)

type CategoryEntrypointHandler struct {
	ObtainCategory entity.ObtainCategoryUseCase
}

func NewCategoryEntrypointHandler(r *gin.RouterGroup, obtainCategory entity.ObtainCategoryUseCase) {
	handler := &CategoryEntrypointHandler{
		ObtainCategory: obtainCategory,
	}

	r.GET("/categories", handler.FindCategory)

}

func (catergoryEntrypointHandler *CategoryEntrypointHandler) FindCategory(c *gin.Context) {
	bfilter := bson.M{}
	categories, err := catergoryEntrypointHandler.ObtainCategory.Get(c.Request.Context(), bfilter)

	if err != nil {
		fmt.Println("Can't find Category", err)
		c.JSON(404, gin.H{"message": "Category not found"})
		return
	}

	c.JSON(200, gin.H{"categories": categories})
}
