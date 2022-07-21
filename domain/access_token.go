package domain

import (
	"os"
	"time"

	"github.com/bshadmehr-com/libs/errs"
	"github.com/bshadmehr-com/libs/logger"
	"github.com/golang-jwt/jwt"
)

type AccessToken struct {
	AccessToken  string
	RefreshToken string
	UserId       int64
	ExpiresAt    int64
}

//go:generate mockgen -destination=../mocks/domain/mockAccessTokenRepository.go -package=domain github.com/bshadmehr-com/auth/domain AccessTokenRepository
type AccessTokenRepository interface {
	IsAuthorized(token AccessToken, route string, vars map[string]string) (bool, *jwt.MapClaims)
	Logout(token AccessToken) *errs.AppError
}

func (at AccessToken) GetExpiresAt() int64 {
	return time.Now().Add(time.Minute * 10).Unix()
}

func (at AccessToken) GetefreshTokenExpiresAt() int64 {
	return time.Now().Add(time.Minute * 10).Unix()
}

func GetNewAccessToken(claims jwt.MapClaims) (*AccessToken, *errs.AppError) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedTokenAsString, err := token.SignedString([]byte(os.Getenv("HMAC_SECRET")))
	if err != nil {
		err := errs.NewUnexpectedError("Error while creating access token")
		return nil, err
	}
	return &AccessToken{
		AccessToken:  signedTokenAsString,
		RefreshToken: "",
	}, nil
}

func GetNewAccessTokenFromRefreshClaims(claims jwt.MapClaims) (*AccessToken, *errs.AppError) {
	delete(claims, "type")
	claims["exp"] = time.Now().Add(TOKEN_DURATION_MINUTES * time.Minute).Unix()
	return GetNewAccessToken(claims)
}

func GetNewRefreshToken(claims jwt.MapClaims) (*AccessToken, *errs.AppError) {
	claims["type"] = "refresh"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedTokenAsString, err := token.SignedString([]byte(os.Getenv("HMAC_SECRET")))
	if err != nil {
		err := errs.NewUnexpectedError("Error while creating access token")
		return nil, err
	}
	return &AccessToken{
		AccessToken:  "",
		RefreshToken: signedTokenAsString,
	}, nil
}

func JwtTokenFromString(tokenString string) (*jwt.Token, *errs.AppError) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("HMAC_SECRET")), nil
	})

	if err != nil {
		logger.Error(err.Error())
		return nil, errs.NewUnexpectedError("Erro whie trying to decode the token")
	}

	return token, nil
}
