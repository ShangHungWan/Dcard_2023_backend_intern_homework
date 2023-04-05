package scheduler

import (
	"key-value-system/db"
	"time"

	"github.com/go-co-op/gocron"
)

func init() {
	scheduler := gocron.NewScheduler(time.UTC)

	registerScheduler(scheduler)

	scheduler.StartAsync()
}

func registerScheduler(scheduler *gocron.Scheduler) {
	scheduler.Every(1).Day().Do(clearOldNodes)
}

func clearOldNodes() {
	db.DB.Exec(db.GetSql(db.CLEAR_OLD_NODES_SQL))
}
