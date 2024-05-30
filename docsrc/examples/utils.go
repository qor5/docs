package examples

import (
	"reflect"
	"runtime"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/theplant/osenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

var dbParamsString = osenv.Get("DB_PARAMS", "database connection string", "user=docs password=docs dbname=docs sslmode=disable host=localhost port=6532 TimeZone=Asia/Tokyo")

func ExampleDB() (r *gorm.DB) {
	if db != nil {
		return db
	}
	var err error
	db, err = gorm.Open(postgres.Open(dbParamsString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Logger = db.Logger.LogMode(logger.Info)
	r = db
	return
}

func SampleURLPathByFunc(v interface{}) string {
	funcNameWithPkg := runtime.FuncForPC(reflect.ValueOf(v).Pointer()).Name()
	segs := strings.Split(funcNameWithPkg, ".")
	return "/samples/" + strcase.ToKebab(segs[len(segs)-1])
}
