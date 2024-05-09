package examples_admin

import (
	"context"
	"os"

	"github.com/qor5/admin/v3/activity"
	"github.com/qor5/admin/v3/presets"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewActivitySample() {
	// @snippet_begin(NewActivitySample)
	presetsBuilder := presets.New()
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_PARAMS")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	activityBuilder := activity.New(db)
	activityBuilder.Install(presetsBuilder)
	// @snippet_end

	// @snippet_begin(ActivityRegisterPresetsModelsSample)
	type Product struct {
		Title string
		Code  string
		Price float64
	}
	productModel := presetsBuilder.Model(&Product{})

	activityBuilder.RegisterModel(productModel).EnableActivityInfoTab().AddKeys("Title").AddIgnoredFields("Code").SkipDelete()
	// @snippet_end

	// @snippet_begin(ActivityRecordLogSample)
	currentCtx := context.WithValue(context.Background(), activity.CreatorContextKey, "user1")

	activityBuilder.AddRecords("Publish", currentCtx, &Product{Title: "Product 1", Code: "P1", Price: 100})

	activityBuilder.AddRecords("Update Price", currentCtx, &Product{Title: "Product 1", Code: "P1", Price: 200})
	// @snippet_end
}
