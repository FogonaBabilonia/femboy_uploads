package auth

import (
	"fmt"

	"github.com/FogonaBabilonia/femboy_uploads/database"
	"github.com/FogonaBabilonia/femboy_uploads/models"
	"golang.org/x/crypto/bcrypt"
)

func ValidatePassword(user *models.SignIn_User) (bool, uint) {
	db_user_data := new(models.User)

	if err := database.DB.Where("name = ?", user.Name).First(db_user_data).Error; err != nil {
		fmt.Println(err)
		return false, 0
	}

	hashed_pass := []byte(db_user_data.Password)
	password := []byte(user.Password)

	if err := bcrypt.CompareHashAndPassword(hashed_pass, password); err != nil {
		return false, 0
	}

	return true, db_user_data.ID

}
