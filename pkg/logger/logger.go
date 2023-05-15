package logger

import (
	"context"
	"github.com/sirupsen/logrus"
)

const (
	TraceIdKey  = "trace_id"
	TagKey      = "tag"
	UserIDKey   = "user_id"
	UserNameKey = "user_name"
	StackKey    = "stack"
)

type Entry = logrus.Entry
type Fields = logrus.Fields
type Level = logrus.Level
type (
	traceIdKey  struct{}
	tagKey      struct{}
	stackKey    struct{}
	userIDKey   struct{}
	userNameKey struct{}
)

func SetFormatter(format string) {
	switch format {
	case "json":
		logrus.SetFormatter(new(logrus.JSONFormatter))
	default:
		logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	}
}

func NewTraceIdCtx(ctx context.Context, traceId string) context.Context {
	return context.WithValue(ctx, traceIdKey{}, traceId)
}

func FromTraceIdCtx(ctx context.Context) string {
	v := ctx.Value(traceIdKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func NewTagCtx(ctx context.Context, tag string) context.Context {
	return context.WithValue(ctx, tagKey{}, tag)
}

func FromTagCtx(ctx context.Context) string {
	v := ctx.Value(tagKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func NewUserIDCtx(ctx context.Context, userID uint64) context.Context {
	return context.WithValue(ctx, userIDKey{}, userID)
}

func FromUserIDCtx(ctx context.Context) uint64 {
	v := ctx.Value(userIDKey{})
	if v != nil {
		if s, ok := v.(uint64); ok {
			return s
		}
	}
	return 0
}

func NewStackCtx(ctx context.Context, stack string) context.Context {
	return context.WithValue(ctx, stackKey{}, stack)
}

func FromStackCtx(ctx context.Context) error {
	v := ctx.Value(stackKey{})
	if v != nil {
		if s, ok := v.(error); ok {
			return s
		}
	}
	return nil
}

func NewUserNameCtx(ctx context.Context, userName string) context.Context {
	return context.WithValue(ctx, userNameKey{}, userName)
}

func FromUserNameCtx(ctx context.Context) string {
	v := ctx.Value(userNameKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func WithStackContext(ctx context.Context, stack error) context.Context {
	return context.WithValue(ctx, stackKey{}, stack)
}

func WithContext(ctx context.Context) *Entry {
	fields := Fields{}

	if v := FromTraceIdCtx(ctx); v != "" {
		fields[TraceIdKey] = v
	}
	if v := FromTagCtx(ctx); v != "" {
		fields[TagKey] = v
	}
	if v := FromUserNameCtx(ctx); v != "" {
		fields[UserNameKey] = v
	}
	if v := FromUserIDCtx(ctx); v != 0 {
		fields[UserIDKey] = v
	}
	if v := FromStackCtx(ctx); v != nil {
		fields[StackKey] = v
	}

	return logrus.WithContext(ctx).WithFields(fields)
}

var (
	SetLevel  = logrus.SetLevel
	SetOutput = logrus.SetOutput
)
