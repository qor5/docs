package example_basics

import (
	"github.com/qor5/admin/presets"
	"github.com/qor5/admin/presets/gorm2op"
	"github.com/qor5/admin/worker"
)

func ActionWorkerExampleMock(b *presets.Builder) {
	if err := DB.AutoMigrate(&ExampleResource{}); err != nil {
		panic(err)
	}

	b.URIPrefix(ActionWorkerExamplePath).
		DataOperator(gorm2op.DataOperator(DB))

	mb := b.Model(&ExampleResource{})
	mb.Listing().ActionsAsMenu(true)

	wb := worker.NewWithQueue(DB, Que)
	wb.Configure(b)
	addActionJobs(mb, wb)
	wb.Listen()
}

const ActionWorkerExamplePath = "/samples/action_worker"
