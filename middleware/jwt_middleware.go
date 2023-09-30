package middleware

import (
	"net/http"

	"github.com/FogonaBabilonia/femboy_uploads/auth"
	"github.com/gin-gonic/gin"
)

func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("jwt")

		if err == http.ErrNoCookie {
			c.Set("authenticated", false)
		} else {
			authenticated, username := auth.AuthWithToken(cookie)
			c.Set("username", username)
			c.Set("authenticated", authenticated)
		}

		c.Next()
	}
}
