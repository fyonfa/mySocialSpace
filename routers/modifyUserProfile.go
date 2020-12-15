package routers

import (
	"encoding/json"
	"github.com/fyonfa/mySocialSpace/bd"
	"github.com/fyonfa/mySocialSpace/models"
	"net/http"
)

//ModifyUserProfile modifies the user profile
func ModifyUserProfile(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Incorrect data "+err.Error(), 400)
		return
	}
	var status bool
	status, err = bd.ModifyRegister(t, IDUSer)
	if err != nil {
		http.Error(w, "Error when trying to modify a register, try again "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "Not able to modify the user register", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
