package buntdb

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBuntdb(t *testing.T) {
	s, err := NewStore("aacc")
	assert.Nil(t, err)
	ctx := context.Background()
	err = s.Set(ctx, "test", time.Hour*time.Duration(1))
	assert.Nil(t, err)
	b, err := s.Check(ctx, "test")
	assert.Nil(t, err)
	assert.Equal(t, true, b)
	err = s.Delete(ctx, "test")
	assert.Nil(t, err)
	b, err = s.Check(ctx, "test")
	assert.Nil(t, err)
	assert.Equal(t, false, b)
	err = s.Release()
	assert.Nil(t, err)
}
