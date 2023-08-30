package controllers

import (
	"fmt"
	"net/http"

	"github.com/FogonaBabilonia/femboy_uploads/database"
	"github.com/FogonaBabilonia/femboy_uploads/models"
	"github.com/FogonaBabilonia/femboy_uploads/services"
	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file_upload")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": fmt.Errorf("error: %s", err.Error()),
		})
		return
	}

	filename := services.Format_filename(file.Filename)
	if err := c.SaveUploadedFile(file, "./uploads/"+filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": fmt.Errorf("error: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"msg": fmt.Sprintf("created file: %v", filename),
	})
}

func Create_User(c *gin.Context) {
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

	c.JSON(http.StatusCreated, gin.H{
		"created_user": db_user,
	})
}
