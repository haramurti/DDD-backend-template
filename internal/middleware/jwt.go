package middleware

import (
	"net/http"
	"strings"

	"hackathon-be/pkg/jwt"

	"github.com/gin-gonic/gin"
)

// Auth validasi JWT, set "uid" & "role" ke context. Token WS boleh lewat ?token=.
func Auth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tok := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
		if tok == "" {
			tok = c.Query("token")
		}
		claims, err := jwt.Parse(secret, tok)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "message": "unauthorized"})
			return
		}
		c.Set("uid", claims.UserID)
		c.Set("role", claims.Role)
		c.Next()
	}
}
