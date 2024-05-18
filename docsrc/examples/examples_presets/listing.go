// @snippet_begin(PresetHelloWorldSample)
package examples_presets

import (
	"fmt"
	"net/url"
	"time"

	"github.com/qor5/admin/v3/presets"
	"github.com/qor5/admin/v3/presets/actions"
	"github.com/qor5/admin/v3/presets/gorm2op"
	v "github.com/qor5/ui/v3/vuetify"
	"github.com/qor5/ui/v3/vuetifyx"
	"github.com/qor5/web/v3"
	"github.com/qor5/x/v3/i18n"
	h "github.com/theplant/htmlgo"
	"github.com/theplant/osenv"
	"golang.org/x/text/language"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Customer struct {
	ID              int
	Name            string
	Email           string
	Description     string
	CompanyID       int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	ApprovedAt      *time.Time
	TermAgreedAt    *time.Time
	ApprovalComment string
}

type Address struct {
	ID       int
	Province string
	City     string
	District string
}

var DB *gorm.DB

func init() {
	DB = setupDB()
}

var dbParamsString = osenv.Get("DB_PARAMS", "database connection string", "")

func setupDB() (db *gorm.DB) {
	var err error
	db, err = gorm.Open(postgres.Open(dbParamsString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Logger.LogMode(logger.Info)
	err = db.AutoMigrate(
		&Customer{},
		&Company{},
		&Address{},
	)
	if err != nil {
		panic(err)
	}
	return
}

func PresetsHelloWorld(b *presets.Builder) (m *presets.ModelBuilder, db *gorm.DB) {
	db = DB

	b.URIPrefix(PresetsHelloWorldPath).
		DataOperator(gorm2op.DataOperator(db))
	m = b.Model(&Customer{})

	return
}

const PresetsHelloWorldPath = "/samples/presets_hello_world"

// @snippet_end

// @snippet_begin(PresetsListingCustomizationFieldsSample)

type Company struct {
	ID   int
	Name string
}

func PresetsListingCustomizationFields(b *presets.Builder) (
	cust *presets.ModelBuilder,
	cl *presets.ListingBuilder,
	ce *presets.EditingBuilder,
	db *gorm.DB,
) {
	b.I18n().
		SupportLanguages(language.English, language.SimplifiedChinese).
		RegisterForModule(language.SimplifiedChinese, presets.ModelsI18nModuleKey, Messages_zh_CN)

	cust, db = PresetsHelloWorld(b)
	b.URIPrefix(PresetsListingCustomizationFieldsPath)

	cl = cust.Listing("ID", "Name", "Company", "Email").
		SearchColumns("name", "email").SelectableColumns(true)
	cl.Field("Company").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) h.HTMLComponent {
		c := obj.(*Customer)
		var comp Company
		if c.CompanyID == 0 {
			return h.Td()
		}

		db.First(&comp, "id = ?", c.CompanyID)
		return h.Td(
			h.A().Text(comp.Name).
				Class("text-decoration-none", "text-blue").
				Href("javascript:void(0)").
				Attr("@click",
					web.POST().EventFunc(actions.Edit).
						Query(presets.ParamID, fmt.Sprint(comp.ID)).
						URL(PresetsListingCustomizationFieldsPath+"/companies").
						Go()),
			h.Text("-"),
			h.A().Text("(Open in Dialog)").
				Class("text-decoration-none", "text-blue").
				Href("javascript:void(0)").
				Attr("@click",
					web.POST().EventFunc(actions.Edit).
						Query(presets.ParamID, fmt.Sprint(comp.ID)).
						Query(presets.ParamOverlay, actions.Dialog).
						URL(PresetsListingCustomizationFieldsPath+"/companies").
						Go(),
				),
		)
	})

	ce = cust.Editing("Name", "CompanyID")

	cust.RegisterEventFunc("updateCompanyList", func(ctx *web.EventContext) (r web.EventResponse, err error) {
		companyID := ctx.ParamAsInt(presets.ParamOverlayUpdateID)
		r.UpdatePortals = append(r.UpdatePortals, &web.PortalUpdate{
			Name: "companyListPortal",
			Body: companyList(ctx, db, companyID),
		})
		return
	})

	ce.Field("CompanyID").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) h.HTMLComponent {
		c := obj.(*Customer)
		return web.Portal(companyList(ctx, db, c.CompanyID)).Name("companyListPortal")
	})

	comp := b.Model(&Company{})
	comp.Editing().ValidateFunc(func(obj interface{}, ctx *web.EventContext) (err web.ValidationErrors) {
		c := obj.(*Company)
		if len(c.Name) < 5 {
			err.GlobalError("name must longer than 5")
		}
		return
	})

	return
}

