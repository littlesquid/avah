package cron

import (
	"fmt"
	"os"

	"github.com/robfig/cron"
)

func Run() {

	cronExpression := os.Getenv("CRON_EXP")

	c := cron.New()
	c.AddFunc(cronExpression, func() { fmt.Println("Every second") })
	c.Start()

}
