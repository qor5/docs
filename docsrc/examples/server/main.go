package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/qor5/docs/v3/docsrc/examples/mux_admin"
	"github.com/theplant/osenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbParamsString = osenv.Get("DB_PARAMS", "database connection string", "")
	port           = osenv.Get("PORT", "The port to serve on", "8800")
	envString      = osenv.Get("ENV", "environment flag", "development")
)

func main() {
	db, err := gorm.Open(postgres.Open(dbParamsString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	go runAtMidNight(db)

	// @snippet_begin(HelloWorldMuxSample1)
	mux := http.NewServeMux()
	// @snippet_end

	// @snippet_begin(HelloWorldMainSample)
	fmt.Println("Starting docs at :" + port)
	err = http.ListenAndServe(":"+port, mux_admin.Mux(mux, "/"))
	if err != nil {
		panic(err)
	}
	// @snippet_end
}

func runAtMidNight(db *gorm.DB) {
	if envString == "development" {
		return
	}

	t := time.Tick(time.Hour)
	for range t {
		if time.Now().Hour() == 0 {
			truncateAllTables(db)
		}
	}
}

func truncateAllTables(db *gorm.DB) {
	if err := db.Exec(`DO
$do$
BEGIN
    EXECUTE
   (SELECT 'TRUNCATE TABLE ' || string_agg(oid::regclass::text, ', ') || ' CASCADE'
    FROM   pg_class
    WHERE  relkind = 'r'  -- only tables
    AND    relnamespace = 'public'::regnamespace
   );
END
$do$;`).
		Error; err != nil {
		panic(err)
	}
}
