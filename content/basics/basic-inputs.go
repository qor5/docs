package basics

import (
	"github.com/qor5/docs/examples"
	"github.com/qor5/docs/examples/e11_vuetify_basic_inputs"
	"github.com/qor5/docs/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var BasicInputs = Doc(
	Markdown(`
Vuetify provides many form basic inputs, and also with error messages display on fields.

Here is one example:
`),
	ch.Code(examples.VuetifyBasicInputsSample).Language("go"),
	utils.Demo("Vuetify Basic Inputs", e11_vuetify_basic_inputs.VuetifyBasicInputsPath, "e11_vuetify_basic_inputs/page.go"),
).Title("Basic Inputs").
	Slug("vuetify-components/basic-inputs")
