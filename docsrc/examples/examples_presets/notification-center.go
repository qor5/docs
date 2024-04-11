package examples_presets

// @snippet_begin(NotificationCenterSample)
import (
	"github.com/qor5/admin/presets"
	"github.com/qor5/admin/presets/gorm2op"
	v "github.com/qor5/ui/vuetify"
	"github.com/qor5/web"
	h "github.com/theplant/htmlgo"
)

func PresetsNotificationCenterSample(b *presets.Builder) {
	db := InitDB()
	b.URIPrefix(NotificationCenterSamplePath).
		DataOperator(gorm2op.DataOperator(db))

	db.AutoMigrate(&Page{})
	b.Model(&Page{})

	b.NotificationFunc(NotifierComponent(), NotifierCount())

	return
}

func NotifierComponent() func(ctx *web.EventContext) h.HTMLComponent {
	return func(ctx *web.EventContext) h.HTMLComponent {
		return v.VList(
			v.VListItem(
				v.VListItemTitle(
					h.A(h.Label("New Notice:"),
						h.Text("unread notes: 3")),
				),
			))
	}
}

func NotifierCount() func(ctx *web.EventContext) int {
	return func(ctx *web.EventContext) int {
		// Use your own count calculation logic here
		return 3
	}
}

// @snippet_end
const NotificationCenterSamplePath = "/samples/notification_center"
