package entrypoint

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sabidos/core/entity"
	"go.mongodb.org/mongo-driver/bson"
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
	r.GET("/account/uid/:uid", handler.FindAccountByUid)
	r.GET("/account/nickname/:nickname", handler.FindAccountByNickName)
	r.POST("/account/", handler.Create)
}

func (accountEntrypointHandler *AccountEntrypointHandler) FindAccountByUid(c *gin.Context) {
	bfilter := bson.M{"uid": c.Param("uid")}
	accounts, err := accountEntrypointHandler.ObtainAccount.Get(c.Request.Context(), bfilter)

	if err != nil {
		fmt.Println("Can't find account", err)
		c.JSON(404, gin.H{"message": "Account not found"})
		return
	}

	c.JSON(200, gin.H{"account": accounts})
}

func (accountEntrypointHandler *AccountEntrypointHandler) FindAccountByNickName(c *gin.Context) {
	bfilter := bson.M{"nickname": c.Param("nickname")}
	accounts, err := accountEntrypointHandler.ObtainAccount.Get(c.Request.Context(), bfilter)

	if err != nil {
		fmt.Println("Can't find account", err)
		c.JSON(404, gin.H{"message": "Account not found"})
		return
	}

	c.JSON(200, gin.H{"account": accounts})
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
		c.JSON(400, gin.H{"message": "Can't create account with the providede params"})
		return
	}

	c.JSON(200, account)
}
