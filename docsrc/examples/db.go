package examples

import (
	"github.com/theplant/osenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

var dbParamsString = osenv.Get("DB_PARAMS", "database connection string", "")

func ExampleDB() (r *gorm.DB) {
	if db != nil {
		return db
	}
	var err error
	db, err = gorm.Open(postgres.Open(dbParamsString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Logger.LogMode(logger.Info)
	r = db
	return
}
