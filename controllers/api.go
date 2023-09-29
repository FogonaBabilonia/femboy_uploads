package controllers

import (
	"fmt"
	"net/http"

	"github.com/FogonaBabilonia/femboy_uploads/auth"
	"github.com/FogonaBabilonia/femboy_uploads/database"
	"github.com/FogonaBabilonia/femboy_uploads/models"
	"github.com/FogonaBabilonia/femboy_uploads/services"
	"github.com/gin-gonic/gin"
)

func HandleUploadFile(c *gin.Context) {
	file, err := c.FormFile("file_upload")

	if err != nil {
		c.String(http.StatusBadRequest, "Error, bad request")
		return
	}

	filename := services.Format_filename(file.Filename)
	if err := c.SaveUploadedFile(file, "./uploads/"+filename); err != nil {
		c.String(http.StatusInternalServerError, "Server error, could not save file")
		return
	}

	c.String(http.StatusCreated, "created file: %v", filename)
}

func HandleCreate_User(c *gin.Context) {
	user := new(models.Create_User)

	if err := c.Bind(user); err != nil {
		fmt.Println(err)
		c.String(http.StatusBadRequest, "could not create the user")
		return
	}

	exists := new(models.User)
	if err := database.DB.Where("name = ?", user.Name).First(exists).Error; err == nil {
		c.String(http.StatusBadRequest, "username already used")
		return
	}

	db_user, err := models.NewUser(user.Name, user.Password)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := database.DB.Create(db_user).Error; err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	generate_jwt_and_redirect(c, db_user.ID)
}

func HandleLogin(c *gin.Context) {
	user := new(models.SignIn_User)

	if err := c.ShouldBind(user); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	authenticated, id := auth.ValidatePassword(user)
	if !authenticated {
		c.String(http.StatusBadRequest, "wrong username or password")
		return
	}

	generate_jwt_and_redirect(c, id)
}

func generate_jwt_and_redirect(c *gin.Context, id uint) {
	jwt_token, err := auth.GenerateToken(id)

	if err != nil {
		c.String(http.StatusInternalServerError, "could not login the user")
		return
	}

	c.SetCookie("jwt", jwt_token, 600, "/", "localhost", false, false)
	c.Redirect(301, "/user")
}
