package routers

import (
	json "encoding/json"
	"github.com/fyonfa/mySocialSpace/bd"
	"net/http"
	"strconv"
)

//ReadTweets reads tweets
func ReadTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Should send the parameter id", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "should send the parameter page", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page")) //Alphabetic to integer
	if err != nil {
		http.Error(w, "Should send page with value greater than 0", http.StatusBadRequest)
		return
	}
	pag := int64(page)
	answer, right := bd.ReadTweet(ID, pag)
	if right == false {
		http.Error(w, "Error when reading tweets", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(answer)
}
