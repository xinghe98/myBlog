package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type Myclaims struct {
	UserName string `json:"user_name"`
	Roles    string `json:"roles"`
	ID       uint   `json:"id"`
	jwt.RegisteredClaims
}

var Secret = []byte(viper.GetString("secret.key"))

// 生成token
func GenToken(username string, roles string, id uint) (string, error) {
	claims := &Myclaims{
		username,
		roles,
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 10)),
			Issuer:    "xinghe",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, *claims)
	return token.SignedString(Secret)
}

func ParseToken(tokenstr string) (*jwt.Token, *Myclaims, error) {
	claims := &Myclaims{}
	token, err := jwt.ParseWithClaims(tokenstr, claims, func(t *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	return token, claims, err
}
