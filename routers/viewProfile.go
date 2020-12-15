package routers

import (
	"encoding/json"
	"github.com/fyonfa/mySocialSpace/bd"
	"net/http"
)

func ViewProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "ID parameter should be sent", http.StatusBadRequest)
		return
	}
	profile, err := bd.SearchProfile(ID)
	if err != nil {
		http.Error(w, "Error happened when trying to find the register"+err.Error(), 400)
		return
	}

	//in case tha profile found
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(profile)
}
