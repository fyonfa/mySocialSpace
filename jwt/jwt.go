package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/fyonfa/mySocialSpace/models"
	"time"
)

func GenerateJWT(t models.User) (string, error) {
	myKey := []byte("DevelopedMaster_FK_Auto")

	//privileges = Mapclaims
	//never in jwt save a password
	payload := jwt.MapClaims{
		"email":     t.Email,
		"name":      t.Name,
		"birthDate": t.BirthDate,
		"biography": t.Biography,
		"location":  t.Location,
		"website":   t.Website,
		"_id":       t.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
