package service

import (
	"belajar-unit-test-mock/entity"
	"belajar-unit-test-mock/repository"
	"errors"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	// jika data category tidak ada maka return nil
	if category == nil {
		return nil, errors.New("Category Not Found")
	} else {
		return category, nil
	}
}
