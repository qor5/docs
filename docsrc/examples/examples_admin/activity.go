package examples_admin

import (
	"context"

	"github.com/qor5/admin/v3/activity"
	"github.com/qor5/admin/v3/presets"
	"github.com/qor5/admin/v3/presets/gorm2op"
)

func NewActivitySample() {
	// @snippet_begin(NewActivitySample)
	db := ExampleDB()
	presetsBuilder := presets.New().DataOperator(gorm2op.DataOperator(db))

	activityBuilder := activity.New(db)
	presetsBuilder.Use(activityBuilder)
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
