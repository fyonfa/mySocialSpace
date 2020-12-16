package routers

import (
	"github.com/fyonfa/mySocialSpace/bd"
	"net/http"
)

func RemoveTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Should send the ID parameter", http.StatusBadRequest)
		return
	}
	err := bd.EraseTweet(ID, IDUSer)
	if err != nil {
		http.Error(w, "Error when trying to erase the tweet"+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
