package setup

import (
	"github.com/golang-jwt/jwt/v5"
	"goApi/internal/app/config"
	"goApi/pkg/auth"
	jwtauth "goApi/pkg/auth/jwt"
	"goApi/pkg/auth/jwt/store/buntdb"
	"goApi/pkg/auth/jwt/store/redis"
)

func Auth() (auth.Author, func(), error) {
	jwtCfg := config.C.JWTAuth

	var opts []jwtauth.Option
	var method jwt.SigningMethod
	opts = append(opts, jwtauth.SetExpired(jwtCfg.Expired))
	switch jwtCfg.SigningMethod {
	case "HS256":
		method = jwt.SigningMethodHS256
	case "HS384":
		method = jwt.SigningMethodHS384
	default:
		method = jwt.SigningMethodHS512
	}
	opts = append(opts, jwtauth.SetSigningMethod(method))
	opts = append(opts, jwtauth.SetSigningKey(jwtCfg.SigningKey))
	opts = append(opts, jwtauth.SetKeyfunc(func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, auth.ErrInvalidToken
		}
		return []byte(jwtCfg.SigningKey), nil
	}))

	var store jwtauth.Store
	switch jwtCfg.Store {
	case "redis":
		redisCfg := config.C.Redis
		store = redis.NewStore(&redis.Config{
			Addr:     redisCfg.Addr,
			DB:       jwtCfg.RedisDB,
			Password: redisCfg.Password,
			Prefix:   jwtCfg.RedisPrefix,
		})
	default:
		s, err := buntdb.NewStore(jwtCfg.FilePath)
		if err != nil {
			return nil, nil, err
		}
		store = s
	}
	jwtAuth := jwtauth.New(store, opts...)
	return jwtAuth, func() {
		store.Release()
	}, nil
}
