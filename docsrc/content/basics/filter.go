package basics

import (
	"github.com/qor5/docs/docsrc/examples/example_basics"
	"github.com/qor5/docs/docsrc/generated"
	"github.com/qor5/docs/docsrc/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var Filter = Doc(
	Markdown(`

Assume we have a ~status~ filed in Post. It has 2 possible values, "draft" and "online". If we want to filter posts by its status. We can add a filter like this:

`),
	ch.Code(generated.FilterSample).Language("go"),
	utils.Demo("Basic filter", example_basics.PresetsBasicFilterPath+"/customers", "e21_presents/filter.go"),

	Markdown(`
### QOR5 now supports 8 types of filter option.

## Filter by String
Set the ~ItemType~ as ~vuetifyx.ItemTypeString~. No ~Options~ needed.
Under this mode, the filter would work in 2 ways,
1. the target value equal to the input string
2. the target value contains the input string

## Filter by Number
Set the ~ItemType~ as ~vuetifyx.ItemTypeNumber~. No ~Options~ needed.
Under this mode, the filter would work in 4 ways
1. the target value equal to the input number
2. the target value is between the input numbers
3. the target value is greater than the input number
4. the target value is less than the input number

TODO:
	3. SelectItem
	4. DateItem
	5. DateRangeItem
	6. DatetimeRangeItem
	7. MultipleSelectItem
	8. LinkageSelectItem
`),
).Title("Filters").
	Slug("basics/filter")
