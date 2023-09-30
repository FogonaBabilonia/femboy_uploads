package controllers

import (
	"fmt"
	"net/http"

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
	authenticated := c.GetBool("authenticated")

	if !authenticated {
		c.String(http.StatusUnauthorized, "You are not logged")
		return
	}

	username := c.GetString("username")
	c.String(http.StatusOK, fmt.Sprintf("You are logged as %v", username))
}
