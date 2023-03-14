package basics

import (
	"github.com/qor5/docs/docsrc/generated"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var Activity = Doc(
	Markdown(`
As an admin user of CMS, You may need to record some important operations and these records should be easily queried and audited,
so QOR5 provides a built-in activity module to record the operation of the model and support the following features:

* Support recording very detailed changes when a model data is changed.
* Support adding ignore fields when diffing a modified data, such as the update time.
* Support customizing how to diff a complex field type, such as the time.Time type.
* Support customizing the keys to indentify the current model data.
* Support both automatic recording and manual recording the CRUD operation.
* Support customizing the action about what the current operation is.
* Support querying the activity log from admin page.

## Initialize the activity
	`),
	ch.Code(generated.NewActivitySample).Language("go"),
	Markdown(`
- The above code will create a new activity instance with the default configuration.
- If you already have a key to fetch the current user from the context, you can use ~~~SetCreatorContextKey~~~ method to set it.
- If you want to fetch the DB from the context, you can use ~~~SetDBContextKey~~~ method to set it.
- If you want to customize the text that displayed on the model edit page, you can use ~~~SetTabHeading~~~ method to customized it.

## Register the preset models
`),
	ch.Code(generated.ActivityRegisterPresetsModelsSample).Language("go"),
	Markdown(`
- The above code will register the product model into activity and the product model will be recorded when it is created, updated or deleted automatically.
- By default, the activity will use the ~~~primary~~~ field as the key to indentify the current model data. You also can use ~~~SetKeys~~~ and ~~~AddKeys~~~ methods to customize the keys.
- By default, the activity will ignore the ~~~ID~~~, ~~~CreatedAt~~~, ~~~UpdatedAt~~~, ~~~DeletedAt~~~ fields when diffing a modified data. You also can use ~~~AddIgnoredFields~~~ and ~~~SetIgnoredFields~~~ methods to customize the ignore fields.
- The activity already handle some special field types, such as the time.Time and media_library.MediaBox, you also can use ~~~AddTypeHanders~~~ methods to add more type handles.
- If you want to skip the automatic recording, you can use ~~~SkipCreate~~~, ~~~SkipUpdate~~~ and ~~~SkipDelete~~~ methods to skip the automatic recording.
- If you want to display the related activity log on the model edit page, you can use ~~~UseDefaultTab~~~ method to enable it.
- If you want to display the link of which the model page is changed, you can use ~~~SetLink~~~ method to set it.

## Record the activity log manually
If you register a preset model into the activity, the activity will record the activity log automatically for the CRUD operation of the model. But if you want to record the activity log manually for some other operations or you want to register a model that is not a preset model, you can use the following sample to record the activity log manually.
	`),
	ch.Code(generated.ActivityRecordLogSample).Language("go"),
).Title("Activity Log")
