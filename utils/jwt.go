package utils

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ritesh-15/notesync-backend/config"
)

type JWTCustomClaim struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

func SignVerificationToken(id string) (string, error) {

	verificationTokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTCustomClaim{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 10)),
		},
	})

	verificationToken, err := verificationTokenClaim.SignedString([]byte(config.VERIFICATION_TOKEN_SECRET))

	return verificationToken, err
}

func GenerateAccessAndRefreshToken(
	id string,
) (string, string) {

	var accessToken string
	var refreshToken string
	var err error

	accessTokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTCustomClaim{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 10)),
		},
	})

	accessToken, err = accessTokenClaim.SignedString([]byte(config.ACCESS_TOKEN_SECRET))

	if err != nil {
		log.Error(err)
	}

	refreshTokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTCustomClaim{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
		},
	})

	refreshToken, err = refreshTokenClaim.SignedString([]byte(config.REFRESH_TOKEN_SECRET))

	if err != nil {
		log.Error(err)
	}

	return accessToken, refreshToken
}

func VerifyToken(receviedToken string, secret string) (*JWTCustomClaim, error) {
	token, err := jwt.ParseWithClaims(receviedToken, &JWTCustomClaim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTCustomClaim); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
