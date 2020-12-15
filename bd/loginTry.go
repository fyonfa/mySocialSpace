package bd

import (
	"github.com/fyonfa/mySocialSpace/models"
	"golang.org/x/crypto/bcrypt"
)

//LoginTry check the login to the DB
func LoginTry(email string, password string) (models.User, bool) {
	usu, found, _ := CheckUserExists(email)
	if found == false {
		return usu, false
	}
	passwordBytes := []byte(password)
	passwordDB := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return usu, false
	}
	return usu, true
}
