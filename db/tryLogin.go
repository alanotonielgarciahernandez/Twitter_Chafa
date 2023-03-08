package db

import (
	"github.com/Brolervi/Twitter_Chafa/models"
	"golang.org/x/crypto/bcrypt"
)

// TryLogin, tries to log in with the user given
func TryLogin(email string, password string) (models.User, bool) {
	us, found, _ := CheckUserAlreadyExists(email)
	if !found {
		return us, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(us.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return us, false
	}
	return us, true
}
