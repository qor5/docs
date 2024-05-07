package advanced_functions

import (
	"github.com/qor5/docs/v3/docsrc/examples/examples_web"
	"github.com/qor5/docs/v3/docsrc/generated"
	"github.com/qor5/docs/v3/docsrc/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
	. "github.com/theplant/htmlgo"
)

var EventHandling = Doc(
	Markdown(`
We extend vue to support the following types of event handling, so you can simply use go code to implement some complex logic.

Using the ~~~Plaid()~~~ method will create an event handler that defaults to using the current ~~~vars~~~ and ~~~plaidForm~~~.
The default http request method is ~~~Post~~~, if you want to use the ~~~Get~~~ method, you can also use the ~~~Get()~~~ method directly to create an event handler
	`),

	utils.Anchor(H2(""), "URL"),
	Markdown(`Request a page.`),
	ch.Code(generated.EventHandlingURLSample).Language("go"),
	utils.Demo("Event Handling", examples_web.EventHandlingPagePath+"?api=url", "examples_web/event-handling.go#L14-L22"),

	utils.Anchor(H2(""), "PushState"),
	Markdown(`Reqest a page and also changing the window location.`),
	ch.Code(generated.EventHandlingPushStateSample).Language("go"),
	utils.Demo("Event Handling", examples_web.EventHandlingPagePath+"?api=pushstate", "examples_web/event-handling.go#27-L35"),

	utils.Anchor(H2(""), "Reload"),
	Markdown(`Refresh page.`),
	ch.Code(generated.EventHandlingReloadSample).Language("go"),
	utils.Demo("Event Handling", examples_web.EventHandlingPagePath+"?api=reload", "examples_web/event-handling.go#40-L49"),

	utils.Anchor(H2(""), "Query"),
	Markdown(`Request a page with a query.`),
	ch.Code(generated.EventHandlingQuerySample).Language("go"),
	utils.Demo("Event Handling", examples_web.EventHandlingPagePath+"?api=query", "examples_web/event-handling.go#L54-L62"),

	utils.Anchor(H2(""), "MergeQuery"),
	Markdown(`Request a page with merging a query.`),
	ch.Code(generated.EventHandlingMergeQuerySample).Language("go"),
	utils.Demo("Event Handling", examples_web.EventHandlingPagePath+"?api=merge_query", "examples_web/event-handling.go#L67-L75"),

	utils.Anchor(H2(""), "ClearMergeQuery"),
	Markdown(`Request a page with clearing a query.`),
	ch.Code(generated.EventHandlingClearMergeQuerySample).Language("go"),
	utils.Demo("Event Handling", examples_web.EventHandlingPagePath+"?api=clear_merge_query", "examples_web/event-handling.go#L80-L88"),

	utils.Anchor(H2(""), "StringQuery"),
	Markdown(`Request a page with a query string.`),
	ch.Code(generated.EventHandlingStringQuerySample).Language("go"),
	utils.Demo("Event Handling", examples_web.EventHandlingPagePath+"?api=string_query", "examples_web/event-handling.go#L93-L101"),

	utils.Anchor(H2(""), "Queries"),
	Markdown(`Request a page with url.Values.`),
	ch.Code(generated.EventHandlingQueriesSample).Language("go"),
	utils.Demo("Event Handling", examples_web.EventHandlingPagePath+"?api=queries", "examples_web/event-handling.go#L106-L114"),

	utils.Anchor(H2(""), "PushStateURL"),
	Markdown(`Request a page with a url and also changing the window location.`),
	ch.Code(generated.EventHandlingQueriesSample).Language("go"),
	utils.Demo("Event Handling", examples_web.EventHandlingPagePath+"?api=pushstateurl", "examples_web/event-handling.go#L119-L127"),

	utils.Anchor(H2(""), "Location"),
	Markdown(`Open a page with more options.`),
	ch.Code(generated.EventHandlingLocationSample).Language("go"),
	utils.Demo("Event Handling", examples_web.EventHandlingPagePath+"?api=location", "examples_web/event-handling.go#L132-L140"),

	utils.Anchor(H2(""), "FieldValue"),
	Markdown(`Fill in a value on form.`),
	ch.Code(generated.EventHandlingFieldValueSample).Language("go"),
	utils.Demo("Event Handling", examples_web.EventHandlingPagePath+"?api=fieldvalue", "examples_web/event-handling.go#L145-L153"),

	utils.Anchor(H2(""), "FormClear"),
	Markdown(`Clear all form data.`),
	ch.Code(generated.EventHandlingFieldValueSample).Language("go"),
	utils.Demo("Event Handling", examples_web.EventHandlingPagePath+"?api=formclear", "examples_web/event-handling.go#L165-L178"),

	utils.Anchor(H2(""), "EventFunc"),
	Markdown(`Register an event func and call it when the event is triggered.`),
	ch.Code(generated.EventHandlingEventFuncSample).Language("go"),
	utils.Demo("Event Handling", examples_web.EventHandlingPagePath+"?api=eventfunc", "examples_web/event-handling.go#L183-L191"),

	utils.Anchor(H2(""), "Script"),
	Markdown(`Run a script code.`),
	ch.Code(generated.EventHandlingBeforeScriptSample).Language("go"),
	utils.Demo("Event Handling", examples_web.EventHandlingPagePath+"?api=script", "examples_web/event-handling.go#L196-L204"),

	utils.Anchor(H2(""), "Raw"),
	Markdown(`Directly call the js method`),
	ch.Code(generated.EventHandlingRawSample).Language("go"),
	utils.Demo("Event Handling", examples_web.EventHandlingPagePath+"?api=raw", "examples_web/event-handling.go#L209-L217"),
).Title("Event Handling").Slug("basics/event-handling")
