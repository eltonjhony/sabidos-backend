package entrypoint

import (
	"github.com/gin-gonic/gin"
	"github.com/sabidos/core/entity"
)

type AccountEntrypointHandler struct {
	ObtainAccount entity.ObtainAccountUseCase
}

func NewAccountEntrypointHandler(r *gin.RouterGroup, us entity.ObtainAccountUseCase) {
	handler := &AccountEntrypointHandler{
		ObtainAccount: us,
	}
	r.GET("/account", handler.FindAccounts)
}

func (accountEntrypointHandler *AccountEntrypointHandler) FindAccounts(c *gin.Context) {
	accounts, _ := accountEntrypointHandler.ObtainAccount.Fetch(c.Request.Context())
	c.JSON(200, gin.H{"data": accounts})
}
