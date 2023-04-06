package basics

import (
	"path"

	"github.com/qor5/docs/docsrc/examples/example_basics"
	"github.com/qor5/docs/docsrc/generated"
	"github.com/qor5/docs/docsrc/utils"
	"github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var I18n = docgo.Doc(
	docgo.Markdown(`
The [i18n package](https://github.com/qor5/x/tree/master/i18n) provides support for internationalization (i18n) in Go applications. 
It allows you to define and load messages in multiple languages, and provides a mechanism 
for selecting the appropriate language for a given user request.
    `),
	docgo.Markdown(`
## Getting Started
To use the i18n package, you first need to import it into your Go application:
`),
	ch.Code(`import "github.com/qor5/x/i18n"`).Language("go"),
	docgo.Markdown(`
Next, create a new ~Builder~ instance using the ~New()~ function. 
If you want to use it in QOR5, use the ~I18n()~ on ~presets.Builder~:
`),
	ch.Code(generated.I18nNew).Language("go"),
	docgo.Markdown(`
The ~Builder~ struct is the central point of the i18n package. 
It holds the supported languages, the messages for each module in each language, 
and the configuration for retrieving the language preference.
`),
	docgo.Markdown(`
## Supporting Languages
To support multiple languages in your web application, you need to define the languages that you support. 
You can do this by calling the ~SupportLanguages~ function on the ~Builder~ struct:
`),
	ch.Code(generated.I18nSupportLanguages).Language("go"),
	docgo.Markdown(`
By default, the ~Builder~ struct only supports the English language. 
You can add more languages by calling the ~SupportLanguages~ function and passing in the language tags.
`),
	docgo.Markdown(`
## Registering Module Messages
Once you have defined the languages that you support, you need to register messages for each module in each language. 
You can do this by calling the ~RegisterForModule~ function on the ~Builder~ struct:
`),
	ch.Code(generated.I18nRegisterForModule).Language("go"),
	docgo.Markdown(`
The ~RegisterForModule~ function takes three arguments: the language tag, the module name, 
and a pointer to a struct that implements the Messages interface. 
The Messages interface is an empty interface that you can use to define your own messages.

Such a struct might look like this:
`),
	ch.Code(generated.I18nMessagesExample).Language("go"),
	docgo.Markdown(`
The ~GetSupportLanguagesFromRequestFunc~ is a method of the ~Builder~ struct in the i18n package. 
It allows you to set a function that retrieves the list of supported languages 
from an HTTP request, which can be useful in scenarios where the list of supported 
languages varies based on the request context.

The ~EnsureLanguage~ function is an HTTP middleware that ensures the request's language 
is properly set and stored. It does this by first checking the query parameters for 
a language value, and if found, setting a cookie with that value. If no language 
value is present in the query parameters, it looks for the language value in the cookie.

The middleware then determines the best-matching language from the supported languages 
based on the "Accept-Language" header of the request. If no match is found, 
it defaults to the first supported language. It then sets the language context for 
the request, which can be retrieved later by calling the ~MustGetModuleMessages~ function.
`),
	docgo.Markdown(`
## Retrieving Messages
To retrieve module messages in your HTTP handler, you can use the ~MustGetModuleMessages~ function:
`),
	ch.Code(generated.I18nMustGetModuleMessages).Language("go"),
	docgo.Markdown(`
The ~MustGetModuleMessages~ function takes three arguments: 
the HTTP request, the module name, and a pointer to a struct 
that implements the Messages interface. The function retrieves the messages 
for the specified module in the language set by the i18n middleware.
`),
	docgo.Markdown(`
## Conclusion
The i18n package provides a simple and efficient way to build multilingual 
web applications in Go. With the package, you can support multiple languages, 
register messages for each module in each language, and serve multilingual content 
based on the user's preferences.
`),
	utils.Demo(
		"I18n",
		path.Join(example_basics.InternationalizationExamplePath, "/home"),
		"example_basics/internationalization.go",
	),
).Slug("basics/i18n").Title("Internationalization")
