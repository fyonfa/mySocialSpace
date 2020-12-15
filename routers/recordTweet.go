package routers

import (
	"encoding/json"
	"github.com/fyonfa/mySocialSpace/bd"
	"github.com/fyonfa/mySocialSpace/models"
	"net/http"
	"time"
)

//RecordTweet allos to record the tweet in the database
func RecordTweet(w http.ResponseWriter, r *http.Request) {
	var message models.Tweet
	err := json.NewDecoder(r.Body).Decode(&message)

	register := models.RecordTweet{
		UserID:  IDUSer,
		Message: message.Message,
		Date:    time.Now(),
	}
	_, status, err := bd.InsertTweet(register)
	if err != nil {
		http.Error(w, "Error when trying to insert the register "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "Tweet cannot be inserted ", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
