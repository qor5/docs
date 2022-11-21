package admin

import (
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() (db *gorm.DB) {
	var err error

	db, err = gorm.Open(postgres.Open(os.Getenv("DB_PARAMS")))
	if err != nil {
		panic(err)
	}

	db.Logger = db.Logger.LogMode(logger.Info)

	if err = db.AutoMigrate(); err != nil {
		panic(err)
	}

	return
}
