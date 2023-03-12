package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	Username string `json:"username"`
	UserID   int64  `json:"user_id"`
	jwt.RegisteredClaims
}

const ExpiresDuration = time.Hour * 2
const RExpiresDuration = time.Hour * 24 * 30

var SECRET = []byte("gin-project")

func GenToken(userID int64, username string) (token, rToken string, err error) {
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		Username: username,
		UserID:   userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ExpiresDuration)),
			Issuer:    "cc",
		},
	}).SignedString(SECRET)

	if err != nil {
		return "", "", err
	}

	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(RExpiresDuration)),
	}).SignedString(SECRET)

	return
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	claims := new(CustomClaims)
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return SECRET, nil
	})

	if err != nil {
		return nil, err
	}

	if token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func ParseRefreshToken(token, refreshToken string) (newToken, newRefreshToken string, err error) {

	if _, err = jwt.Parse(refreshToken, func(t *jwt.Token) (interface{}, error) {
		return SECRET, nil
	}); err != nil {
		return
	}

	claims := new(CustomClaims)
	_, err = jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return SECRET, nil
	})

	if errors.Is(err, jwt.ErrTokenExpired) {
		return GenToken(claims.UserID, claims.Username)
	}
	return
}
