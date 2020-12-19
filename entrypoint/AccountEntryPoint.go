package entrypoint

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sabidos/core/entity"
	"github.com/sabidos/entrypoint/model"
	"go.mongodb.org/mongo-driver/bson"
)

const NICKNAME_ALREADY_IN_USE_CODE = 1
const DATA_NOT_FOUND_CODE = 2
const SUCCESS = 3

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
	r.POST("/account/validate", handler.validate)
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

	bfilter := bson.M{"$or": []bson.M{bson.M{"nickname": accountModel.NickName}, bson.M{"uid": accountModel.Uid}}}

	_, err := accountEntrypointHandler.ObtainAccount.Get(c.Request.Context(), bfilter)

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

	account := entity.Account{
		Uid: accountModel.Uid, 
		Name: accountModel.Name, 
		NickName: accountModel.NickName,
		Avatar: entity.Avatar{
			Id: accountModel.DefaultAvatarId,
		},
		Email: accountModel.Email, 
		IsAnonymous: accountModel.IsAnonymous, 
		Phone: accountModel.Phone,
	}

	account, err := accountEntrypointHandler.InsertAccount.Insert(c.Request.Context(), account)
	if err != nil {
		fmt.Println("Can't create account", err)
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, account)
}
