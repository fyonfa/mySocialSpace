package routers

import (
	"encoding/json"
	"github.com/fyonfa/mySocialSpace/bd"
	"net/http"
	"strconv"
)

//UserList reads the list of all the users
func UserList(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "page parameter should be greater than 0 ", http.StatusBadRequest)
		return
	}
	pag := int64(pagTemp)

	result, status := bd.ReadAllUsers(IDUSer, pag, search, typeUser)
	if status == false {
		http.Error(w, "Error when reading users ", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(result)

}
