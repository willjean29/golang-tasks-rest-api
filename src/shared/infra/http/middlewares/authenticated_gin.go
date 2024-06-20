package middlewares

import (
	error "app/src/shared/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthenticatedGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")
		println(cookie)
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
		// Agregar el token decodificado al contexto de Gin, si es necesario
		c.Set("tokenDecoded", tokenDecoded)
		c.Next() // Continuar con el siguiente middleware o controlador
	}
}
