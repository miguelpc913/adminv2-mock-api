package db

import (
	"fmt"

	dbHelpers "github.com/tiqueteo/adminv2-mock-api/db/utils"
	"gorm.io/gorm"
)

func InitDB() (db *gorm.DB, err error) {
	db, err = dbHelpers.ConnectDb()
	if err != nil {
		return nil, err
	}
	dbHelpers.MigrateModels(db)
	dbHelpers.SeedEntities(db)
	fmt.Println("Successfully connected!", db)
	return db, nil
}
