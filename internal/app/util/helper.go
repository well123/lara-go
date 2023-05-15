package util

import (
	"github.com/jinzhu/copier"
)

func Copy(s, t any) error {
	return copier.Copy(t, s)
}
