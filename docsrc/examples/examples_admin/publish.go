package examples_admin

import (
	"context"
	"fmt"
	"strings"

	"github.com/qor/oss"
	"github.com/qor5/admin/v3/presets"
	"github.com/qor5/admin/v3/presets/gorm2op"
	"github.com/qor5/admin/v3/publish"
	"gorm.io/gorm"
)

// @snippet_begin(PublishInjectModules)
type Product struct {
	gorm.Model

	Name  string
	Price int

	publish.Status
	publish.Schedule
	publish.Version
}

// @snippet_end

// @snippet_begin(PublishImplementSlugInterfaces)
var (
	_ presets.SlugEncoder = (*Product)(nil)
	_ presets.SlugDecoder = (*Product)(nil)
)

func (p *Product) PrimarySlug() string {
	return fmt.Sprintf("%v_%v", p.ID, p.Version.Version)
}

func (p *Product) PrimaryColumnValuesBySlug(slug string) map[string]string {
	segs := strings.Split(slug, "_")
	if len(segs) != 2 {
		panic("wrong slug")
	}

	return map[string]string{
		"id":      segs[0],
		"version": segs[1],
	}
}

// @snippet_end

// @snippet_begin(PublishImplementPublishInterfaces)
var (
	_ publish.PublishInterface   = (*Product)(nil)
	_ publish.UnPublishInterface = (*Product)(nil)
)

func (p *Product) GetPublishActions(db *gorm.DB, ctx context.Context, storage oss.StorageInterface) (objs []*publish.PublishAction, err error) {
	// create publish actions
	return
}

func (p *Product) GetUnPublishActions(db *gorm.DB, ctx context.Context, storage oss.StorageInterface) (objs []*publish.PublishAction, err error) {
	// create unpublish actions
	return
}

// @snippet_end

func PublishExample(b *presets.Builder, db *gorm.DB) {
	err := db.AutoMigrate(&Product{})
	if err != nil {
		panic(err)
	}

	b.URIPrefix(PublishExamplePath).
		DataOperator(gorm2op.DataOperator(db))

	// @snippet_begin(PublishConfigureView)
	mb := b.Model(&Product{})
	mb.Editing("StatusBar", "ScheduleBar", "Name", "Price")

	publisher := publish.New(db, nil).Models(mb)
	publisher.Install(b)
	// run the publisher job if Schedule is used
	go publish.RunPublisher(db, nil, publisher)
	// @snippet_end
}

const PublishExamplePath = "/samples/publish"
