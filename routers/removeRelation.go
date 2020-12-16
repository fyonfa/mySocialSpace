package routers

import (
	"github.com/fyonfa/mySocialSpace/bd"
	"github.com/fyonfa/mySocialSpace/models"
	"net/http"
)

//RemoveRelation removes the relation between users
func RemoveRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Relation
	t.UserID = IDUSer
	t.UserRelationID = ID

	status, err := bd.EraseRelation(t)
	if err != nil {
		http.Error(w, "Error occurred when trying to remove relation "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "Not able to remove relation "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated) //we can add here also the application json, it is optional
}
