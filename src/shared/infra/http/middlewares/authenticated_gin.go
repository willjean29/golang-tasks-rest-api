package middlewares

import (
	error "app/src/shared/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AuthenticatedGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, error.New("Unauthorized", http.StatusUnauthorized, nil))
			return
		}
		token := cookie
		isValidToken, tokenDecoded, err := tokenAdapter.ValidateToken(token)
		if !isValidToken {
			c.AbortWithStatusJSON(http.StatusUnauthorized, error.New("Unauthorized (Token invalid)", http.StatusUnauthorized, err))
			return
		}

		id := tokenDecoded["userId"].(string)
		userId, err := strconv.Atoi(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, error.New("Unauthorized (Error decoded token)", http.StatusUnauthorized, err))
			return

		}

		c.Set(string(userKey), userId)
		c.Next()
	}
}
