package cron

import (
	"time"
)

type AutoSetOrder struct {
}

func (*AutoSetOrder) Run() {
	println("Every hour thirty " + time.Now().String())
}
