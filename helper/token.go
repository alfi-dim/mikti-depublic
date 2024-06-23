package helper

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type TokenUseCase interface {
	GenerateAccessToken(claims JwtCustomClaims) (string, error)
	DecodeTokenPayload(token string) (JwtCustomClaims, error)
}

type TokenUseCaseImpl struct{}

type JwtCustomClaims struct {
	ID    string `json:"user_id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func NewTokenUseCase() *TokenUseCaseImpl {
	return &TokenUseCaseImpl{}
}

func (t *TokenUseCaseImpl) GenerateAccessToken(claims JwtCustomClaims) (string, error) {
	plainToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	encodedToken, err := plainToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return encodedToken, nil

}

func (t *TokenUseCaseImpl) DecodeTokenPayload(token string) (JwtCustomClaims, error) {
	payload, err := jwt.ParseWithClaims(token, &JwtCustomClaims{}, func(jwtToken *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return JwtCustomClaims{}, err
	}

	claims, ok := payload.Claims.(*JwtCustomClaims)
	if ok && payload.Valid {
		return *claims, nil
	}
	return JwtCustomClaims{}, errors.New("failed when decode payload")
}
