package service

import (
	"errors"
	"golang-unit-test/entity"
	"golang-unit-test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	result := service.Repository.FindById(id)
	if result == nil {
		return nil, errors.New("catgory not Found")

	} else {
		return result, nil
	}
}
