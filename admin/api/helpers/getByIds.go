package helpers

import (
	"errors"

	"gorm.io/gorm"
)

func GetByIds[T any](modelArray []T, ids []int, db *gorm.DB) ([]T, error) {
	for _, id := range ids {
		var model T
		if err := db.First(&model, id).Error; err != nil {
			return nil, errors.New("invalid ids")
		}
		modelArray = append(modelArray, model)
	}
	return modelArray, nil
}
