package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/fyonfa/mySocialSpace/bd"
	"github.com/fyonfa/mySocialSpace/models"
)

//UploadBanner uploads the avatar to the server
func UploadBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("banner")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archive = "uploads/banners/" + IDUSer + "." + extension
	f, err := os.OpenFile(archive, os.O_WRONLY|os.O_CREATE, 0666) //read and execution permits to this file
	if err != nil {
		http.Error(w, "Error when uploading image"+err.Error(), http.StatusBadRequest)
		return
	}
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error when copying image"+err.Error(), http.StatusBadRequest)
		return
	}
	var user models.User
	var status bool

	user.Banner = IDUSer + "." + extension
	status, err = bd.ModifyRegister(user, IDUSer)
	if err != nil || status == false {
		http.Error(w, "Error when recording banner in the DB! "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
