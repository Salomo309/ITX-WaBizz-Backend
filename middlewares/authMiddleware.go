package middlewares

import (
	"itx-wabizz/repositories"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

/*
Function: VerifyEmailMiddleware

Provide a function for router to protect vital endpoints
*/
func VerifyEmailMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if request sent with Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Error": "Authorization header missing"})
			return
		}

		// Check if Authorization header contain "Bearer " prefix
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Error": "Invalid authorization header"})
			return
		}

		// Get authorization token and ask its validity with Google API
		email := strings.TrimPrefix(authHeader, "Bearer ")
		existingUser, err := repositories.UserRepo.GetUser(email)

		// Go to next handler if token valid, if not then abort
		if err == nil && existingUser.IsActive {
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized"})
			return
		}
	}
}
