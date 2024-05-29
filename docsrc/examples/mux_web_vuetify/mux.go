package mux_web_vuetify

import (
	"net/http"
	"strings"

	"github.com/qor5/docs/v3/docsrc/assets"
	"github.com/qor5/docs/v3/docsrc/examples/examples_vuetify"
	"github.com/qor5/docs/v3/docsrc/examples/examples_vuetifyx"
	"github.com/qor5/docs/v3/docsrc/examples/examples_web"
	"github.com/qor5/ui/v3/tiptap"
	. "github.com/qor5/ui/v3/vuetify"
	"github.com/qor5/ui/v3/vuetifyx"
	"github.com/qor5/web/v3"
	. "github.com/theplant/htmlgo"
)

type IndexMux struct {
	Mux   *http.ServeMux
	paths []string
}

func (im *IndexMux) Page(ctx *web.EventContext) (r web.PageResponse, err error) {
	ul := Ul()
	for _, p := range im.paths {
		ul.AppendChildren(Li(A().Href(p).Text(p)))
	}
	r.Body = ul
	return
}

func (im *IndexMux) Handle(pattern string, handler http.Handler) {
	im.paths = append(im.paths, pattern)
	im.Mux.Handle(pattern, handler)
	im.Mux.Handle(pattern+"/", handler)
}

func AddGA(ctx *web.EventContext) {
	if strings.Index(ctx.R.Host, "localhost") >= 0 {
		return
	}
	ctx.Injector.HeadHTML(`
<!-- Global site tag (gtag.js) - Google Analytics -->
<script async src="https://www.googletagmanager.com/gtag/js?id=UA-149605708-1"></script>
<script>
  window.dataLayer = window.dataLayer || [];
  function gtag(){dataLayer.push(arguments);}
  gtag('js', new Date());

  gtag('config', 'UA-149605708-1');
</script>
`)
}

// @snippet_begin(DemoLayoutSample)
func demoLayout(in web.PageFunc) (out web.PageFunc) {
	return func(ctx *web.EventContext) (pr web.PageResponse, err error) {
		AddGA(ctx)

		ctx.Injector.HeadHTML(`
			<script src='/assets/vue.js'></script>
		`)

		ctx.Injector.TailHTML(`
			<script src='/assets/main.js'></script>
		`)
		ctx.Injector.HeadHTML(`
		<style>
			[v-cloak] {
				display: none;
			}
		</style>
		`)

		var innerPr web.PageResponse
		innerPr, err = in(ctx)
		if err != nil {
			panic(err)
		}

		pr.Body = innerPr.Body

		return
	}
}

// @snippet_end

// @snippet_begin(TipTapLayoutSample)
func tiptapLayout(in web.PageFunc) (out web.PageFunc) {
	return func(ctx *web.EventContext) (pr web.PageResponse, err error) {
		AddGA(ctx)

		ctx.Injector.HeadHTML(`
			<link rel="stylesheet" href="/assets/tiptap.css">
			<script src='/assets/vue.js'></script>
		`)

		ctx.Injector.TailHTML(`
<script src='/assets/tiptap.js'></script>
<script src='/assets/main.js'></script>
`)
		ctx.Injector.HeadHTML(`
		<style>
			[v-cloak] {
				display: none;
			}
		</style>
		`)

		var innerPr web.PageResponse
		innerPr, err = in(ctx)
		if err != nil {
			panic(err)
		}

		pr.Body = innerPr.Body

		return
	}
}

// @snippet_end

