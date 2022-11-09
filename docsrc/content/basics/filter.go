package basics

import (
	"github.com/qor5/docs/docsrc/examples"
	"github.com/qor5/docs/docsrc/examples/e21_presents"
	"github.com/qor5/docs/docsrc/utils"
	"github.com/theplant/docgo/ch"

	. "github.com/theplant/docgo"
)

var Filter = Doc(
	Markdown(`

To add a basic filter to the list page

For example:
`),
	ch.Code(examples.FilterSample).Language("go"),
	utils.Demo("Basic filter", e21_presents.PresetsBasicFilterPath+"/customers", "e21_presents/filter.go"),
	Markdown(`
	Call ~FilterDataFunc~ on a ~ListingBuilder~
`),
).Title("Filters").
	Slug("basics/filter")
