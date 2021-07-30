package service

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pluveto/bloggo/pkg/errcode"
)

func (s *Service) TokenIsDestroyed(token string) (destroyed bool) {
	return s.repo.TokenIsDestroyed(token)
}

// 创建一个token
func (s *Service) TokenCreate(claims jwt.StandardClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("some key here"))
}

func (s *Service) TokenCreateClaims(adminId uint64) *jwt.StandardClaims {
	var times time.Time = time.Now()

	return &jwt.StandardClaims{
		Audience:  "client",
		ExpiresAt: times.AddDate(0, 0, 1).Unix(),
		Id:        "grpcid",
		IssuedAt:  times.Unix(),
		Issuer:    "bloggo_server",
		NotBefore: times.Unix(),
		Subject:   strconv.FormatUint(adminId, 10),
	}
}

func (s *Service) TokenParse(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte("some key here"), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errcode.ApiError(errcode.ErrTokenMalformed)
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errcode.ApiError(errcode.ErrTokenExpired)
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errcode.ApiError(errcode.ErrTokenNotValidYet)
			} else {
				return nil, errcode.ApiError(errcode.ErrTokenInvalid)
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, errcode.ApiError(errcode.ErrTokenInvalid)

	} else {
		return nil, errcode.ApiError(errcode.ErrTokenInvalid)

	}

}
