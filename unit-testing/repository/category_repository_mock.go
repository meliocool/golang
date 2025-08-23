package repository

import (
	"github.com/stretchr/testify/mock"
	"unit-testing/entity"
)

// * CategoryRepositoryMock is a FAKE repository for tests.
// * The Mock field is Testify's "brain": it stores rules you set with .On(...).Return(...)
type CategoryRepositoryMock struct {
	Mock mock.Mock
}

// * FindById in the MOCK does not hit a real DB.
// * It simply returns whatever the test told it to return for this input id.
// * Flow:
// *   1) test says: Mock.On("FindById", "2").Return(category)
// *   2) service calls: Repository.FindById("2")
// *   3) we look up the scripted return via Mock.Called("2")
// *   4) we hand that value back to the service
func (repository CategoryRepositoryMock) FindById(id string) *entity.Category {
	// * Ask Testify: "What did the test say to return when FindById is called with this id?"
	arguments := repository.Mock.Called(id)

	// * If the test set Return(nil), Get(0) will be nil → simulate "not found".
	if arguments.Get(0) == nil {
		return nil
	} else {
		// * Otherwise, the test returned a VALUE (entity.Category).
		// * Cast it back to entity.Category, then return its POINTER to match the interface type.
		category := arguments.Get(0).(entity.Category)
		return &category
	}
	// * NOTE: This mock uses a VALUE receiver (func (repository CategoryRepositoryMock) ...).
	// * That means the mock struct is copied when the method is called.
	// * It still works here because you're always using the same instance,
	// * but for larger tests people often prefer a POINTER receiver to avoid surprises.
}
