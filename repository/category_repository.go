package repository

import "golang-unit-test/entity"

type CategoryRepository interface {
	FindById(is string) *entity.Category
}
