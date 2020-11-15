package entrypoint

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sabidos/core/entity"
	"go.mongodb.org/mongo-driver/bson"
)

type AvatarEntrypointHandler struct {
	ObtainAvatar entity.ObtainAvatarUseCase
}

func NewAvatarEntrypointHandler(r *gin.RouterGroup, obtainAvatar entity.ObtainAvatarUseCase) {
	handler := &AvatarEntrypointHandler{
		ObtainAvatar: obtainAvatar,
	}

	r.GET("/avatars", handler.FindAvatar)

}

func (avatarEntrypointHandler *AvatarEntrypointHandler) FindAvatar(c *gin.Context) {
	bfilter := bson.M{}
	avatars, err := avatarEntrypointHandler.ObtainAvatar.Get(c.Request.Context(), bfilter)

	if err != nil {
		fmt.Println("Can't find Avatar", err)
		c.JSON(404, gin.H{"message": "Avatar not found"})
		return
	}

	c.JSON(200, gin.H{"avatars": avatars})
}
