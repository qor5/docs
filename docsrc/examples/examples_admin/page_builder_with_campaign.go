package examples_admin

import (
	"github.com/qor5/admin/v3/pagebuilder"
	"github.com/qor5/admin/v3/presets"
	"github.com/qor5/admin/v3/publish"
	"github.com/qor5/ui/v3/vuetify"
	"github.com/qor5/web/v3"
	. "github.com/theplant/htmlgo"
	"gorm.io/gorm"
)

type MyContent struct {
	Text string
}

type Campaign struct {
	Title  string
	Banner string
}

type Product struct {
	Name string
}

func PageBuilderExample(b *presets.Builder, db *gorm.DB) {
	pb := pagebuilder.New("", ExampleDB(), b.I18n())

	header := pb.RegisterContainer("Header").Group("Navigation").
		RenderFunc(func(obj interface{}, input *pagebuilder.RenderInput, ctx *web.EventContext) HTMLComponent {
			c := obj.(*MyContent)
			return Div().Text(c.Text)
		})

	ed := header.Model(&MyContent{}).Editing("Text")
	ed.Field("Color").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) HTMLComponent {
		return vuetify.VTextField().
			Variant(vuetify.FieldVariantUnderlined).
			Label(field.Label).
			Attr(web.VField(field.FormKey, field.Value(obj))...)
	})

	// Campaigns Menu
	campaignModelBuilder := b.Model(&Campaign{})
	campaignModelBuilder.Listing("Title")

	ed1 := campaignModelBuilder.Editing("Title")

	detail := campaignModelBuilder.Detailing(
		publish.VersionsPublishBar,
		pagebuilder.PageBuilderPreviewCard,
		"CampaignDetail",
	)

	detail.Field("CampaignDetail").
		SetSwitchable(true).
		ShowComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) HTMLComponent {
			return Div()
		})

	pb.RegisterModelContainer("CampaignContent", ed1).RenderFunc(func(obj interface{}, input *pagebuilder.RenderInput, ctx *web.EventContext) HTMLComponent {
		c := obj.(*Campaign)
		return Div().Text(c.Banner)
	})

	campaign2ModelBuilder := b.Model(&Campaign{}).URIName("campaign2").InMenu(false)
	ed2 := campaign2ModelBuilder.Editing("Banner")
	pb.RegisterModelContainer("CampaignBanner", ed2).RenderFunc(func(obj interface{}, input *pagebuilder.RenderInput, ctx *web.EventContext) HTMLComponent {
		return Div()
	})

	campaignModelBuilder.Use(pb)

	// Products Menu
	productModelBuilder := b.Model(&Product{})
	productModelBuilder.Listing("Name")

	ed3 := campaignModelBuilder.Editing("Name")

	detail2 := campaignModelBuilder.Detailing(
		publish.VersionsPublishBar,
		pagebuilder.PageBuilderPreviewCard,
		"ProductDetail",
	)

	detail2.Field("ProductDetail").
		SetSwitchable(true).
		ShowComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) HTMLComponent {
			return Div()
		})

	pb.RegisterModelContainer("ProductDetail", ed3).RenderFunc(func(obj interface{}, input *pagebuilder.RenderInput, ctx *web.EventContext) HTMLComponent {
		c := obj.(*Campaign)
		return Div().Text(c.Banner)
	})

	productModelBuilder.Use(pb)
}
