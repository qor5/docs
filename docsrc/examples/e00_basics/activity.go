package e00_basics

import (
	"context"
	"os"

	"github.com/qor5/admin/activity"
	"github.com/qor5/admin/presets"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewActivitySample() {
	// @snippet_begin(NewActivitySample)
	presetsBuilder := presets.New()
	db, err := gorm.Open(postgres.Open(os.Getenv("DBURL")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	activityBuilder := activity.New(presetsBuilder, db)
	// @snippet_end

	// @snippet_begin(ActivityRegisterPresetsModelsSample)
	type Product struct {
		Title string
		Code  string
		Price float64
	}
	productModel := presetsBuilder.Model(&Product{})

	activityBuilder.RegisterModel(productModel).UseDefaultTab().AddKeys("Title").AddIgnoredFields("Code").SkipDelete()
	// @snippet_end

	// @snippet_begin(ActivityRecordLogSample)
	currentCtx := context.WithValue(context.Background(), activity.CreatorContextKey, "user1")

	activityBuilder.AddRecords("Publish", currentCtx, &Product{Title: "Product 1", Code: "P1", Price: 100})

	activityBuilder.AddRecords("Update Price", currentCtx, &Product{Title: "Product 1", Code: "P1", Price: 200})
	// @snippet_end
}
