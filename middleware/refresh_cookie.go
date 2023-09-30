package middleware

import (
	"net/http"

	"github.com/FogonaBabilonia/femboy_uploads/auth"
	"github.com/gin-gonic/gin"
)

func RefreshCookie() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("jwt")

		if err == nil && cookie != "" {
			valid, username := auth.AuthWithToken(cookie)

			if valid {
				jwt_token, err := auth.GenerateToken(username)

				if err != nil {
					c.String(http.StatusInternalServerError, "server error")
					return
				}

				c.SetCookie("jwt", jwt_token, 300, "/", "localhost", false, false)
			}
		}

		c.Next()
	}
}
