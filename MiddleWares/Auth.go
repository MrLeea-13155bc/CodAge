package MiddleWares

import (
	"QuestionSearch/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"sync"
	"time"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err == nil && token != "" {
			token = strings.ReplaceAll(token, " ", "+") // !!!!!! token里面的+会变成空格!!!!!!!!!!!
			data, err := Utils.AesDecryptCBC(token)
			if err == nil {
				info, err := Utils.ParseToken(data)
				if err == nil {
					c.Set("userId", info.UserId)
					wg := sync.WaitGroup{}
					//创建协程cookie续杯
					go func(info *Utils.CustomClaims) {
						wg.Add(1)
						defer wg.Done()
						t := time.Unix(info.ExpiresAt, 0)
						timeExceed := int(t.Sub(time.Now()).Seconds())
						if timeExceed < Utils.MAXAGE/2 {
							newToken := Utils.CreateToken(info.UserId)
							http.SetCookie(c.Writer, &http.Cookie{
								Name:     "token",
								Value:    newToken,
								Path:     "/",
								Domain:   "",
								MaxAge:   Utils.MAXAGE,
								Secure:   true,
								SameSite: 4,
							})
						}
					}(info)
					c.Next()
					wg.Wait()
					return
				}
			}
		}
		c.JSON(http.StatusUnauthorized, nil)
		c.Abort()
	}
}
