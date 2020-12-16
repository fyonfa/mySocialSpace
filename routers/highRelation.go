package routers

import (
	"github.com/fyonfa/mySocialSpace/bd"
	"github.com/fyonfa/mySocialSpace/models"
	"net/http"
)

//HighRelation makes the register of the relation between users
func HighRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "The parameter ID is mandatory ", http.StatusBadRequest)
		return
	}
	var t models.Relation
	t.UserID = IDUSer
	t.UserRelationID = ID

	status, err := bd.InsertRelation(t)
	if err != nil {
		http.Error(w, "Error occurred when trying to insert relation "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "Not able to insert relation "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated) //we can add here also the application json, it is optional
}
