package infrastructure

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context){
		// Get the Authorization header
		auth := ctx.GetHeader("Authorization") //header
		if !strings.HasPrefix(auth, "Bearer "){ //starts with Bearer
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			ctx.Abort()
			return
		}
		// Parse the token
		token := strings.TrimPrefix(auth, "Bearer ")
		cl, err := ParseToken(token)
		if err != nil{
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			ctx.Abort()
			return
		}
		ctx.Set("username", cl.Username)
		ctx.Set("role", cl.Role)
		ctx.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context){
		if ctx.GetString("role") != "admin"{
			ctx.JSON(http.StatusForbidden, gin.H{"error": "admin only"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}