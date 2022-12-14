package basics

import (
	"github.com/qor5/docs/docsrc/examples/e21_presents"
	"github.com/qor5/docs/docsrc/generated"
	"github.com/qor5/docs/docsrc/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var Permissions = Doc(
	Markdown(`## To list all the permissions in your project`),
	ch.Code(`perm.Verbose = true`).Language("go"),
	Markdown(`Then reboot your app, you can see all the permissions in the console`),

	Markdown(`
## Permissions sample:
`),
	ch.Code(generated.PresetsPermissionsSample).Language("go"),
	utils.Demo("Permissions Demo", e21_presents.PresetsPermissionsPath+"/customers", "e21_presents/permissions.go"),
).Title("Permissions").
	Slug("presets-guide/permissions")
