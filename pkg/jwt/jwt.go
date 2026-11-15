package jwt

import (
	"errors"
	"time"

	gojwt "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID uint   `json:"uid"`
	Role   string `json:"role"`
	gojwt.RegisteredClaims
}

func Sign(secret string, uid uint, role string, ttl time.Duration) (string, error) {
	c := Claims{UserID: uid, Role: role, RegisteredClaims: gojwt.RegisteredClaims{
		ExpiresAt: gojwt.NewNumericDate(time.Now().Add(ttl)),
		IssuedAt:  gojwt.NewNumericDate(time.Now()),
	}}
	return gojwt.NewWithClaims(gojwt.SigningMethodHS256, c).SignedString([]byte(secret))
}

func Parse(secret, token string) (*Claims, error) {
	c := &Claims{}
	t, err := gojwt.ParseWithClaims(token, c, func(t *gojwt.Token) (any, error) {
		if _, ok := t.Method.(*gojwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil || !t.Valid {
		return nil, errors.New("invalid token")
	}
	return c, nil
}
