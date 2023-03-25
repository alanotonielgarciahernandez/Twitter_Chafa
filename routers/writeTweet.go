package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Brolervi/Twitter_Chafa/db"
	"github.com/Brolervi/Twitter_Chafa/models"
)

// WriteTweet, Writes a Tweet
func WriteTweet(w http.ResponseWriter, r *http.Request) {
	var message models.Tweet

	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, "Error in data received "+err.Error(), http.StatusBadRequest)
		return
	}

	reg := models.RecordTweet{
		UserID:  IDUser,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := db.InsertTweet(reg)
	if err != nil {
		http.Error(w, "Error while trying to insert the Tweet, try again "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "Can't insert the Tweet", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
