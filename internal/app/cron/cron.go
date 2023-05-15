package cron

import (
	"github.com/robfig/cron"
)

type Cron struct {
	cron *cron.Cron
}

func (a *Cron) Start() *Cron {
	c := cron.New()
	//c.AddFunc("0 30 * * * *", func() { println("Every hour on the half hour") })
	//c.AddFunc("@hourly", func() { println("Every hour") })
	//c.AddFunc("@every 5s", func() { println("Every hour thirty " + time.Now().String()) })
	//c.AddFunc("* * * * * *", func() {
	//	println("Every minutes " + time.Now().String())
	//})
	//c.AddJob("*/5 * * * * *", &AutoSetOrder{})

	c.Start()
	a.cron = c

	return a
}

func (a *Cron) Stop() {
	a.cron.Stop()
}
