package setup

import "goApi/internal/app/cron"

func Cron() (*cron.Cron, func(), error) {
	c2 := cron.Cron{}
	c2.Start()
	return &c2, c2.Stop, nil
}
