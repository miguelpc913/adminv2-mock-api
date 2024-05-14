package dbHelpers

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDb() (db *gorm.DB, err error) {
	var userDb = os.Getenv("MYSQL_USER")
	var pass = os.Getenv("MYSQL_PASSWORD")
	var url = os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_PORT")
	var dbname = os.Getenv("MYSQL_DATABASE")
	var dsn = fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		userDb,
		pass,
		url,
		dbname,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}
