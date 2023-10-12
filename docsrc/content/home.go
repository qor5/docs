package content

import (
	"github.com/qor5/docs/docsrc/utils"
	. "github.com/theplant/docgo"
	. "github.com/theplant/htmlgo"
)

var Home = Doc(
	Markdown(`
QOR5 is a Go library designed to help developers build web applications with ease and high customization. By focusing on static typing in the Go language and minimizing the need for JavaScript or TypeScript, QOR5 streamlines the development process and encourages reusability of components.

In QOR5, the traditional approach of using template languages for rendering HTML is discouraged. Instead, QOR5 encourages developers to write HTML [using static typing in the Go language](/advanced-functions/the-go-html-builder.html). This design choice provides several benefits:

- Improved Readability and Maintainability: By using Go's static typing, you can maintain a consistent coding style throughout the entire project, making it easier to read and maintain.

- Better Error Checking: Static typing allows for compile-time error checking, which can help catch issues before they cause problems in production.

- Enhanced Reusability: QOR5 promotes the use of components, which can be easily abstracted and reused across different parts of your application. Since components are written in Go, using third-party components from other Go packages is as simple as importing and using regular Go packages.

- Simplified Development Process: By minimizing the need for JavaScript or TypeScript, QOR5 streamlines the development process and reduces the complexity of building interactive web applications.

QOR5's approach to rendering HTML using Go's static typing eliminates the need for developers to learn and work with multiple template languages. This results in a more consistent and streamlined development experience, allowing developers to focus on the core functionality of their web applications.

	`),

	utils.Anchor(H2(""), "How is this document organized"),
	Markdown(`
Most of latter examples are based on the initial sample project. In another word, we will demonstrate how to build a rich functioned website by this document.

- Quick Sample Project: We will begin with a brief overview of a sample project, giving you a visual idea of QOR5's capabilities and functionalities.

- Basic Functions: In this section, we will explore the core features of QOR5, starting from listing pages to editing pages. This section covers common features found in admin websites.

- QOR5 Essentials and Advanced Functions: We will dive into the inner workings of QOR5, covering topics such as rendering pages and advanced features like partial page refreshing.

- Digging Deeper: In the final section, you will learn how to create new components for QOR5, extending its capabilities and adapting it to your specific needs.


**Join the Discord community**: https://discord.gg/76YPsVBE4E
`)).Title("Introduction").
	Slug("/")
