package util

import (
	"context"
	"goApi/internal/app/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type Model struct {
	ID        uint64 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	trans, ok := util.FromTrans(ctx)
	if ok && !util.FromNoTrans(ctx) {
		db, ok := trans.(*gorm.DB)
		if ok {
			if util.FromTransLock(ctx) {
				db = db.Clauses(clause.Locking{Strength: "UPDATE"})
			}
		}
		return db
	}
	return defDB
}

func GetDBWithModel(ctx context.Context, defDB *gorm.DB, model any) *gorm.DB {
	return GetDB(ctx, defDB).Model(model)
}

func First(db *gorm.DB, model any) (bool, error) {
	result := db.First(model)
	if err := result.Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return false, err
		}
		return false, nil
	}
	return true, nil
}
