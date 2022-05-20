package Utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"time"
)

const (
	SECRETARY = "blockChain"
	MAXAGE    = 60 * 60 * 24 // 1天
)

type CustomClaims struct {
	UserId int64
	jwt.StandardClaims
}

func CreateToken(UserId int64) string {
	c := &CustomClaims{
		UserId: UserId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(MAXAGE) * time.Second).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, err := token.SignedString([]byte(SECRETARY))
	if err != nil {
		log.Panicln(err)
	}
	return AesEncryptCBC(tokenString)
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRETARY), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*CustomClaims)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("OutDate")
}
