package helpers

import (
	"net/http"

	"github.com/go-chi/chi"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetById[T any](model T, r *http.Request, db *gorm.DB) error {
	id := chi.URLParam(r, "id")
	return db.Preload(clause.Associations).First(model, id).Error
}
