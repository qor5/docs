package basics

import (
	"github.com/qor5/docs/v3/docsrc/examples"
	"github.com/qor5/docs/v3/docsrc/examples/examples_presets"
	"github.com/qor5/docs/v3/docsrc/generated"
	"github.com/qor5/docs/v3/docsrc/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var NotificationCenter = Doc(
	Markdown(`
To enable notification center: Call ~NotificationFunc~ on ~presets.Builder~ With 2 function parameters
like this ~builder.NotificationFunc(NotifierComponent(), NotifierCount())~

The first function is for rendering the content of the popup after user clicked the "bell icon".
The second function is for rendering the number at the top right corner of the "bell icon".
`),

	ch.Code(generated.NotificationCenterSample).Language("go"),
	utils.DemoWithSnippetLocation("Notification Center", examples.URLPathByFunc(examples_presets.PresetsNotificationCenterSample)+"/pages", generated.NotificationCenterSampleLocation),
).Slug("basics/notification-center").Title("Notification Center")
