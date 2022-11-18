package basics

import (
	"github.com/qor5/docs/docsrc/examples/e21_presents"
	"github.com/qor5/docs/docsrc/generated"
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

	ch.Code(generated.BrandTitleSample).Language("go"),
	utils.Demo("Brand Title", e21_presents.PresetsBrandTitlePath+"/pages", "e21_presents/brand.go"),

	Markdown(`
## Full customization
When you opt-in to full brand customization, you can use ~BrandFunc~ to be responsible for drawing for the entire brand area, such as you can put your own logo image in it.
`),

	ch.Code(generated.BrandFuncSample).Language("go"),
	utils.Demo("Brand Func", e21_presents.PresetsBrandFuncPath+"/pages", "e21_presents/brand.go"),

	Markdown(`
## Profile
Profile is below the brand area, where you can put the current user's information or others. We provide ~ProfileFunc~ to customize it.
`),

	ch.Code(generated.ProfileSample).Language("go"),
	utils.Demo("Profile", e21_presents.PresetsProfilePath+"/pages", "e21_presents/profile.go"),
).Title("Brand").
	Slug("basics/brand")
