package entrypoint

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sabidos/core/usecase/AvatarUseCase"
)

type AvatarEntrypointHandler struct {
	ObtainAvatar AvatarUseCase.ObtainAvatarUseCaseProtocol
}

func NewAvatarEntrypointHandler(r *gin.RouterGroup, obtainAvatar AvatarUseCase.ObtainAvatarUseCaseProtocol) {
	handler := &AvatarEntrypointHandler{
		ObtainAvatar: obtainAvatar,
	}

	r.GET("/avatars", handler.FindAvatar)
}

func (avatarEntrypointHandler *AvatarEntrypointHandler) FindAvatar(c *gin.Context) {

	avatars, err := avatarEntrypointHandler.ObtainAvatar.GetAll(c.Request.Context())

	if err != nil {
		fmt.Println("Can't find Avatar", err)
		c.JSON(404, gin.H{"message": "Avatar not found"})
		return
	}

	c.JSON(200, gin.H{"avatars": avatars})
}
