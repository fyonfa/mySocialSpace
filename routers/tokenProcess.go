package routers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/fyonfa/mySocialSpace/bd"
	"github.com/fyonfa/mySocialSpace/models"
	"strings"
)

//Email is the email used in all the endpoints
var Email string

//IDUser is the ID given back from the model, that we sill use in the endpoints
var IDUSer string

func TokenProcess(tk string) (*models.Claim, bool, string, error) {
	myKey := []byte("DevelopedMaster_FK_Auto")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer") //1 vector 2 strings [0] bearer and [1] the rest of the token
	if len(splitToken) != 2 {
		return claims, false, "", errors.New("invalid format token")
	}

	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil {
		_, found, _ := bd.CheckUserExists(claims.Email)
		if found == true {
			Email = claims.Email
			IDUSer = claims.ID.Hex()
		}
		return claims, found, IDUSer, nil
	}
	if !tkn.Valid {
		return claims, false, "", errors.New("invalid token")
	}
	return claims, false, "", err
}
