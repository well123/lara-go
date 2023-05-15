package buntdb

import (
	"context"
	"github.com/tidwall/buntdb"
	"os"
	"path/filepath"
	"time"
)

type Store struct {
	db *buntdb.DB
}

func NewStore(filePath string) (*Store, error) {
	if filePath != ":memory:" {
		_ = os.MkdirAll(filepath.Dir(filePath), 0777)
		_, _ = os.Create("buntdb.db")
	}

	db, err := buntdb.Open(filePath)
	if err != nil {
		return nil, err
	}
	return &Store{db: db}, nil
}

func (s *Store) Set(ctx context.Context, tokenString string, expiration time.Duration) error {
	return s.db.Update(func(tx *buntdb.Tx) error {
		var opts *buntdb.SetOptions
		if expiration > 0 {
			opts = &buntdb.SetOptions{Expires: true, TTL: expiration}
		}
		_, _, err := tx.Set(tokenString, "1", opts)
		return err
	})
}

func (s *Store) Delete(ctx context.Context, tokenString string) error {
	return s.db.Update(func(tx *buntdb.Tx) error {
		_, err := tx.Delete(tokenString)
		if err != nil && err != buntdb.ErrNotFound {
			return err
		}
		return nil
	})
}

func (s *Store) Check(ctx context.Context, tokenString string) (bool, error) {
	var exists bool
	err := s.db.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get(tokenString)
		if err != nil && err != buntdb.ErrNotFound {
			return err
		}
		exists = val == "1"
		return nil
	})
	return exists, err
}

func (s *Store) Release() error {
	return s.db.Close()
}
