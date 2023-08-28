package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRoot(c *gin.Context) {
	c.HTML(http.StatusOK, "index", gin.H{
		"title": "Femboy Uploads",
		"msg":   "Welcome to femboy uploads",
	})
}

func HandleRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register", gin.H{
		"title": "Femboy Uploads - Create Account",
	})
}

func HandleLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login", gin.H{
		"title": "Femboy Uploads - Login",
	})
}

func HandlePublic(c *gin.Context) {
	c.HTML(http.StatusOK, "public", gin.H{
		"title": "Femboy Uploads - public files",
	})
}
