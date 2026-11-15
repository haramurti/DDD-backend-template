package middleware

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

// RequireRole role GLOBAL dari JWT.
func RequireRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !slices.Contains(roles, c.GetString("role")) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"success": false, "message": "forbidden"})
			return
		}
		c.Next()
	}
}

// TODO: RequireContextRole(db, roles...) buat role PER-KONTEKS (event/project/cohort):
// cek tabel membership ke DB tiap request, set "ctx_role", 404 kalau bukan member.
