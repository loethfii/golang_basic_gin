package auth

import (
	"errors"
	"github.com/golang-jwt/jwt"
	// "github.com/golang-jwt/jwt"
	"time"
)

// sebaiknya jangan hardcord taro di env atau secret
var jwtKey = []byte("rahasia")

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     uint   `json:"role"`
	jwt.StandardClaims
}

func GenereteJWT(email, username string, role uint) (string, error) {
	expTime := time.Now().Add(5 * time.Minute)

	claims := &JWTClaim{
		Username: username,
		Email:    email,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

func ValidateToken(signedToken string) (email string, role uint, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JWTClaim)

	//jika claim gagal
	if !ok {
		err = errors.New("Cloud parse claims for token")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("Token expired")
		return
	}

	role = claims.Role
	email = claims.Email

	return
}
