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
type WithPublishProduct struct {
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
	_ presets.SlugEncoder = (*WithPublishProduct)(nil)
	_ presets.SlugDecoder = (*WithPublishProduct)(nil)
)

func (p *WithPublishProduct) PrimarySlug() string {
	return fmt.Sprintf("%v_%v", p.ID, p.Version.Version)
}

func (p *WithPublishProduct) PrimaryColumnValuesBySlug(slug string) map[string]string {
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
	_ publish.PublishInterface   = (*WithPublishProduct)(nil)
	_ publish.UnPublishInterface = (*WithPublishProduct)(nil)
)

func (p *WithPublishProduct) GetPublishActions(db *gorm.DB, ctx context.Context, storage oss.StorageInterface) (objs []*publish.PublishAction, err error) {
	// create publish actions
	return
}

func (p *WithPublishProduct) GetUnPublishActions(db *gorm.DB, ctx context.Context, storage oss.StorageInterface) (objs []*publish.PublishAction, err error) {
	// create unpublish actions
	return
}

// @snippet_end

func PublishExample(b *presets.Builder, db *gorm.DB) {
	err := db.AutoMigrate(&WithPublishProduct{})
	if err != nil {
		panic(err)
	}

	b.URIPrefix(PublishExamplePath).
		DataOperator(gorm2op.DataOperator(db))

	// @snippet_begin(PublishConfigureView)
	mb := b.Model(&WithPublishProduct{})
	mb.RightDrawerWidth("1000")
	mb.Editing(publish.EditingFieldControlBar, "Name", "Price").
		Creating("Name", "Price")

	publisher := publish.New(db, nil)
	b.Plugins(publisher)
	mb.Plugins(publisher)
	// run the publisher job if Schedule is used
	go publish.RunPublisher(db, nil, publisher)
	// @snippet_end
}

const PublishExamplePath = "/samples/publish"
