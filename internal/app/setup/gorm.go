package setup

import (
	"errors"
	"goApi/internal/app/config"
	"goApi/internal/app/dao"
	gormw "goApi/pkg/gorm"
	"gorm.io/gorm"
)

func Gorm() (*gorm.DB, func(), error) {
	cfg := config.C.Gorm
	db, err := newGorm()

	if err != nil {
		println(err.Error())
	}
	clearGormFunc := func() {}

	if cfg.EnableAutoMigrate {
		err := dao.AutoMigrate(db)
		if err != nil {
			return nil, clearGormFunc, err
		}
	}

	return db, clearGormFunc, nil
}

func newGorm() (*gorm.DB, error) {
	c := config.C
	var dsn string
	switch c.Gorm.DBType {
	case "mysql":
		dsn = c.Mysql.DSN()
	default:
		return nil, errors.New("unknown db")
	}
	return gormw.New(&gormw.Config{
		Gorm: c.Gorm,
		DSN:  dsn,
	})
}
