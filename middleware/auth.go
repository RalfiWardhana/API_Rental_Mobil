package middleware

import (
	"context"
	"rental/repository/user"
	"strings"

	"github.com/gin-gonic/gin"
)

func WithAuthentication(userUc user.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer") {
			c.JSON(401, map[string]string{
				"message": "unauthorize",
			})
			c.Abort()
		}

		auth := strings.Split(authHeader, " ")
		userData, err := userUc.DecryptJWT(auth[1])
		if err != nil {
			c.JSON(401, map[string]string{
				"message": "unauthorize",
			})
			c.Abort()
		}
		ctxUserID := context.WithValue(c.Request.Context(), "user_id", userData["id"])
		c.Request = c.Request.WithContext(ctxUserID)
		c.Next()
	}
}
