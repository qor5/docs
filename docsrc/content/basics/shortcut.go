package basics

import (
	"github.com/qor5/docs/docsrc/examples/web_examples"
	"github.com/qor5/docs/docsrc/generated"
	"github.com/qor5/docs/docsrc/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var ShortCut = Doc(
	Markdown(`
To add keyboard shortcut to a button:

Trigger the event by [GlobalEvents](https://www.npmjs.com/package/vue-global-events).
You can configure your own keyboard event like ~@keyup.ctrl.enter~ to trigger the event.

Also you can setup the ~filter~ function to limit when this event can be triggered by shortcut.
In the example, the event would only be triggered when ~locals.shortCutEnabled~ is opened.
`),

	ch.Code(generated.ShortCutSample).Language("go"),
	utils.Demo("Shortcut", web_examples.ShortCutSamplePath, "e00_basics/shortcut.go"),
).Slug("basics/shortcut").Title("Keyboard Shortcut")
