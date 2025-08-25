package service

import (
	"errors"
	"unit-testing/entity"
	"unit-testing/repository"
)

// * Service layer = business logic.
// * It depends on the repository INTERFACE, so we can swap real DB / mock easily.
type CategoryService struct {
	Repository repository.CategoryRepository
}

// * Get tries to fetch a category by id.
// * If the repo returns nil → we translate that to a "not found" error.
// * If not nil → we just pass the category back up.
func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("Category not found!")
	} else {
		return category, nil
	}
}
