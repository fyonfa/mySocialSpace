package routers

import (
	"encoding/json"
	"github.com/fyonfa/mySocialSpace/bd"
	"github.com/fyonfa/mySocialSpace/models"
	"net/http"
)

//Register is the function to create the DB the register of the users
func Register(w http.ResponseWriter, r *http.Request) {

	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t) //Body of an HTTP is an string, once read, its destroyed, careful when using in other parts of the program

	if err != nil {
		http.Error(w, "Error in the received data "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "Email of user is required ", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Password least than 6 characters ", 400)
		return
	}
	_, found, _ := bd.CheckUserExists(t.Email) //I can't load the same information every time,
	if found == true {
		http.Error(w, "User already exist with this email ", 400)
		return
	}
	_, status, err := bd.InsertRegister(t)
	if err != nil {
		http.Error(w, "Error when tried to register user "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "Can't insert the user register ", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
