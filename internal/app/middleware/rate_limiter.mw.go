package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/go-redis/redis_rate"
	"goApi/internal/app/config"
	"goApi/internal/app/util"
	"goApi/pkg/errors"
	"golang.org/x/time/rate"
	"strconv"
	"time"
)

func RateLimiterMiddleware(skippers ...SkipperFunc) gin.HandlerFunc {

	cfg := config.C.RateLimiter

	if !cfg.Enable {
		return EmptyHandler()
	}

	rc := config.C.Redis
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": rc.Addr,
		},
		Password: rc.Password,
		DB:       cfg.RedisDB,
	})

	limter := redis_rate.NewLimiter(ring)
	limter.Fallback = rate.NewLimiter(rate.Inf, 0)

	return func(context *gin.Context) {
		if SkipHandler(context, skippers...) {
			context.Next()
			return
		}

		userID := util.FromUint64Key(context.Request.Context(), util.UserIDCtx{})

		if userID != 0 {
			limit := cfg.Count
			count, delay, allow := limter.AllowMinute(fmt.Sprintf("%d", userID), limit)
			if !allow {
				h := context.Writer.Header()
				h.Set("X-RateLimit-Limit", strconv.FormatInt(limit, 10))
				h.Set("X-RateLimit-Remaining", strconv.FormatInt(limit-count, 10))
				delaySec := int64(delay / time.Second)
				h.Set("X-RateLimit-Delay", strconv.FormatInt(delaySec, 10))
				util.ResError(context, errors.ErrTooManyRequests)
				return
			}
		}

		context.Next()
	}
}
