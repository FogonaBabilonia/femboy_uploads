package controllers

import (
	"fmt"
	"net/http"

	"github.com/FogonaBabilonia/femboy_uploads/auth"
	"github.com/gin-gonic/gin"
)

func HandleRootPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index", gin.H{
		"title": "Femboy Uploads",
		"msg":   "Welcome to femboy uploads",
	})
}

func HandleRegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register", gin.H{
		"title": "Femboy Uploads - Create Account",
	})
}

func HandleLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login", gin.H{
		"title": "Femboy Uploads - Login",
	})
}

func HandlePublicPage(c *gin.Context) {
	c.HTML(http.StatusOK, "public", gin.H{
		"title": "Femboy Uploads - public files",
	})
}

func HandleUserPage(c *gin.Context) {
	cookie, err := c.Cookie("jwt")

	if err == http.ErrNoCookie {
		c.String(http.StatusUnauthorized, "You are not logged")
		return
	}

	logged, username := auth.AuthWithToken(cookie)

	if !logged {
		c.String(http.StatusUnauthorized, "You are not logged")
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("You are logged as %v", username))
}
