package dbHelpers

import (
	"fmt"

	"gorm.io/gorm"
)

func InitDB(shouldSeed bool) (db *gorm.DB, err error) {
	db, err = ConnectDb()
	if err != nil {
		return nil, err
	}
	MigrateModels(db)
	if shouldSeed {
		SeedEntities(db)
	}
	fmt.Println("Successfully connected!", db)
	return db, nil
}
