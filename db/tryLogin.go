package db

import (
	"github.com/Brolervi/Twitter_Chafa/models"
	"golang.org/x/crypto/bcrypt"
)

// TryLogin, Attempts to log the given user in
func TryLogin(email string, password string) (models.User, bool) {
	us, found, _ := CheckUserAlreadyExists(email)
	if !found {
		return us, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(us.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return us, false
	}
	return us, true
}
