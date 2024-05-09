package advanced_functions

import (
	"github.com/qor5/docs/v3/docsrc/examples/examples_presets"
	"github.com/qor5/docs/v3/docsrc/generated"
	"github.com/qor5/docs/v3/docsrc/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
	. "github.com/theplant/htmlgo"
)

var DetailPageForComplexObject = Doc(
	Markdown(`
By default, presets will only generate the listing page, editing page for a model,
It's for simple objects. But for a complicated object with a lots of relationships and connections,
and as the main data model of your system, It's better to have detail page for them. In there
You can add all kinds of operations conveniently.
`),
	ch.Code(generated.PresetsDetailPageTopNotesSample).Language("go"),
	utils.DemoWithSnippetLocation("Presets Detail Page Top Notes", examples_presets.PresetsDetailPageTopNotesPath+"/customers", generated.PresetsDetailPageTopNotesSampleLocation),
	Markdown(`
- The name of detailing fields are just a place holder for decide ordering
- ~CellComponentFunc~ customize how the cell display
- ~vx.DataTable~ create a data table, Which the Listing page uses the same component
- ~LoadMoreAt~ will only show for example 2 rows of data, and you can click load more to display all
- ~vx.Card~ display a card with toolbar you can setup action buttons
- We reference the new form drawer that ~b.Model(&Note{})~ creates, but hide notes in the menu
`),
	utils.Anchor(H2(""), "Details Info components and actions"),
	Markdown(`
A ~vx.DetailInfo~ component is used for display main detail field of the model.
And you can add any actions to the detail page with ease:
`),
	ch.Code(generated.PresetsDetailPageDetailsSample).Language("go"),
	utils.DemoWithSnippetLocation("Presets Detail Page Details", examples_presets.PresetsDetailPageDetailsPath+"/customers", generated.PresetsDetailPageDetailsSampleLocation),
	Markdown(`
- The ~vx.Card~ Actions links to two event functions: Agree Terms, and Update Details
- Agree Terms show a drawer popup that edit the ~term_agreed_at~ field
- Update Details reuse the edit customer form
`),

	utils.Anchor(H2(""), "More Usage for Data Table"),
	Markdown(`
A ~vx.DataTable~ component is very featured rich, Here check out the row expandable example:
`),
	ch.Code(generated.PresetsDetailPageCardsSample).Language("go"),
	utils.DemoWithSnippetLocation("Presets Detail Page Credit Cards", examples_presets.PresetsDetailPageCardsPath+"/customers", generated.PresetsDetailPageCardsSampleLocation),
	Markdown(`
- ~RowExpandFunc~ config the content when data table row expand
- ~cc.Editing~ setup the fields when edit
- ~cc.Creating~ setup the fields when create
`),
).Title("Detailing").
	Slug("presets-guide/detail-page-for-complex-object")
