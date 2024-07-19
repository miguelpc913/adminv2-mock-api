package dbHelpers

import (
	"fmt"

	"gorm.io/gorm"
)

func InitDB() (db *gorm.DB, err error) {
	db, err = ConnectDb()
	if err != nil {
		return nil, err
	}
	MigrateModels(db)
	SeedEntities(db)
	fmt.Println("Successfully connected!", db)
	return db, nil
}
