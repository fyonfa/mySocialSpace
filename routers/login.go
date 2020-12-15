package routers

import (
	"encoding/json"
	"github.com/fyonfa/mySocialSpace/bd"
	"github.com/fyonfa/mySocialSpace/jwt"
	"github.com/fyonfa/mySocialSpace/models"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request)  {
	w.Header().Add("content-type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w,"User or password wrong "+err.Error(),400)
		return
	}
	if len(t.Email)==0{
		http.Error(w, "Email is required ", 400)
		return
	}
	document , exist := bd.LoginTry(t.Email,t.Password)
	if exist == false{
		http.Error(w, "User or password wrong ", 400)
		return
	}
	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "Error when tried to generate token "+err.Error(), 400)
		return
	}
	resp := models.LoginResponse{
		Token: jwtKey,
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(resp)

	//Set a cookie in the user side
	expirationTime := time.Now().Add(24*time.Hour)
	http.SetCookie(w,&http.Cookie{
		Name: "token",
		Value: jwtKey,
		Expires: expirationTime,
	})

}