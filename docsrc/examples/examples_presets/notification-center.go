package examples_presets

// @snippet_begin(NotificationCenterSample)
import (
	"github.com/qor5/admin/v3/presets"
	"github.com/qor5/admin/v3/presets/gorm2op"
	v "github.com/qor5/ui/v3/vuetify"
	"github.com/qor5/web/v3"
	h "github.com/theplant/htmlgo"
	"gorm.io/gorm"
)

func PresetsNotificationCenterSample(b *presets.Builder, db *gorm.DB) {
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
