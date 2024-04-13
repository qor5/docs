package examples_presets

import (
	"fmt"
	"io"
	"net/http"

	"github.com/qor5/admin/v3/presets"
	"github.com/qor5/admin/v3/presets/gorm2op"
	"github.com/qor5/ui/v3/tiptap"
	v "github.com/qor5/ui/v3/vuetify"
	"github.com/qor5/web/v3"
	"github.com/sunfmin/reflectutils"
	h "github.com/theplant/htmlgo"
	"gorm.io/gorm"
)

// @snippet_begin(PresetsEditingCustomizationDescriptionSample)

func PresetsEditingCustomizationDescription(b *presets.Builder) (
	cust *presets.ModelBuilder,
	cl *presets.ListingBuilder,
	ce *presets.EditingBuilder,
	db *gorm.DB,
) {
	cust, cl, ce, db = PresetsListingCustomizationBulkActions(b)
	b.URIPrefix(PresetsEditingCustomizationDescriptionPath)
	b.ExtraAsset("/tiptap.js", "text/javascript", tiptap.JSComponentsPack())
	b.ExtraAsset("/tiptap.css", "text/css", tiptap.CSSComponentsPack())

	ce.Only("Name", "CompanyID", "Description")

	ce.Field("Description").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) h.HTMLComponent {
		return tiptap.TipTapEditor().
			Attr(web.VField(field.Name, field.Value(obj).(string))...)
	})
	return
}

const PresetsEditingCustomizationDescriptionPath = "/samples/presets-editing-customization-description"

// @snippet_end

// @snippet_begin(PresetsEditingCustomizationFileTypeSample)

type MyFile string

type Product struct {
	ID        int
	Title     string
	MainImage MyFile
}

func PresetsEditingCustomizationFileType(b *presets.Builder) (
	cust *presets.ModelBuilder,
	cl *presets.ListingBuilder,
	ce *presets.EditingBuilder,
	db *gorm.DB,
) {
	cust, cl, ce, db = PresetsEditingCustomizationDescription(b)
	err := db.AutoMigrate(&Product{})
	if err != nil {
		panic(err)
	}

	b.URIPrefix(PresetsEditingCustomizationFileTypePath)
	b.FieldDefaults(presets.WRITE).
		FieldType(MyFile("")).
		ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) h.HTMLComponent {
			val := field.Value(obj).(MyFile)
			var img h.HTMLComponent
			if len(string(val)) > 0 {
				img = h.Img(string(val))
			}
			var er h.HTMLComponent
			if len(field.Errors) > 0 {
				er = h.Div().Text(field.Errors[0]).Style("color:red")
			}
			return h.Div(
				img,
				er,
				h.Input("").Type("file").Attr("@change", fmt.Sprintf("form.%s_NewFile = $event.target.files[0]", field.Name)),
			)
		}).
		SetterFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) (err error) {
			ff, _, _ := ctx.R.FormFile(fmt.Sprintf("%s_NewFile", field.Name))

			if ff == nil {
				return
			}
			var req *http.Request
			req, err = http.NewRequest("PUT", "https://transfer.sh/myfile.png", ff)
			if err != nil {
				return
			}
			var res *http.Response
			res, err = http.DefaultClient.Do(req)
			if err != nil {
				panic(err)
			}
			var b []byte
			b, err = io.ReadAll(res.Body)
			if err != nil {
				return
			}
			if res.StatusCode == 500 {
				err = fmt.Errorf("%s", string(b))
				return
			}
			err = reflectutils.Set(obj, field.Name, MyFile(b))
			return
		})

	mb := b.Model(&Product{})
	mb.Editing("Title", "MainImage")
	return
}

const PresetsEditingCustomizationFileTypePath = "/samples/presets-editing-customization-file-type"

// @snippet_end

// @snippet_begin(PresetsEditingCustomizationValidationSample)

func PresetsEditingCustomizationValidation(b *presets.Builder) (
	cust *presets.ModelBuilder,
	cl *presets.ListingBuilder,
	ce *presets.EditingBuilder,
	db *gorm.DB,
) {
	cust, cl, ce, db = PresetsEditingCustomizationDescription(b)
	b.URIPrefix(PresetsEditingCustomizationValidationPath)

	ce.ValidateFunc(func(obj interface{}, ctx *web.EventContext) (err web.ValidationErrors) {
		cus := obj.(*Customer)
		if len(cus.Name) < 10 {
			err.FieldError("Name", "name is too short")
		}
		return
	})
	return
}

const PresetsEditingCustomizationValidationPath = "/samples/presets-editing-customization-validation"

// @snippet_end

// @snippet_begin(PresetsEditingCustomizationTabsSample)

func PresetsEditingCustomizationTabs(b *presets.Builder) {
	db := setupDB()
	b.URIPrefix(PresetsEditingCustomizationTabsPath).DataOperator(gorm2op.DataOperator(db))
	mb := b.Model(&Company{})
	mb.Listing("ID", "Name")
	mb.Editing().AppendTabsPanelFunc(func(obj interface{}, ctx *web.EventContext) (tab, content h.HTMLComponent) {
		c := obj.(*Company)
		tab = v.VTab(h.Text("New Tab")).Value("2")
		content = v.VWindowItem(
			v.VListItemTitle(h.Text(fmt.Sprintf("Name: %s", c.Name))),
		).Value("2").Class("pa-4")
		return
	})
}

// @snippet_end

const PresetsEditingCustomizationTabsPath = "/samples/presets_editing_customization_tabs"
