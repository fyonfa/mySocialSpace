package routers

import (
	"encoding/json"
	"github.com/fyonfa/mySocialSpace/bd"
	"github.com/fyonfa/mySocialSpace/models"
	"net/http"
)

//QueryRelation checks if the relation between 2 users exists
func QueryRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relation
	t.UserID = IDUSer
	t.UserRelationID = ID

	var answ models.AnswerConsultRelation

	status, err := bd.ConsultRelation(t) //the important here is to know if there is relation or not, only true or false
	if err != nil || status == false {
		answ.Status = false
	} else {
		answ.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(answ)

}
