package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessionID, err := ctx.Cookie("session_id")
		if err != nil || sessionID != "logged_in" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 1, "message": "Unauthorized"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
