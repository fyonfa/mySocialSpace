package routers

import (
	"github.com/fyonfa/mySocialSpace/bd"
	"io"
	"net/http"
	"os"
)

//GetBanner send the banner to the http
func GetBanner(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "ID parameter should be sent", http.StatusBadRequest)
		return
	}
	profile, err := bd.SearchProfile(ID)
	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	openFile, err := os.Open("uploads/banners/" + profile.Banner)
	if err != nil {
		http.Error(w, "Image not found", http.StatusBadRequest)
		return
	}
	_, err = io.Copy(w, openFile) //sends to w the file in binary mode that has been opened
	if err != nil {
		http.Error(w, "Error when copying the image", http.StatusBadRequest)
	}
}