// @snippet_begin(DemoBootstrapLayoutSample)
func demoBootstrapLayout(in web.PageFunc) (out web.PageFunc) {
	return func(ctx *web.EventContext) (pr web.PageResponse, err error) {
		AddGA(ctx)

		ctx.Injector.HeadHTML(`
<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">
<script src='/assets/vue.js'></script>
		`)

		ctx.Injector.TailHTML(`
<script src="https://code.jquery.com/jquery-3.3.1.slim.min.js" integrity="sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo" crossorigin="anonymous"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js" integrity="sha384-UO2eT0CpHqdSJQ6hJty5KVphtPhzWj9WO1clHTMGa3JDZwrnQq4sF86dIHNDz0W1" crossorigin="anonymous"></script>
<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js" integrity="sha384-JjSmVgyd0p3pXB1rRibZUAYoIIy6OrQ6VrjIEaFf/nJGzIxFDsf4x0xIM+B07jRM" crossorigin="anonymous"></script>
<script src='/assets/main.js'></script>

`)
		ctx.Injector.HeadHTML(`
		<style>
			[v-cloak] {
				display: none;
			}
		</style>
		`)

		var innerPr web.PageResponse
		innerPr, err = in(ctx)
		if err != nil {
			panic(err)
		}

		pr.Body = innerPr.Body

		return
	}
}

// @snippet_end

// @snippet_begin(DemoVuetifyLayoutSample)
func DemoVuetifyLayout(in web.PageFunc) (out web.PageFunc) {
	return func(ctx *web.EventContext) (pr web.PageResponse, err error) {
		AddGA(ctx)

		ctx.Injector.HeadHTML(`
			<link rel="stylesheet" href="/vuetify/assets/index.css">
			<script src='/assets/vue.js'></script>
		`)

		ctx.Injector.TailHTML(`
			<script src='/assets/main.js'></script>
		`)
		ctx.Injector.HeadHTML(`
		<style>
			[v-cloak] {
				display: none;
			}
		</style>
		`)

		var innerPr web.PageResponse
		innerPr, err = in(ctx)
		if err != nil {
			panic(err)
		}

		pr.Body = VApp(
			VMain(
				innerPr.Body,
			),
		)

		return
	}
}

// @snippet_end

func Mux(mux *http.ServeMux, prefix string) http.Handler {
	// @snippet_begin(ComponentsPackSample)
	mux.Handle("/assets/main.js",
		web.PacksHandler("text/javascript",
			vuetifyx.JSComponentsPack(),
			Vuetify(),
			JSComponentsPack(),
			web.JSComponentsPack(),
		),
	)

	mux.Handle("/assets/vue.js",
		web.PacksHandler("text/javascript",
			web.JSVueComponentsPack(),
		),
	)

	// @snippet_end

	// @snippet_begin(TipTapComponentsPackSample)
	mux.Handle("/assets/tiptap.js",
		web.PacksHandler("text/javascript",
			tiptap.JSComponentsPack(),
		),
	)

	mux.Handle("/assets/tiptap.css",
		web.PacksHandler("text/css",
			tiptap.CSSComponentsPack(),
		),
	)
	// @snippet_end

	// @snippet_begin(VuetifyComponentsPackSample)
	HandleMaterialDesignIcons("", mux)
	// @snippet_end

	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.Write(assets.Favicon)
		return
	})

	return mux
}

type Muxer interface {
	Handle(pattern string, handler http.Handler)
}

