package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"unit-testing/entity"
	"unit-testing/repository"
)

// * Create ONE mock repo and ONE service that uses it.
// * The tests will "script" this mock's behavior using .On(...).Return(...)
var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

// * Test: when repo says "not found" (returns nil), service should return an error.
func TestCategoryService_GetNotFound(t *testing.T) {
	// * Script the mock: "If FindById('1') is called, return nil"
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	// * Call the service (this will call the mock under the hood)
	category, err := categoryService.Get("1")

	// * Expect: no category, but there IS an error.
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

// * Test: when repo returns a real category, service should pass it through without error.
func TestCategoryService_GetSuccess(t *testing.T) {
	// * Prepare a fake category that the mock will return
	category := entity.Category{
		Id:   "2",
		Name: "Laptop",
	}

	// * Script the mock: "If FindById('2') is called, return this category value"
	categoryRepository.Mock.On("FindById", "2").Return(category)

	// * Call the service
	result, err := categoryService.Get("2")

	// * Expect: no error, result is not nil, and fields match
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}
