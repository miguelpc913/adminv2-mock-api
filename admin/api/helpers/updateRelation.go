package helpers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

func UpdateRelation[J any, T any](r *http.Request, model J, modelToAssociate T, db *gorm.DB, associationName string) error {
	req := []int{}
	id := chi.URLParam(r, "id")
	err := db.First(&model, id).Error
	if err != nil {
		return errors.New("there is no entity with that id")
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return errors.New("invalid request")
	}
	modelSet := []T{}
	for _, id := range req {
		relationModel := modelToAssociate
		if err := db.First(&relationModel, id).Error; err != nil {
			return errors.New("ids are not valid")
		}
		modelSet = append(modelSet, relationModel)
	}

	db.Model(&model).Association(associationName).Replace(modelSet)
	return nil
}
