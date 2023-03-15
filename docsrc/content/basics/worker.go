package basics

import (
	"fmt"
	"path"

	"github.com/qor5/docs/docsrc/examples/example_basics"
	"github.com/qor5/docs/docsrc/generated"
	"github.com/qor5/docs/docsrc/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var Worker = Doc(
	Markdown(fmt.Sprintf(`
Worker runs a single Job in the background, it can do so immediately or at a scheduled time.  
Once registered with QOR Admin, Worker will provide a Workers section in the navigation tree, containing pages for listing and managing the following aspects of Workers:

- All Jobs.
- Running: Jobs that are currently running.
- Scheduled: Jobs which have been scheduled to run at a time in the future.
- Done: finished Jobs.
- Errors: any errors output from any Workers that have been run.

## Note
- The default que GoQueQueue(https://github.com/tnclong/go-que) only supports postgres for now.
- To make a job abortable, you need to check %s channel in job handler and stop the handler func.
    `, "`ctx.Done()`")),
	Markdown(`
## Example
`),
	ch.Code(generated.WorkerExample).Language("go"),
	utils.Demo(
		"Worker",
		path.Join(example_basics.WorkerExamplePath, "/workers"),
		"example_basics/worker.go",
	),
).Slug("basics/worker").Title("Worker")
