package utils

import (
	"time"

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
