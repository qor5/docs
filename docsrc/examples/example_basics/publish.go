package example_basics

import (
	"context"
	"fmt"
	"strings"

	"github.com/qor/oss"
	"github.com/qor5/admin/presets"
	"github.com/qor5/admin/presets/gorm2op"
	"github.com/qor5/admin/publish"
	publish_view "github.com/qor5/admin/publish/views"
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
var _ presets.SlugEncoder = (*Product)(nil)
var _ presets.SlugDecoder = (*Product)(nil)

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
var _ publish.PublishInterface = (*Product)(nil)
var _ publish.UnPublishInterface = (*Product)(nil)

func (p *Product) GetPublishActions(db *gorm.DB, ctx context.Context, storage oss.StorageInterface) (objs []*publish.PublishAction, err error) {
	// create publish actions
	return
}

func (p *Product) GetUnPublishActions(db *gorm.DB, ctx context.Context, storage oss.StorageInterface) (objs []*publish.PublishAction, err error) {
	// create unpublish actions
	return
}

// @snippet_end

func PublishExample(b *presets.Builder) {
	b.URIPrefix(PublishExamplePath).
		DataOperator(gorm2op.DataOperator(DB))

		// @snippet_begin(PublishConfigureView)
	mb := b.Model(&Product{})
	mb.Editing("Status", "Schedule", "Name", "Price")

	publisher := publish.New(DB, nil)
	publish_view.Configure(b, DB, nil, publisher, mb)

	// run the publisher job if Schedule is used
	go publish.RunPublisher(DB, nil, publisher)
	// @snippet_end
}

const PublishExamplePath = "/samples/publish"
