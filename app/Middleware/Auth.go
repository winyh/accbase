package Middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		hmacSampleSecret := []byte("my_secret_key")
		tokenString := c.Request.Header.Get("token")

		fmt.Println(tokenString)


		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return hmacSampleSecret, nil
		})

		if err != nil{
			log.Print("请检查 token 参数！")
			c.JSON(http.StatusOK, gin.H{
				"status": false,
				"msg":    "token校验失败！",
			})
			c.Abort()
			return
		}


		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims["username"], claims["nbf"])

			c.JSON(200, gin.H{
				"username": claims["username"],
				"password": claims["password"],
			})

		} else {
			fmt.Println(err)
			c.JSON(200, gin.H{
				"message": "获取失败！",
			})
		}

		fmt.Println(token, err)

		// 请求前

		c.Next()

		// 请求后
		fmt.Println("wwww")

	}
}