func SamplesHandler(mux Muxer, prefix string) {
	emptyUb := web.New().LayoutFunc(web.NoopLayoutFunc)

	mux.Handle(examples_web.TypeSafeBuilderSamplePath, examples_web.TypeSafeBuilderSamplePFPB.Builder(emptyUb))

	// @snippet_begin(HelloWorldMuxSample2)
	mux.Handle(examples_web.HelloWorldPath, examples_web.HelloWorldPB)
	// @snippet_end

	// @snippet_begin(HelloWorldReloadMuxSample1)
	mux.Handle(
		examples_web.HelloWorldReloadPath,
		examples_web.HelloWorldReloadPB.Wrap(demoLayout),
	)
	// @snippet_end

	mux.Handle(
		examples_web.HelloButtonPath,
		examples_web.HelloButtonPB.Wrap(demoLayout),
	)

	mux.Handle(
		examples_web.Page1Path,
		examples_web.Page1PB.Wrap(demoLayout),
	)
	mux.Handle(
		examples_web.Page2Path,
		examples_web.Page2PB.Wrap(demoLayout),
	)

	mux.Handle(
		examples_web.ReloadWithFlashPath,
		examples_web.ReloadWithFlashPB.Wrap(demoLayout),
	)

	mux.Handle(
		examples_web.PartialUpdatePagePath,
		examples_web.PartialUpdatePagePB.Wrap(demoLayout),
	)

	mux.Handle(
		examples_web.PartialReloadPagePath,
		examples_web.PartialReloadPagePB.Wrap(demoLayout),
	)

	mux.Handle(
		examples_web.MultiStatePagePath,
		examples_web.MultiStatePagePB.Wrap(demoLayout),
	)

	mux.Handle(
		examples_web.FormHandlingPagePath,
		examples_web.FormHandlingPagePB.Wrap(demoLayout),
	)

	mux.Handle(
		examples_web.CompositeComponentSample1PagePath,
		examples_web.CompositeComponentSample1PagePB.Wrap(demoBootstrapLayout),
	)

	mux.Handle(
		examples_web.HelloWorldTipTapPath,
		examples_web.HelloWorldTipTapPB.Wrap(tiptapLayout),
	)

	mux.Handle(
		examples_web.EventExamplePagePath,
		examples_web.ExamplePagePB.Wrap(DemoVuetifyLayout),
	)

	mux.Handle(
		examples_web.EventHandlingPagePath,
		examples_web.EventHandlingPagePB.Wrap(DemoVuetifyLayout),
	)

	mux.Handle(
		examples_web.WebScopeUseLocalsPagePath,
		examples_web.UseLocalsPB.Wrap(DemoVuetifyLayout),
	)

	mux.Handle(
		examples_web.WebScopeUsePlaidFormPagePath,
		examples_web.UsePlaidFormPB.Wrap(demoLayout),
	)

	mux.Handle(
		examples_web.ShortCutSamplePath,
		examples_web.ShortCutSamplePB.Wrap(DemoVuetifyLayout),
	)

	mux.Handle(
		examples_vuetify.HelloVuetifyListPath,
		examples_vuetify.HelloVuetifyListPB.Wrap(DemoVuetifyLayout),
	)

	mux.Handle(
		examples_vuetify.HelloVuetifyMenuPath,
		examples_vuetify.HelloVuetifyMenuPB.Wrap(DemoVuetifyLayout),
	)

	mux.Handle(
		examples_vuetify.VuetifyGridPath,
		examples_vuetify.VuetifyGridPB.Wrap(DemoVuetifyLayout),
	)

	mux.Handle(
		examples_vuetify.VuetifyBasicInputsPath,
		examples_vuetify.VuetifyBasicInputsPB.Wrap(DemoVuetifyLayout),
	)

	mux.Handle(
		examples_vuetify.VuetifyVariantSubFormPath,
		examples_vuetify.VuetifyVariantSubFormPB.Wrap(DemoVuetifyLayout),
	)

	mux.Handle(
		examples_vuetify.VuetifyComponentsKitchenPath,
		examples_vuetify.VuetifyComponentsKitchenPB.Wrap(DemoVuetifyLayout),
	)

	mux.Handle(
		examples_vuetify.VuetifyNavigationDrawerPath,
		examples_vuetify.VuetifyNavigationDrawerPB.Wrap(DemoVuetifyLayout),
	)

	mux.Handle(
		examples_vuetify.LazyPortalsAndReloadPath,
		examples_vuetify.LazyPortalsAndReloadPB.Wrap(DemoVuetifyLayout),
	)

	mux.Handle(
		examples_vuetify.VuetifySnackBarsPath,
		examples_vuetify.VuetifySnackBarsPB.Wrap(DemoVuetifyLayout),
	)

	mux.Handle(
		examples_vuetifyx.VuetifyComponentsLinkageSelectPath,
		examples_vuetifyx.VuetifyComponentsLinkageSelectPB.Wrap(DemoVuetifyLayout),
	)

	return
}
