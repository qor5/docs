package basics

import (
	"github.com/qor5/docs/docsrc/examples/vuetify_examples"
	"github.com/qor5/docs/docsrc/generated"
	"github.com/qor5/docs/docsrc/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var VariantSubForm = Doc(
	Markdown(`
VSelect changes, the form below it will change to a new form accordingly.

By use of ~web.Portal()~ and ~VSelect~'s ~OnInput~
`),
	ch.Code(generated.VuetifyVariantSubForm).Language("go"),
	utils.Demo("Vuetify Variant Sub Form", vuetify_examples.VuetifyVariantSubFormPath, "e22_vuetify_variant_sub_form/page.go"),
).Title("Variant Sub Form").
	Slug("vuetify-components/variant-sub-form")
