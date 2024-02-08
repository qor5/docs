package e00_basics

import (
	"github.com/qor5/docs/docsrc/utils"
	. "github.com/qor5/ui/vuetify"
	"github.com/qor5/web"
	. "github.com/theplant/htmlgo"
)

// @snippet_begin(WebScopeUseLocalsSample1)
func UseLocals(ctx *web.EventContext) (pr web.PageResponse, err error) {
	pr.Body = VCard(
		VBtn("Test Can Not Change Other Scope").Attr("@click", `locals.btnLabel = "YES"`),
		web.Scope(
			VCard(
				VBtn("").
					Attr("v-text", "locals.btnLabel").
					Attr("@click", `
if (locals.btnLabel == "Add") {
	locals.items.push({text: "B", icon: "done"});
	locals.btnLabel = "Remove";
} else {
	locals.items.pop();
	locals.btnLabel = "Add";
}`),

				VList(
					VListSubheader(
						Text("REPORTS"),
					),
					VListGroup(
						VListItem().Attr("v-for", "(item, i) in locals.items").
							Attr("x-bind:key", "i").
							PrependIcon("item.icon").
							Title("item.text"),
					).Attr("v-model", "locals.selectedItem").
						Attr("color", "primary"),
				).Attr("dense", ""),
			).Class("mx-auto").
				Attr("max-width", "300").
				Attr("tile", ""),
		).Init(`{ selectedItem: 1, btnLabel:"Add", items: [{text: "A", icon: "clock"}]}`).
			VSlot("{ locals }"),
	)
	return
}

var UseLocalsPB = web.Page(UseLocals)

// @snippet_end

// @snippet_begin(WebScopeUsePlaidFormSample1)
var materialID, materialName, rawMaterialID, rawMaterialName, countryID, countryName, productName string

func UsePlaidForm(ctx *web.EventContext) (pr web.PageResponse, err error) {

	pr.Body = Div(
		H3("Form Content"),
		utils.PrettyFormAsJSON(ctx),

		Div(
			Div(
				Fieldset(
					Legend("Product Form"),
					Div(
						Label("Product Name"),
						Input("").
							Type("text").
							Attr(web.VField("ProductName", productName)...),
					),
					Div(
						Label("Material ID"),
						Input("").
							Type("text").Disabled(true).
							Attr(web.VField("MaterialID", materialID)...),
					),

					web.Scope(
						Fieldset(
							Legend("Material Form"),

							Div(
								Label("Material Name"),
								Input("").
									Type("text").
									Attr(web.VField("MaterialName", materialName)...),
							),
							Div(
								Label("Raw Material ID"),
								Input("").
									Type("text").Disabled(true).
									Attr(web.VField("RawMaterialID", rawMaterialID)...),
							),
							web.Scope(
								Fieldset(
									Legend("Raw Material Form"),

									Div(
										Label("Raw Material Name"),
										Input("").
											Type("text").
											Attr(web.VField("RawMaterialName", rawMaterialName)...),
									),

									Button("Send").Style(`background: orange;`).Attr("@click", web.POST().EventFunc("updateValue").Go()),
								).Style(`background: orange;`),
							).VSlot("{ plaidForm }"),

							Button("Send").Style(`background: brown;`).Attr("@click", web.POST().EventFunc("updateValue").Go()),
						).Style(`background: brown;`),
					).VSlot("{ plaidForm }"),

					Div(
						Label("Country ID"),
						Input("").
							Type("text").Disabled(true).
							Attr(web.VField("CountryID", countryID)...),
					),

					web.Scope(
						Fieldset(
							Legend("Country Of Origin Form"),

							Div(
								Label("Country Name"),
								Input("").
									Type("text").
									Attr(web.VField("CountryName", countryName)...),
							),

							Button("Send").Style(`background: red;`).Attr("@click", web.POST().EventFunc("updateValue").Go()),
						).Style(`background: red;`),
					).VSlot("{ plaidForm }"),

					Div(
						Button("Send").Style(`background: grey;`).Attr("@click", web.POST().EventFunc("updateValue").Go())),
				).Style(`background: grey;`)),
		).Style(`width:600px;`),
	)

	return
}

func updateValue(ctx *web.EventContext) (er web.EventResponse, err error) {
	ctx.R.ParseForm()
	if v := ctx.R.Form.Get("ProductName"); v != "" {
		productName = v
	}
	if v := ctx.R.Form.Get("MaterialName"); v != "" {
		materialName = v
		materialID = "66"
	}
	if v := ctx.R.Form.Get("RawMaterialName"); v != "" {
		rawMaterialName = v
		rawMaterialID = "88"
	}
	if v := ctx.R.Form.Get("CountryName"); v != "" {
		countryName = v
		countryID = "99"
	}
	er.Reload = true
	return
}

var UsePlaidFormPB = web.Page(UsePlaidForm).
	EventFunc("updateValue", updateValue)

// @snippet_end

const WebScopeUseLocalsPagePath = "/samples/web-scope-use-locals"
const WebScopeUsePlaidFormPagePath = "/samples/web-scope-use-plaid-form"
