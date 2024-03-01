package basics

import (
	"github.com/qor5/docs/docsrc/examples/vuetify_examples"
	"github.com/qor5/docs/docsrc/generated"
	"github.com/qor5/docs/docsrc/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var BasicInputs = Doc(
	Markdown(`
Vuetify provides many form basic inputs, and also with error messages display on fields.

Here is one example:
`),
	ch.Code(generated.VuetifyBasicInputsSample).Language("go"),
	utils.Demo("Vuetify Basic Inputs", vuetify_examples.VuetifyBasicInputsPath, "e11_vuetify_basic_inputs/page.go"),
).Title("Basic Inputs").
	Slug("vuetify-components/basic-inputs")
