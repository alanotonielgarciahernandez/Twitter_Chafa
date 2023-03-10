package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Brolervi/Twitter_Chafa/db"
	"github.com/Brolervi/Twitter_Chafa/models"
)

// RecordTweet, allow record the Tweet in the db
func RecordTweet(w http.ResponseWriter, r *http.Request) {
	var message models.Tweet

	err := json.NewDecoder(r.Body).Decode(&message)

	reg := models.RecordTweet{
		UserID:  IDUser,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := db.InsertTweet(reg)
	if err != nil {
		http.Error(w, "Error while trying insert the log, try again "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "Can't insert the Tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
