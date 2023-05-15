package redis

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStore(t *testing.T) {

	store := NewStore(&Config{
		Addr:     "127.0.0.1:6379",
		DB:       1,
		Password: "foobared",
		Prefix:   "aaa",
	})

	defer store.Release()

	key := "test"
	ctx := context.Background()
	err := store.Set(ctx, key, 0)
	assert.Nil(t, err)

	b, err := store.Check(ctx, key)
	assert.Nil(t, err)
	assert.Equal(t, true, b)

	b, err = store.Delete(ctx, key)
	assert.Nil(t, err)
	assert.Equal(t, true, b)

}
