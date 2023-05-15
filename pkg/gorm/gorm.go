package gorm

import (
	"database/sql"
	"errors"
	"fmt"
	mysqlDriver "github.com/go-sql-driver/mysql"
	"goApi/internal/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strings"
	"time"
)

type Config struct {
	config.Gorm
	DSN string
}

func New(c *Config) (*gorm.DB, error) {
	var dialector gorm.Dialector

	switch strings.ToLower(c.DBType) {
	case "mysql":
		cfg, err := mysqlDriver.ParseDSN(c.DSN)
		if err != nil {
			return nil, err
		}
		err = createDatabaseWithMySQL(cfg)
		if err != nil {
			return nil, err
		}
		dialector = mysql.Open(c.DSN)
	default:
		return nil, errors.New("unknown db")
	}

	gconfig := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.TablePrefix,
			SingularTable: true,
		}}

	db, err := gorm.Open(dialector, gconfig)

	if err != nil {
		return nil, err
	}

	if c.Debug {
		db.Debug()
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(c.MaxIdleConns)
	sqlDB.SetMaxOpenConns(c.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(c.MaxLifetime) * time.Second)

	return db, nil

}

func createDatabaseWithMySQL(cfg *mysqlDriver.Config) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/", cfg.User, cfg.Passwd, cfg.Addr)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer db.Close()
	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET = `utf8mb4`;", cfg.DBName)
	_, err = db.Exec(query)
	return err
}
