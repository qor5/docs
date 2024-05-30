package examples_admin

import (
	"github.com/qor5/admin/v3/presets"
	"github.com/qor5/admin/v3/presets/gorm2op"
	"github.com/qor5/admin/v3/worker"
	"gorm.io/gorm"
)

func ActionWorkerExample(b *presets.Builder, db *gorm.DB) {
	if err := db.AutoMigrate(&ExampleResource{}); err != nil {
		panic(err)
	}

	b.DataOperator(gorm2op.DataOperator(db))

	mb := b.Model(&ExampleResource{})
	mb.Listing().ActionsAsMenu(true)

	wb := worker.NewWithQueue(db, Que)
	b.Use(wb)
	addActionJobs(mb, wb)
	wb.Listen()
}

const ActionWorkerExamplePath = "/samples/action_worker"
