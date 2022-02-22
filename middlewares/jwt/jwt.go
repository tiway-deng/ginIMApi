package jwt

import (
	"ginIMApi/packages/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = 200
		token := c.Query("token")
		if token == "" {
			authorization := c.GetHeader("Authorization")
			authStr := strings.Split(authorization, " ")
			token = authStr[1]
		}

		claims, err := utils.ParseToken(token)
		if token == "" {
			code = 401
		} else {
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = 402
				default:
					code = 403
				}
			}
		}

		if code != 200 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  "登录失效",
				"data": data,
			})

			c.Abort()
			return
		}

		c.Set("user_id", claims.UserId)

		c.Next()
	}
}
