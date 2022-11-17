package basics

import (
	"github.com/qor5/docs/docsrc/examples"
	"github.com/qor5/docs/docsrc/examples/e21_presents"
	"github.com/qor5/docs/docsrc/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var Brand = Doc(
	Markdown(`
Brand refers to the top area of the left menu bar, we provide two functions ~BrandTitle~ and ~BrandFunc~ to customize it.

## Simple customization
If you want only to change the brand string, you can use ~BrandTitle~ to set the string, the string will be displayed in the brand area with ~<H1>~ tag.
`),

	ch.Code(`b = presets.New()
b.URIPrefix("/admin").BrandTitle("Admin")`),

	Markdown(`
## Full customization
When you opt-in to full brand customization, you can use ~BrandFunc~ to be responsible for drawing for the entire brand area, such as you can put your own logo image in it.
`),

	ch.Code(examples.BrandSample).Language("go"),
	utils.Demo("Brand", e21_presents.PresetsBrandPath+"/pages", "e21_presents/brand.go"),
).Title("Brand").
	Slug("basics/brand")
