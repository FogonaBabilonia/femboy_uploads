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

	if err := c.ShouldBind(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "could not create the user",
		})
		return
	}

	db_user, err := models.NewUser(user.Name, user.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if err := database.DB.Create(db_user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	generate_jwt_and_redirect(c, db_user)
}

func HandleLogin(c *gin.Context) {
	user := new(models.SignIn_User)

	if err := c.ShouldBind(user); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	authenticated := auth.ValidatePassword(user)
	if !authenticated {
		c.String(http.StatusBadRequest, "wrong username or password")
		return
	}

	db_user, err := models.NewUser(user.Name, user.Password)

	if err != nil {
		c.String(http.StatusInternalServerError, "could not create user due to server error")
		return
	}

	generate_jwt_and_redirect(c, db_user)
}

func generate_jwt_and_redirect(c *gin.Context, db_user *models.User) {
	jwt_token, err := auth.GenerateToken(db_user.ID)
	fmt.Println("generating token for user", db_user.ID, db_user.Name)

	if err != nil {
		c.String(http.StatusInternalServerError, "could not login the user")
		return
	}

	c.SetCookie("jwt", jwt_token, 600, "/", "localhost", false, false)
	//c.Redirect(http.StatusPermanentRedirect, "/user")
	c.String(http.StatusOK, "Ok")
}
