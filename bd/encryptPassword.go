package bd

import "golang.org/x/crypto/bcrypt"

//EncryptPassword is the routine that allows us to encrypt passwords
func EncryptPassword(pass string) (string, error) {
	cost := 8 //quantity, alg, 2^cost... grater cost, better security will have the pass but longer, standard 6 for normal user, 8 for admin user, min 6
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass),cost)//slice of bytes
	return string(bytes), err
}

