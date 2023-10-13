package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/qor5/docs/docsrc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8800"
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DB_PARAMS")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	go runAtMidNight(db)

	// @snippet_begin(HelloWorldMainSample)
	fmt.Println("Starting docs at :" + port)
	err = http.ListenAndServe(":"+port, docsrc.Mux("/"))
	if err != nil {
		panic(err)
	}
	// @snippet_end
}

func runAtMidNight(db *gorm.DB) {
	if os.Getenv("ENV") == "development" {
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
