package util

import "context"

type (
	TransCtx     struct{}
	NoTransCtx   struct{}
	TransLockCtx struct{}
	UserNameCtx  struct{}
	UserIDCtx    struct{}
	TraceIDCtx   struct{}
)

func NewTransKey(ctx context.Context, val any) context.Context {
	return context.WithValue(ctx, TransCtx{}, val)
}

func FromTrans(ctx context.Context) (any, bool) {
	v := ctx.Value(TransCtx{})
	return v, v != nil
}

func NewNoTransKey(ctx context.Context, val any) context.Context {
	return context.WithValue(ctx, NoTransCtx{}, val)
}

func FromNoTrans(ctx context.Context) bool {
	v := ctx.Value(NoTransCtx{})
	return v != nil && v.(bool)
}

func NewTransLockKey(ctx context.Context, val any) context.Context {
	return context.WithValue(ctx, TransLockCtx{}, val)
}

func FromTransLock(ctx context.Context) bool {
	v := ctx.Value(TransLockCtx{})
	return v != nil && v.(bool)
}

func FromStringKey(ctx context.Context, key any) string {
	v := ctx.Value(key)
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func FromIntKey(ctx context.Context, key any) int {
	v := ctx.Value(key)
	if v != nil {
		if s, ok := v.(int); ok {
			return s
		}
	}
	return 0
}

func FromUint64Key(ctx context.Context, key any) uint64 {
	v := ctx.Value(key)
	if v != nil {
		if s, ok := v.(uint64); ok {
			return s
		}
	}
	return 0
}
