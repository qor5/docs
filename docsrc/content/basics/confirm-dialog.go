package basics

import (
	"fmt"
	"github.com/qor5/docs/docsrc/examples/presets_examples"
	"path"
	"strings"

	"github.com/qor5/docs/docsrc/generated"
	"github.com/qor5/docs/docsrc/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var ConfirmDialog = Doc(
	Markdown(fmt.Sprintf("`%s`", strings.TrimRight(generated.OpenConfirmDialog, ","))+`
 is a pre-defined event used to show a confirm dialog for user to do confirm before executing the actual action.
`+
		`
### Queries
`+fmt.Sprintf("`%s`  ", strings.TrimRight(generated.ConfirmDialogConfirmEvent, ","))+
		`
required  
Usually the value will be *web.Plaid().EventFunc(the actual action event)....Go()*.  
  
`+fmt.Sprintf("`%s`  ", strings.TrimRight(generated.ConfirmDialogPromptText, ","))+
		`
optional  
To customize the prompt text.  
  
`+fmt.Sprintf("`%s`  ", strings.TrimRight(generated.ConfirmDialogDialogPortalName, ","))+
		`
optional  
To use a custom portal for dialog.  
`),
	Markdown(`
## Example
`),
	ch.Code(generated.ConfirmDialogSample).Language("go"),
	utils.Demo(
		"Confirm Dialog",
		path.Join(presets_examples.PresetsConfirmDialogPath, "/confirm-dialog"),
		"example_basics/confirm-dialog.go",
	),
).Slug("basics/confirm-dialog").Title("Confirm Dialog")
