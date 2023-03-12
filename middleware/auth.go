package middleware

import (
	"gin-project/constants"
	"gin-project/pkg/jwt"
	"gin-project/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")

		if authHeader == "" {
			utils.ResponseError(c, constants.CodeEmptyToken)
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.ResponseError(c, constants.CodeInvalidToken)
			c.Abort()
			return
		}

		mc, err := jwt.ParseToken(parts[1])

		if err != nil {
			utils.ResponseError(c, constants.CodeInvalidToken)
			c.Abort()
			return
		}

		c.Set("username", mc.Username)
		c.Next()
	}
}
