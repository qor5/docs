package presets_examples

import (
	"github.com/qor5/admin/presets"
	"github.com/qor5/admin/presets/gorm2op"
	"github.com/qor5/ui/vuetify"
	"github.com/qor5/web"
	"github.com/theplant/htmlgo"
	"gorm.io/gorm"
)

type confirmDialog struct{}

func PresetsConfirmDialog(db *gorm.DB, b *presets.Builder) {
	_ = []interface{}{
		// @snippet_begin(OpenConfirmDialog)
		presets.OpenConfirmDialog,
		// @snippet_end
		// @snippet_begin(ConfirmDialogConfirmEvent)
		presets.ConfirmDialogConfirmEvent,
		// @snippet_end
		// @snippet_begin(ConfirmDialogPromptText)
		presets.ConfirmDialogPromptText,
		// @snippet_end
		// @snippet_begin(ConfirmDialogDialogPortalName)
		presets.ConfirmDialogDialogPortalName,
		// @snippet_end
	}

	b.URIPrefix(PresetsConfirmDialogPath).
		DataOperator(gorm2op.DataOperator(db))

	mb := b.Model(&confirmDialog{}).
		URIName("confirm-dialog").
		Label("Confirm Dialog")

	mb.Listing().PageFunc(func(ctx *web.EventContext) (r web.PageResponse, err error) {
		r.Body = htmlgo.Div(
			// @snippet_begin(ConfirmDialogSample)
			vuetify.VBtn("Delete File").
				Attr("@click",
					web.Plaid().
						EventFunc(presets.OpenConfirmDialog).
						Query(presets.ConfirmDialogConfirmEvent,
							`alert("file deleted")`,
						).
						Go(),
				),
			// @snippet_end
		).Class("ma-8")
		return r, nil
	})
}

const PresetsConfirmDialogPath = "/samples/confirm_dialog"
