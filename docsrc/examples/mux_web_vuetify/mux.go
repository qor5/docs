package mux_web_vuetify

import (
	"github.com/qor5/docs/docsrc/assets"
	"github.com/qor5/docs/docsrc/examples/vuetify_examples"
	"github.com/qor5/docs/docsrc/examples/vuetifyx_examples"
	"github.com/qor5/docs/docsrc/examples/web_examples"
	"github.com/qor5/ui/tiptap"
	. "github.com/qor5/ui/vuetify"
	"github.com/qor5/web"
	. "github.com/theplant/htmlgo"
	"net/http"
	"strings"
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
			<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto+Mono" async>
			<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto:300,400,500" async>
			<link href="https://cdn.jsdelivr.net/npm/@mdi/font@5.x/css/materialdesignicons.min.css" rel="stylesheet" async>
			<link rel="stylesheet" href="/assets/vuetify.css">
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
			Vuetify(""),
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
	mux.Handle("/assets/vuetify.css",
		web.PacksHandler("text/css",
			CSSComponentsPack(),
		),
	)
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

	mux.Handle(web_examples.TypeSafeBuilderSamplePath, web_examples.TypeSafeBuilderSamplePFPB.Builder(emptyUb))

	// @snippet_begin(HelloWorldMuxSample2)
	mux.Handle(web_examples.HelloWorldPath, web_examples.HelloWorldPB)
	// @snippet_end

	// @snippet_begin(HelloWorldReloadMuxSample1)
	mux.Handle(
		web_examples.HelloWorldReloadPath,
		web_examples.HelloWorldReloadPB.Wrap(demoLayout),
	)
	// @snippet_end

	mux.Handle(
		web_examples.HelloButtonPath,
		web_examples.HelloButtonPB.Wrap(demoLayout),
	)

	mux.Handle(
		web_examples.Page1Path,
		web_examples.Page1PB.Wrap(demoLayout),
	)
	mux.Handle(
		web_examples.Page2Path,
		web_examples.Page2PB.Wrap(demoLayout),
	)

	mux.Handle(
		web_examples.ReloadWithFlashPath,
		web_examples.ReloadWithFlashPB.Wrap(demoLayout),
	)

	mux.Handle(
		web_examples.PartialUpdatePagePath,
		web_examples.PartialUpdatePagePB.Wrap(demoLayout),
	)

	mux.Handle(
		web_examples.PartialReloadPagePath,
		web_examples.PartialReloadPagePB.Wrap(demoLayout),
	)

	mux.Handle(
		web_examples.MultiStatePagePath,
		web_examples.MultiStatePagePB.Wrap(demoLayout),
	)

	mux.Handle(
		web_examples.FormHandlingPagePath,
		web_examples.FormHandlingPagePB.Wrap(demoLayout),
	)

	mux.Handle(
		web_examples.CompositeComponentSample1PagePath,
		web_examples.CompositeComponentSample1PagePB.Wrap(demoBootstrapLayout),
	)

	mux.Handle(
		web_examples.HelloWorldTipTapPath,
		web_examples.HelloWorldTipTapPB.Wrap(tiptapLayout),
	)

	mux.Handle(
		web_examples.EventExamplePagePath,
		web_examples.ExamplePagePB.Wrap(DemoVuetifyLayout),
	)

	mux.Handle(
		web_examples.EventHandlingPagePath,
		web_examples.EventHandlingPagePB.Wrap(DemoVuetifyLayout),
	)

	mux.Handle(
		web_examples.WebScopeUseLocalsPagePath,
		web_examples.UseLocalsPB.Wrap(DemoVuetifyLayout),
	)

	mux.Handle(
		web_examples.WebScopeUsePlaidFormPagePath,
		web_examples.UsePlaidFormPB.Wrap(demoLayout),
	)

	mux.Handle(
		web_examples.ShortCutSamplePath,
		web_examples.ShortCutSamplePB.Wrap(DemoVuetifyLayout),
	)

	mux.Handle(
		vuetify_examples.HelloVuetifyListPath,
		vuetify_examples.HelloVuetifyListPB.Wrap(DemoVuetifyLayout),
	)

	mux.Handle(
		vuetify_examples.HelloVuetifyMenuPath,
		vuetify_examples.HelloVuetifyMenuPB.Wrap(DemoVuetifyLayout),
	)

	mux.Handle(
		vuetify_examples.VuetifyGridPath,
		vuetify_examples.VuetifyGridPB.Wrap(DemoVuetifyLayout),
	)

	mux.Handle(
		vuetify_examples.VuetifyBasicInputsPath,
		vuetify_examples.VuetifyBasicInputsPB.Wrap(DemoVuetifyLayout),
	)

	mux.Handle(
		vuetify_examples.VuetifyVariantSubFormPath,
		vuetify_examples.VuetifyVariantSubFormPB.Wrap(DemoVuetifyLayout),
	)

	mux.Handle(
		vuetify_examples.VuetifyComponentsKitchenPath,
		vuetify_examples.VuetifyComponentsKitchenPB.Wrap(DemoVuetifyLayout),
	)

	mux.Handle(
		vuetify_examples.VuetifyNavigationDrawerPath,
		vuetify_examples.VuetifyNavigationDrawerPB.Wrap(DemoVuetifyLayout),
	)

	mux.Handle(
		vuetify_examples.LazyPortalsAndReloadPath,
		vuetify_examples.LazyPortalsAndReloadPB.Wrap(DemoVuetifyLayout),
	)

	mux.Handle(
		vuetify_examples.VuetifySnackBarsPath,
		vuetify_examples.VuetifySnackBarsPB.Wrap(DemoVuetifyLayout),
	)

	mux.Handle(
		vuetifyx_examples.VuetifyComponentsLinkageSelectPath,
		vuetifyx_examples.VuetifyComponentsLinkageSelectPB.Wrap(DemoVuetifyLayout),
	)

	return
}
