package CategoryUseCase

import (
	"context"
	"fmt"

	"github.com/sabidos/core/entity"
	"go.mongodb.org/mongo-driver/bson"
)

type ObtainCategoryUseCase struct {
	categoryRepository entity.CategoryDataProvider
}

func NewObtainCategoryUsecase(a entity.CategoryDataProvider) entity.ObtainCategoryUseCase {
	return &ObtainCategoryUseCase{
		categoryRepository: a,
	}
}

func (a *ObtainCategoryUseCase) Get(c context.Context, filter bson.M) (av []entity.Category, err error) {

	av, err = a.categoryRepository.Get(c, filter)
	if err != nil {
		fmt.Printf("Error %s ", err)
		return av, err
	}
	return av, err
}
