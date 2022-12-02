package main

import (
	"github.com/qor5/docs/docsrc"
	"github.com/theplant/docgo"
)

func main() {
	docgo.New().
		Assets("/assets/", docsrc.Assets).
		MainPageTitle("QOR5 Document").
		SitePrefix("/docs/").
		DocTree(docsrc.DocTree...).
		BuildStaticSite("../docs")
}
