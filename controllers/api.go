package controllers

import (
	"fmt"
	"net/http"

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

}
