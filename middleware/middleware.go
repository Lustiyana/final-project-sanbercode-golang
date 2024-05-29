package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"simple-social-media/helpers"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")

		if auth == "" {
			unauthorized(c, "Unauthorized: No credentials provided")
			c.Abort()
			return
		}

		token, err := extractToken(auth)
		if err != nil {
			unauthorized(c, "Unauthorized: Invalid credentials format")
			c.Abort()
			return
		}

		valid := checkCredentials(token)
		if !valid {
			unauthorized(c, "Unauthorized: Invalid credential")
			c.Abort()
			return
		}

		c.Next()
	}
}

func extractToken(auth string) (string, error) {
	token := strings.TrimPrefix(auth, "Bearer ")
	if token == auth {
		return "", http.ErrNotSupported
	}
	return token, nil
}

func checkCredentials(token string) bool {
	_, err := helpers.VerifyToken(token)
	if err != nil {
		return false
	}

	return true
}

func unauthorized(ctx *gin.Context, message string) {
	helpers.GeneralResponse(ctx, http.StatusUnauthorized, false, message, nil, nil)
}
