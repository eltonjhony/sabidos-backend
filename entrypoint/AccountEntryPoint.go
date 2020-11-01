package entrypoint

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sabidos/core/entity"
)

type AccountEntrypointHandler struct {
	ObtainAccount entity.ObtainAccountUseCase
	InsertAccount entity.InsertAccountUseCase
}

func NewAccountEntrypointHandler(r *gin.RouterGroup, obtainAccount entity.ObtainAccountUseCase, insertAcc entity.InsertAccountUseCase) {
	handler := &AccountEntrypointHandler{
		ObtainAccount: obtainAccount,
		InsertAccount: insertAcc,
	}
	r.GET("/account/:id", handler.FindAccount)
	r.POST("/account/", handler.Create)
}

func (accountEntrypointHandler *AccountEntrypointHandler) FindAccount(c *gin.Context) {
	accounts, _ := accountEntrypointHandler.ObtainAccount.Get(c.Request.Context(), c.Param("id"))
	c.JSON(200, gin.H{"data": accounts})
}

func (accountEntrypointHandler *AccountEntrypointHandler) Create(c *gin.Context) {

	var account entity.Account

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(account)

	err := accountEntrypointHandler.InsertAccount.Insert(c.Request.Context(), account)
	if err != nil {
		fmt.Println("Can't create account", err)
		return
	}

	c.JSON(200, gin.H{"data": account})
}
