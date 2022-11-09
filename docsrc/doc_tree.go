package docsrc

import (
	"github.com/qor5/docs/docsrc/content"
	advanced_functions "github.com/qor5/docs/docsrc/content/advanced-functions"
	"github.com/qor5/docs/docsrc/content/basics"
	digging_deeper "github.com/qor5/docs/docsrc/content/digging-deeper"
	getting_started "github.com/qor5/docs/docsrc/content/getting-started"
	"github.com/qor5/docs/docsrc/utils"
	"github.com/theplant/docgo"
)

var DocTree = []interface{}{
	content.Home,
	&docgo.DocsGroup{
		Title: "Getting Started",
		Docs: []*docgo.DocBuilder{
			getting_started.OneMinuteQuickStart,
		},
	},
	&docgo.DocsGroup{
		Title: "Basics",
		Docs: []*docgo.DocBuilder{
			basics.Listing,
			basics.Filter,
			basics.EditingCustomizations,
			basics.FormHandling,
			basics.BasicInputs,
			basics.AutoComplete,
			basics.ShortCut,
			basics.VariantSubForm,
			basics.LinkageSelect,
			basics.Permissions,
			basics.NotificationCenter,
		},
	},

	&docgo.DocsGroup{
		Title: "Advanced Functions",
		Docs: []*docgo.DocBuilder{
			advanced_functions.PageFuncAndEventFunc,
			advanced_functions.TheGoHTMLBuilder,
			advanced_functions.ATasteOfUsingVuetifyInGo,
			advanced_functions.ItsTheWholeHouse,
			advanced_functions.NavigationDrawer,
			advanced_functions.LazyPortalsAndReload,
			advanced_functions.LayoutFunctionAndPageInjector,
			advanced_functions.SwitchPagesWithPushState,
			advanced_functions.ReloadPageWithAFlash,
			advanced_functions.PartialRefreshWithPortal,
			advanced_functions.ManipulatePageURLInEventFunc,
			advanced_functions.SummaryOfEventResponse,
			advanced_functions.WebScope,
			advanced_functions.EventHandling,
			advanced_functions.DetailPageForComplexObject,
		},
	},
	&docgo.DocsGroup{
		Title: "Digging Deeper",
		Docs: []*docgo.DocBuilder{
			digging_deeper.CompositeNewComponentWithGo,
			digging_deeper.IntegrateAHeavyVueComponent,
		},
	},
	&docgo.DocsGroup{
		Title: "Appendix",
		Docs: []*docgo.DocBuilder{
			docgo.Doc(utils.ExamplesDoc()).
				Title("All Demo Examples").
				Slug("appendix/all-demo-examples"),
		},
	},
}
