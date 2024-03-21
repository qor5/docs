package vuetify_examples

// @snippet_begin(VuetifySnackBarsSample)
import (
	. "github.com/qor5/ui/vuetify"
	"github.com/qor5/web"
)

func VuetifyCheckbox(ctx *web.EventContext) (pr web.PageResponse, err error) {
	//VCheckbox().
	//		FieldName(field.FormKey).
	//		Lable(field.Label).
	//		Value(reflectutils.MustGet(obj, field.Name).(bool)).
	//		ErrorMessages(field.Errors...).
	//		Disabled(field.Disabled)
	pr.Body = VContainer(
		VCheckbox().Attr("label", "Checkbox"),
	)
	return
}

var VuetifyCheckboxPB = web.Page(VuetifyCheckbox)

const VuetifyCheckboxPath = "/samples/vuetify-Checkbox"

// @snippet_end
