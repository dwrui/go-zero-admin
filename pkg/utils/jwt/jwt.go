package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtConfig struct {
	AccessSecret string
	AccessExpire int64
}

var (
	ErrTokenExpired     = errors.New("token is expired")
	ErrTokenNotValidYet = errors.New("token not active yet")
	ErrTokenMalformed   = errors.New("that's not even a token")
	ErrTokenInvalid     = errors.New("couldn't handle this token")
)

type CustomClaims struct {
	Data map[string]interface{} `json:"data"`
	jwt.RegisteredClaims
}

// GenerateToken 生成token
func GenerateToken(config JwtConfig, data map[string]interface{}) (string, error) {
	now := time.Now().Unix()
	//expireTime := now.Add(time.Duration(config.AccessExpire) * time.Second)

	//claims := CustomClaims{
	//	Data: data,
	//	RegisteredClaims: jwt.RegisteredClaims{
	//		ExpiresAt: jwt.NewNumericDate(expireTime),
	//		IssuedAt:  jwt.NewNumericDate(now),
	//		NotBefore: jwt.NewNumericDate(now),
	//	},
	//}
	data["exp"] = now + config.AccessExpire
	data["iat"] = now
	claims := make(jwt.MapClaims)
	for k, v := range data {
		claims[k] = v
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(config.AccessSecret))
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//return token.SignedString([]byte(config.AccessSecret))
}

// ParseToken 解析token
func ParseToken(config JwtConfig, tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AccessSecret), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}
