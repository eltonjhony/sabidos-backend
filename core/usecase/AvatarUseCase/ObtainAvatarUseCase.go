package AvatarUseCase

import (
	"context"
	"fmt"

	"github.com/sabidos/core/entity"
)

type ObtainAvatarUseCase struct {
	avatarRepository entity.AvatarDataProvider
}

func NewObtainAvatarUsecase(a entity.AvatarDataProvider) ObtainAvatarUseCaseProtocol {
	return &ObtainAvatarUseCase{
		avatarRepository: a,
	}
}

func (a *ObtainAvatarUseCase) GetAll(c context.Context) (av []entity.Avatar, err error) {

	av, err = a.avatarRepository.GetAll(c)
	if err != nil {
		fmt.Printf("Error %s ", err)
		return av, err
	}
	return av, err
}
