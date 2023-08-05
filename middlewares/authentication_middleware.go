package middlewares

import (
	"github.com/ekantbajaj/duty-allocation/token"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get token from header
		// Verify token
		// If verified, set user in context
		// If not verified, abort with error
		if ctx.Request.URL.Path != "/users/login" {
			authorizationHeader := ctx.GetHeader("Authorization")
			if authorizationHeader == "" {
				ctx.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
				return
			}
			// if token is provided then it should be of for bearer token
			fields := strings.Fields(authorizationHeader)
			if len(fields) != 2 || strings.ToLower(fields[0]) != "bearer" {
				ctx.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
				return
			}
			accessToken := fields[1]
			// Verify token
			payload, err := tokenMaker.VerifyToken(accessToken)
			if err != nil {
				ctx.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
				return
			}
			// Set payload in context
			ctx.Set("authorization_payload", payload)
		}
		ctx.Next()
	}
}
