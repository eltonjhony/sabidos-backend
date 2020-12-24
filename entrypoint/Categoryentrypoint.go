package entrypoint

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sabidos/core/usecase/CategoryUseCase"
)

type CategoryEntrypointHandler struct {
	ObtainCategory CategoryUseCase.ObtainCategoryUseCaseProtocol
}

func NewCategoryEntrypointHandler(r *gin.RouterGroup, obtainCategory CategoryUseCase.ObtainCategoryUseCaseProtocol) {
	handler := &CategoryEntrypointHandler{
		ObtainCategory: obtainCategory,
	}

	r.GET("/categories", handler.FindAllCategories)

}

func (catergoryEntrypointHandler *CategoryEntrypointHandler) FindAllCategories(c *gin.Context) {

	categories, err := catergoryEntrypointHandler.ObtainCategory.GetAll(c.Request.Context())

	if err != nil {
		fmt.Println("Can't find Category", err)
		c.JSON(404, gin.H{"message": "Category not found"})
		return
	}

	c.JSON(200, gin.H{"categories": categories})
}
