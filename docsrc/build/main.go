package main

import (
	"github.com/theplant/docgo"
	"github.com/qor5/docs/docsrc"
)

func main() {
	docgo.New().
		Assets("/assets/", docsrc.Assets).
		MainPageTitle("My Document").
		SitePrefix("/docs/").
		DocTree(docsrc.DocTree...).
		BuildStaticSite("../docs")
}