func companyList(ctx *web.EventContext, db *gorm.DB, companyID int) h.HTMLComponent {
	msgr := i18n.MustGetModuleMessages(ctx.R, presets.ModelsI18nModuleKey, Messages_en_US).(*Messages)
	var comps []Company
	db.Find(&comps)
	return h.Div(
		v.VSelect().
			Label(msgr.CustomersCompanyID).
			Variant("underlined").
			Items(comps).
			Attr(web.VField("CompanyID", companyID)...).
			ItemTitle("Name").ItemValue("ID"),
		h.A().Text("Add Company").
			Class("text-decoration-none", "text-blue").
			Href("javascript:void(0)").Attr("@click",
			web.POST().
				URL(PresetsListingCustomizationFieldsPath+"/companies").
				EventFunc(actions.New).
				Query(presets.ParamOverlay, actions.Dialog).
				Query(presets.ParamOverlayAfterUpdateScript,
					web.POST().EventFunc("updateCompanyList").Go()).
				Go(),
		),
	)
}

const PresetsListingCustomizationFieldsPath = "/samples/presets_listing_customization_fields"

// @snippet_end

// @snippet_begin(PresetsListingCustomizationFiltersSample)

func PresetsListingCustomizationFilters(b *presets.Builder) (
	cust *presets.ModelBuilder,
	cl *presets.ListingBuilder,
	ce *presets.EditingBuilder,
	db *gorm.DB,
) {
	cust, cl, ce, db = PresetsListingCustomizationFields(b)
	b.URIPrefix(PresetsListingCustomizationFiltersPath)

	cl.FilterDataFunc(func(ctx *web.EventContext) vuetifyx.FilterData {
		msgr := i18n.MustGetModuleMessages(ctx.R, presets.ModelsI18nModuleKey, Messages_en_US).(*Messages)
		var companyOptions []*vuetifyx.SelectItem
		err := db.Model(&Company{}).Select("name as text, id as value").Scan(&companyOptions).Error
		if err != nil {
			panic(err)
		}

		return []*vuetifyx.FilterItem{
			{
				Key:      "created",
				Label:    msgr.CustomersFilterCreated,
				ItemType: vuetifyx.ItemTypeDatetimeRange,
				// SQLCondition: `cast(strftime('%%s', created_at) as INTEGER) %s ?`,
				SQLCondition: `created_at %s ?`,
			},
			{
				Key:      "approved",
				Label:    msgr.CustomersFilterApproved,
				ItemType: vuetifyx.ItemTypeDatetimeRange,
				// SQLCondition: `cast(strftime('%%s', created_at) as INTEGER) %s ?`,
				SQLCondition: `created_at %s ?`,
			},
			{
				Key:          "name",
				Label:        msgr.CustomersFilterName,
				ItemType:     vuetifyx.ItemTypeString,
				SQLCondition: `name %s ?`,
			},
			{
				Key:          "company",
				Label:        msgr.CustomersFilterCompany,
				ItemType:     vuetifyx.ItemTypeSelect,
				SQLCondition: `company_id %s ?`,
				Options:      companyOptions,
			},
		}
	})
	return
}

const PresetsListingCustomizationFiltersPath = "/samples/presets-listing-customization-filters"

// @snippet_end

// @snippet_begin(PresetsListingCustomizationTabsSample)

