package getting_started

import (
	. "github.com/theplant/docgo"
)

var OneMinuteQuickStart = Doc(
	Markdown(`
This article try to let you use the shortest time to get a taste of how powerful QOR5 is.

One of the QOR5 module called ~presets~ that can quickly create admin interface like [these](/samples/presets-detail-page-cards/customers):

Install the command line tool with:

~~~
$ go install github.com/qor5/docs/cmd/qor5@latest
~~~

And run:

~~~
$ qor5
~~~

It will promote you to input a Go package, and create the admin app in current directory.

Change to the created package directory, and use ~docker-compose up~ to start the database, and then
Use a new terminal to run ~source dev_env && go run main.go~ to start the admin app

`),
).Title("1 Minute Quick Start").
	Slug("getting-started/one-minute-quick-start")
