package AvatarUseCase

import (
	"context"
	"fmt"

	"github.com/sabidos/core/entity"
	"go.mongodb.org/mongo-driver/bson"
)

type ObtainAvatarUseCase struct {
	avatarRepository entity.AvatarDataProvider
}

func NewObtainAvatarUsecase(a entity.AvatarDataProvider) entity.ObtainAvatarUseCase {
	return &ObtainAvatarUseCase{
		avatarRepository: a,
	}
}

func (a *ObtainAvatarUseCase) Get(c context.Context, filter bson.M) (av []entity.Avatar, err error) {

	av, err = a.avatarRepository.Get(c, filter)
	if err != nil {
		fmt.Printf("Error %s ", err)
		return av, err
	}
	return av, err
}
