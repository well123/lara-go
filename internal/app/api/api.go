package api

import "github.com/google/wire"

var Set = wire.NewSet(LoginSet, UserSet)
