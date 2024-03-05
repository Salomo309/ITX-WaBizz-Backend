package middlewares

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func VerifyTokenMiddleware() gin.HandlerFunc {
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
		token := strings.TrimPrefix(authHeader, "Bearer ")

		res, err := http.Get("https://www.googleapis.com/oauth2/v3/tokeninfo?id_token=" + token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": "Failed to get response"})
			return
		}
		defer res.Body.Close()

		// Decode Google response
		var response map[string]interface{}
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": "Failed to decode response"})
			return
		}

		// Check for error in Google response
		if _, ok := response["error_description"]; ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Error": "Token verification failed"})
		}

		// Go to next handler if token valid, if not then abort
		if _, ok := response["sub"]; ok {
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Error": "Invalid token"})
		}
	}
}
