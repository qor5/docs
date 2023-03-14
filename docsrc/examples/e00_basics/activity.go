package e00_basics

import (
	"context"

	"github.com/qor5/admin/activity"
	"github.com/qor5/admin/presets"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewActivitySample() {
	// @snippet_begin(NewActivitySample)
	presetsBuilder := presets.New()
	db, err := gorm.Open(sqlite.Open("/tmp/activity.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	activity.New(presetsBuilder, db)
	// @snippet_end
}

func ActivityRegisterModelsSample() {
	// @snippet_begin(ActivityRegisterPresetsModelsSample)
	presetsBuilder := presets.New()
	db, err := gorm.Open(sqlite.Open("/tmp/activity.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	type Product struct {
		Title string
		Code  string
		Price float64
	}

	productModel := presetsBuilder.Model(&Product{})
	activityBuilder := activity.New(presetsBuilder, db)
	activityBuilder.RegisterModel(productModel).UseDefaultTab().AddKeys("Title").AddIgnoredFields("Code").SkipDelete()
	// @snippet_end
}

func ActivityRecordLogSample() {
	// @snippet_begin(ActivityRecordLogSample)
	presetsBuilder := presets.New()
	db, err := gorm.Open(sqlite.Open("/tmp/activity.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	type Product struct {
		Title string
		Code  string
		Price float64
	}

	activityBuilder := activity.New(presetsBuilder, db)
	activityBuilder.RegisterModel(&Product{})
	currentCtx := context.WithValue(context.Background(), activity.CreatorContextKey, "user1")
	activityBuilder.AddRecords("Publish", currentCtx, &Product{Title: "Product 1", Code: "P1", Price: 100}) // custmize the action name
	activityBuilder.AddRecords("Update Price", currentCtx, &Product{Title: "Product 1", Code: "P1", Price: 200})

	// @snippet_end
}
