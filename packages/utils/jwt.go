package utils

import (
	"ginIMApi/models"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret []byte

type Claims struct {
	UserId   string `json:"user_id"`
	Mobile   string `json:"mobile"`
	Nickname string `json:"nickname"`
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(user models.User) (string, int64, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour).Unix()

	claims := Claims{
		strconv.Itoa(user.ID),
		user.Mobile,
		user.Nickname,
		jwt.StandardClaims{
			ExpiresAt: expireTime,
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, expireTime, err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
