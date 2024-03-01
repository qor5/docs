package admin_examples

// @snippet_begin(L10nFullExample)
import (
	"fmt"
	"net/http"
	"strings"

	"github.com/qor5/admin/l10n"
	l10n_view "github.com/qor5/admin/l10n/views"
	"github.com/qor5/admin/presets"
	"github.com/qor5/admin/presets/gorm2op"
	"gorm.io/gorm"
)

// @snippet_begin(L10nModelExample)
type L10nModel struct {
	gorm.Model
	Title string

	l10n.Locale
}

func (lm *L10nModel) PrimarySlug() string {
	return fmt.Sprintf("%v_%v", lm.ID, lm.LocaleCode)
}

func (lm *L10nModel) PrimaryColumnValuesBySlug(slug string) map[string]string {
	segs := strings.Split(slug, "_")
	if len(segs) != 2 {
		panic("wrong slug")
	}

	return map[string]string{
		"id":          segs[0],
		"locale_code": segs[1],
	}
}

// @snippet_end

func LocalizationExampleMock(b *presets.Builder) {
	if err := DB.AutoMigrate(&L10nModel{}); err != nil {
		panic(err)
	}

	b.URIPrefix(LocalizationExamplePath).
		DataOperator(gorm2op.DataOperator(DB))

	// @snippet_begin(L10nBuilderExample)
	l10nBuilder := l10n.New()
	l10nBuilder.
		RegisterLocales("International", "international", "International").
		RegisterLocales("China", "cn", "China").
		RegisterLocales("Japan", "jp", "Japan").
		GetSupportLocaleCodesFromRequestFunc(func(R *http.Request) []string {
			return l10nBuilder.GetSupportLocaleCodes()[:]
		})
	// @snippet_end

	// @snippet_begin(L10nConfigureExample)
	mb := b.Model(&L10nModel{}).URIName("l10n-models")
	l10n_view.Configure(b, DB, l10nBuilder, nil, mb)
	mb.Listing("ID", "Title", "Locale")
	// @snippet_end
	// @snippet_end

}

const LocalizationExamplePath = "/samples/l10n"
