package utils

import (
	"cmp"
	"encoding/json"
	"fmt"
	"os"

	"github.com/qor5/web/v3"
	"github.com/shurcooL/sanitized_anchor_name"
	"github.com/sunfmin/snippetgo/parse"
	. "github.com/theplant/htmlgo"
)

func Anchor(h *HTMLTagBuilder, text string) HTMLComponent {
	anchorName := sanitized_anchor_name.Create(text)
	return h.Children(
		Text(text),
		A().Class("anchor").Href(fmt.Sprintf("#%s", anchorName)),
	).Id(anchorName)
}

type Example struct {
	Title      string
	DemoPath   string
	SourcePath string
}

var LiveExamples []*Example

var envGitBranch string

func init() {
	envGitBranch = cmp.Or(os.Getenv("GIT_BRANCH"), "main")
}

func DemoWithSnippetLocation(title string, demoPath string, location parse.Location) HTMLComponent {
	return Demo(title, demoPath, fmt.Sprintf("%s#L%d-L%d", location.File, location.StartLine, location.EndLine))
}

func Demo(title string, demoPath string, sourcePath string) HTMLComponent {
	if sourcePath != "" {
		sourcePath = fmt.Sprintf("https://github.com/qor5/docs/tree/%s/docsrc/%s", envGitBranch, sourcePath)
	}
	ex := &Example{
		Title:      title,
		DemoPath:   demoPath,
		SourcePath: sourcePath,
	}

	if title != "" {
		LiveExamples = append(LiveExamples, ex)
	}

	return Div(
		Div(
			A().Text("Check the demo").Href(ex.DemoPath).Target("_blank"),
			Iff(ex.SourcePath != "", func() HTMLComponent {
				return Components(
					Text(" | "),
					A().Text("Source on GitHub").
						Href(ex.SourcePath).
						Target("_blank"),
				)
			}),
		).Class("demo"),
	)
}

func ExamplesDoc() HTMLComponent {
	u := Ul()
	for _, le := range LiveExamples {
		u.AppendChildren(
			Li(
				A().Href(le.DemoPath).Text(le.Title).Target("_blank"),
				Text(" | "),
				A().Href(le.SourcePath).Text("Source").Target("_blank"),
			),
		)
	}
	return u
}

func PrettyFormAsJSON(ctx *web.EventContext) HTMLComponent {
	if ctx.R.MultipartForm == nil {
		return nil
	}

	formData, err := json.MarshalIndent(ctx.R.MultipartForm, "", "\t")
	if err != nil {
		panic(err)
	}

	return Pre(
		string(formData),
	)
}
