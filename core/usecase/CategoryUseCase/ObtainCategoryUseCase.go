package CategoryUseCase

import (
	"context"
	"fmt"

	"github.com/sabidos/core/entity"
)

type ObtainCategoryUseCase struct {
	categoryRepository entity.CategoryDataProvider
}

func NewObtainCategoryUsecase(a entity.CategoryDataProvider) ObtainCategoryUseCaseProtocol {
	return &ObtainCategoryUseCase{
		categoryRepository: a,
	}
}

func (a *ObtainCategoryUseCase) GetAll(c context.Context) (av []entity.Category, err error) {

	av, err = a.categoryRepository.GetAll(c)
	if err != nil {
		fmt.Printf("Error %s ", err)
		return av, err
	}
	return av, err
}
