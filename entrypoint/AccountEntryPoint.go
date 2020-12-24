package entrypoint

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sabidos/core/usecase/AccountUseCase"
	"github.com/sabidos/entrypoint/model"
)

const NICKNAME_ALREADY_IN_USE_CODE = 1
const DATA_NOT_FOUND_CODE = 2
const SUCCESS = 3

type AccountEntrypointHandler struct {
	ObtainAccount   AccountUseCase.ObtainAccountUseCaseProtocol
	InsertAccount   AccountUseCase.InsertAccountUseCaseProtocol
	ValidateAccount AccountUseCase.ValidateAccountUseCaseProtocol
	UpdateAccount   AccountUseCase.UpdateAccountUseCaseProtocol
}

func NewAccountEntrypointHandler(r *gin.RouterGroup, obtainAccount AccountUseCase.ObtainAccountUseCaseProtocol,
	insertAcc AccountUseCase.InsertAccountUseCaseProtocol,
	validateAcc AccountUseCase.ValidateAccountUseCaseProtocol,
	updateAcc AccountUseCase.UpdateAccountUseCaseProtocol) {
	handler := &AccountEntrypointHandler{
		ObtainAccount:   obtainAccount,
		InsertAccount:   insertAcc,
		ValidateAccount: validateAcc,
		UpdateAccount:   updateAcc,
	}

	r.GET("/account/uid/:uid", handler.FindAccountByUid)
	r.GET("/account/nickname/:nickname", handler.FindAccountByNickName)
	r.POST("/account/validate", handler.validate)
	r.POST("/account/", handler.Create)
	r.PATCH("/account/:uid", handler.Update)
}

func (accountEntrypointHandler *AccountEntrypointHandler) FindAccountByUid(c *gin.Context) {
	uid := c.Param("uid")
	accounts, err := accountEntrypointHandler.ObtainAccount.GetByUid(c.Request.Context(), uid)

	if err != nil {
		fmt.Println("Can't find account", err)
		c.JSON(404, gin.H{"message": "Account not found"})
		return
	}

	c.JSON(200, gin.H{"account": accounts})
}

func (accountEntrypointHandler *AccountEntrypointHandler) FindAccountByNickName(c *gin.Context) {
	nickname := c.Param("nickname")
	accounts, err := accountEntrypointHandler.ObtainAccount.GetByNickname(c.Request.Context(), nickname)

	if err != nil {
		fmt.Println("Can't find account", err)
		c.JSON(404, gin.H{"code": DATA_NOT_FOUND_CODE, "message": "Account not found"})
		return
	}

	c.JSON(200, gin.H{"account": accounts})
}

func (accountEntrypointHandler *AccountEntrypointHandler) validate(c *gin.Context) {

	var accountModel model.AccountModel

	if err := c.ShouldBindJSON(&accountModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := accountEntrypointHandler.ValidateAccount.Validate(c.Request.Context(), accountModel.NickName, accountModel.Uid)

	if err == nil {
		fmt.Println("\nNickname already exists", err)
		c.JSON(400, gin.H{"code": NICKNAME_ALREADY_IN_USE_CODE, "message": "Nickname already exists"})
		return
	}

	c.JSON(200, gin.H{"code": SUCCESS, "message": "Nickname does not exists"})
}

func (accountEntrypointHandler *AccountEntrypointHandler) Create(c *gin.Context) {

	var accountModel model.AccountModel

	if err := c.ShouldBindJSON(&accountModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(accountModel)

	account, err := accountEntrypointHandler.InsertAccount.Insert(c.Request.Context(), accountModel)
	if err != nil {
		fmt.Println("Can't create account", err)
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, account)
}

func (accountEntrypointHandler *AccountEntrypointHandler) Update(c *gin.Context) {

	var updateAccountModel model.UpdateAccountModel

	if err := c.ShouldBindJSON(&updateAccountModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(updateAccountModel)

	uid := c.Param("uid")

	account, err := accountEntrypointHandler.UpdateAccount.Update(c.Request.Context(), uid, updateAccountModel)
	if err != nil {
		fmt.Println("Can't update account", err)
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, account)
}
