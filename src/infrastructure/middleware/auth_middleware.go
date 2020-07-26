package middleware

import (
	"CleanArchitectureSample_golang/common"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// RequireAuth detects if signed in.
func RequireAuth(ctx *gin.Context) {
	session := sessions.Default(ctx)
	user := session.Get(common.SessionUserIDKey)
	if user == nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "Unauthorized", "message": "unauthorized"})
		return
	}
	ctx.Next()
}
