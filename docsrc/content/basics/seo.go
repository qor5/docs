package basics

import (
	"github.com/qor5/docs/docsrc/generated"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
	"github.com/theplant/htmlgo"
)

var SEO = Doc(
	Markdown(`
The SEO library facilitates the optimization of Search Engine results by managing and injecting dynamic data into HTML tags.

## Usage
Initialize a ~~Collection~~ instance. The ~~Collection~~ manages all the registered models and hold global seo settings
~~~go
collection := seo.NewCollection()

// Turn off the default inherit the upper level SEO data when the current SEO data is missing
collection.SetInherited(false)
~~~

### Register models to SEO
~~~go
// Register mutiple SEO by name
collection.RegisterSEOByNames("Product", "Announcement")

// Register a SEO by model
type Product struct{
	Name  string
	Setting Setting
}
collection.RegisterSEO(&Product{})
~~~

### Remove models from SEO
~~~
// Remove by struct
collection.RemoveSEO(&Product{})
// Remove by name
collection.RemoveSEO("Not Found")
~~~

## Configuration

### Change the default SEO name
~~~go
collection.RegisterSEO(&Product{}).SetName("My Product")
~~~

### Register customized variables
~~~go
collection.RegisterSEO(&Product{}).
	RegisterContextVariables("og:image", func(obj interface{}, _ *Setting, _ *http.Request) string {
		// this will render "og:image" with the value of the object in the current request
		return obj.image.url
	}).
	RegisterContextVariables("Name", func(obj interface{}, _ *Setting, _ *http.Request) string {
		return obj.Name
	})
~~~

### Register setting variable
This variable will be saved in the database and available as a global variable while editing SEO settings.

~~~go
collection.RegisterSEO(&Product{}).RegisterSettingVaribles(struct{ProductTag string}{})
~~~

### Render SEO html data

~~~go
// Render Global SEO
collection.RenderGlobal(request)

// Render SEO by name
collection.Render("product", request)

// Render SEO by model
collection.Render(Product{}, request)
~~~

## Customization
`),
	Markdown(`
You can customize your SEO settings by implementing the interface and adding functions such as l10n and publish.`),
	ch.Code(generated.QorSEOSettingInterface).Language("go"),

	htmlgo.H2("Example"),
	ch.Code(generated.SeoExample).Language("go"),

	Markdown(`
## Definition
~~Collection~~ manages all the registered models and hold global seo settings.`),
	ch.Code(generated.SeoCollectionDefinition).Language("go"),

	Markdown(`
~~SEO~~ provides system-level default page matadata.`),
	ch.Code(generated.SeoDefinition).Language("go"),

	Markdown(`
You can use seo setting at the model level, but you need to register the model to the system SEO`),
	ch.Code(generated.SeoModelExample).Language("go"),
	ch.Code(`collection.RegisterSEO(&Product{})`).Language("go"),
).Title("SEO")
