package admin

import (
	"os"

	"github.com/qor5/docs/v3/cmd/qor5/admin-template/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() (db *gorm.DB) {
	var err error
	// Create database connection
	db, err = gorm.Open(postgres.Open(os.Getenv("DB_PARAMS")))
	if err != nil {
		panic(err)
	}

	// Set db log level
	db.Logger = db.Logger.LogMode(logger.Info)

	// Create data table in the database
	err = db.AutoMigrate(models.Post{})
	if err != nil {
		panic(err)
	}

	return
}
