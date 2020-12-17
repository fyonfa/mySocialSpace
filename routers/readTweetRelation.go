package routers

import (
	"encoding/json"
	"github.com/fyonfa/mySocialSpace/bd"
	"net/http"
	"strconv"
)
//ReadTweetFollowers read the tweets of all our followers
func ReadTweetFollowers(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Page number should be sent ", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page")) //page to integer
	if err != nil {
		http.Error(w, "Page parameter should be greater than 0 ", http.StatusBadRequest)
		return
	}
	answer, right := bd.ReadTweetFollowers(IDUSer, page)
	if right == false {
		http.Error(w, "Error when reading tweets ", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(answer)

}
