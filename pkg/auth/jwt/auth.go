package jwt

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"goApi/pkg/auth"
	"time"
)

var defaultKey = "go-api"

var defaultOptions = options{
	signingMethod: jwt.SigningMethodHS512,
	signingKey:    []byte(defaultKey),
	tokenType:     "Bearer",
	expired:       7200,
	keyfunc: func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return nil, auth.ErrInvalidToken
		}
		return []byte(defaultKey), nil
	},
}

type options struct {
	signingMethod jwt.SigningMethod
	signingKey    any
	tokenType     string
	expired       int
	keyfunc       jwt.Keyfunc
}

type JWTAuth struct {
	store Store
	opts  *options
}

type Option func(*options)

func SetSigningMethod(method jwt.SigningMethod) Option {
	return func(o *options) {
		o.signingMethod = method
	}
}

func SetSigningKey(key interface{}) Option {
	return func(o *options) {
		o.signingKey = key
	}
}

func SetKeyfunc(keyFunc jwt.Keyfunc) Option {
	return func(o *options) {
		o.keyfunc = keyFunc
	}
}

func SetExpired(expired int) Option {
	return func(o *options) {
		o.expired = expired
	}
}

func New(store Store, opts ...Option) *JWTAuth {
	o := defaultOptions
	for _, opt := range opts {
		opt(&o)
	}

	return &JWTAuth{
		store: store,
		opts:  &o,
	}
}

func (a *JWTAuth) GenerateToken(ctx context.Context, userID string) (auth.TokenInfo, error) {

	now := time.Now()
	expiresAt := now.Add(time.Duration(a.opts.expired) * time.Second)
	token := jwt.NewWithClaims(a.opts.signingMethod, &jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: &jwt.NumericDate{Time: expiresAt},
		NotBefore: &jwt.NumericDate{Time: now},
		IssuedAt:  &jwt.NumericDate{Time: now},
	})
	tokenString, err := token.SignedString([]byte(a.opts.signingKey.(string)))
	if err != nil {
		return nil, err
	}

	return &TokenInfo{
		AccessToken: tokenString,
		TokenType:   a.opts.tokenType,
		ExpiresAt:   expiresAt.Unix(),
	}, nil
}

func (a *JWTAuth) parseToken(accessToken string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &jwt.RegisteredClaims{}, a.opts.keyfunc)
	if err != nil || !token.Valid {
		return nil, auth.ErrInvalidToken
	}
	return token.Claims.(*jwt.RegisteredClaims), nil
}

func (a *JWTAuth) callStore(fn func(Store) error) error {
	if store := a.store; store != nil {
		return fn(store)
	}
	return nil
}

func (a *JWTAuth) DestroyToken(ctx context.Context, accessToken string) error {
	claim, err := a.parseToken(accessToken)
	if err != nil {
		return err
	}

	return a.callStore(func(store Store) error {
		expiration, _ := claim.GetExpirationTime()
		expirationTime := expiration.Time.Unix()
		expired := time.Unix(expirationTime, 0).Sub(time.Now())
		return store.Set(ctx, accessToken, expired)
	})
}

func (a *JWTAuth) ParseUserID(ctx context.Context, accessToken string) (string, error) {
	if accessToken == "" {
		return "", auth.ErrInvalidToken
	}

	claims, err := a.parseToken(accessToken)
	if err != nil {
		return "", err
	}

	err = a.callStore(func(store Store) error {
		if exists, err := store.Check(ctx, accessToken); err != nil {
			return err
		} else if exists {
			return auth.ErrInvalidToken
		}
		return nil
	})

	if err != nil {
		return "", err
	}
	return claims.GetSubject()
}

func (a *JWTAuth) Release() error {
	return a.callStore(func(store Store) error {
		return store.Release()
	})
}
