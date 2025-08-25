package repository

import "unit-testing/entity"

// * Repository layer = "how we fetch/save data" contract.
// * This is an INTERFACE so the service can depend on behavior, not a concrete DB.
// * Return type is *entity.Category so we can:
// *   - avoid copying the struct
// *   - return nil to mean "not found"
type CategoryRepository interface {
	// * Input: id (string), Output: pointer to Category (or nil if not found)
	FindById(id string) *entity.Category
}