func PresetsListingCustomizationTabs(b *presets.Builder) (
	cust *presets.ModelBuilder,
	cl *presets.ListingBuilder,
	ce *presets.EditingBuilder,
	db *gorm.DB,
) {
	cust, cl, ce, db = PresetsListingCustomizationFilters(b)
	b.URIPrefix(PresetsListingCustomizationTabsPath)

	cl.FilterTabsFunc(func(ctx *web.EventContext) []*presets.FilterTab {
		var c Company
		db.First(&c)
		return []*presets.FilterTab{
			{
				Label: "Felix",
				Query: url.Values{"name.ilike": []string{"felix"}},
			},
			{
				Label: "The Plant",
				Query: url.Values{"company": []string{fmt.Sprint(c.ID)}},
			},
			{
				Label: "Approved",
				Query: url.Values{"approved.gt": []string{fmt.Sprint(1)}},
			},
			{
				Label: "All",
				Query: url.Values{"all": []string{"1"}},
			},
		}
	})
	return
}

const PresetsListingCustomizationTabsPath = "/samples/presets-listing-customization-tabs"

// @snippet_end

// @snippet_begin(PresetsListingCustomizationBulkActionsSample)

func PresetsListingCustomizationBulkActions(b *presets.Builder) (
	cust *presets.ModelBuilder,
	cl *presets.ListingBuilder,
	ce *presets.EditingBuilder,
	db *gorm.DB,
) {
	cust, cl, ce, db = PresetsListingCustomizationTabs(b)
	b.URIPrefix(PresetsListingCustomizationBulkActionsPath)

	cl.BulkAction("Approve").Label("Approve").
		UpdateFunc(func(selectedIds []string, ctx *web.EventContext) (err error) {
			comment := ctx.R.FormValue("ApprovalComment")
			if len(comment) < 10 {
				ctx.Flash = "comment should larger than 10"
				return
			}
			err = db.Model(&Customer{}).
				Where("id IN (?)", selectedIds).
				Updates(map[string]interface{}{"approved_at": time.Now(), "approval_comment": comment}).Error
			if err != nil {
				ctx.Flash = err.Error()
			}
			return
		}).
		ComponentFunc(func(selectedIds []string, ctx *web.EventContext) h.HTMLComponent {
			comment := ctx.R.FormValue("ApprovalComment")
			errorMessage := ""
			if ctx.Flash != nil {
				errorMessage = ctx.Flash.(string)
			}
			return v.VTextField().
				Variant("underlined").
				Attr(web.VField("ApprovalComment", comment)...).
				Label("Comment").
				ErrorMessages(errorMessage)
		})

	cl.BulkAction("Delete").Label("Delete").
		UpdateFunc(func(selectedIds []string, ctx *web.EventContext) (err error) {
			err = db.Where("id IN (?)", selectedIds).Delete(&Customer{}).Error
			return
		}).
		ComponentFunc(func(selectedIds []string, ctx *web.EventContext) h.HTMLComponent {
			return h.Div().Text(fmt.Sprintf("Are you sure you want to delete %s ?", selectedIds)).Class("title deep-orange--text")
		})

	return
}

const PresetsListingCustomizationBulkActionsPath = "/samples/presets-listing-customization-bulk-actions"

// @snippet_end

// @snippet_begin(PresetsListingCustomizationSearcherSample)

func PresetsListingCustomizationSearcher(b *presets.Builder) {
	db := setupDB()
	b.URIPrefix(PresetsListingCustomizationSearcherPath).DataOperator(gorm2op.DataOperator(db))
	mb := b.Model(&Customer{})
	mb.Listing().SearchFunc(func(model interface{}, params *presets.SearchParams, ctx *web.EventContext) (r interface{}, totalCount int, err error) {
		// only display approved customers
		qdb := db.Where("approved_at IS NOT NULL")
		return gorm2op.DataOperator(qdb).Search(model, params, ctx)
	})
}

// @snippet_end

const PresetsListingCustomizationSearcherPath = "/samples/presets_listing_customization_searcher"
