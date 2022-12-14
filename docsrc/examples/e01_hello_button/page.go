package e01_hello_button

import (
	"github.com/qor5/web"
	. "github.com/theplant/htmlgo"
)

type mystate struct {
	Message string
}

func HelloButton(ctx *web.EventContext) (pr web.PageResponse, err error) {

	var s = &mystate{}
	if ctx.Flash != nil {
		s = ctx.Flash.(*mystate)
	}

	pr.Body = Div(
		Button("Hello").Attr("@click", web.POST().EventFunc("reload").Go()),
		Tag("input").
			Attr("type", "text").
			Attr("value", s.Message).
			Attr("@input", web.POST().
				EventFunc("reload").
				FieldValue("Message", web.Var("$event.target.value")).
				Go()),
		Div().
			Style("font-family: monospace;").
			Text(s.Message),
	)
	return
}

func reload(ctx *web.EventContext) (r web.EventResponse, err error) {
	var s = &mystate{}
	ctx.MustUnmarshalForm(s)
	ctx.Flash = s

	r.Reload = true
	return
}

var HelloButtonPB = web.Page(HelloButton).
	EventFunc("reload", reload)
