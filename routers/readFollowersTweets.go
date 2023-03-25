package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Brolervi/Twitter_Chafa/db"
)

// ReadFollowersTweets, Requests all Tweets from the given user's followers
func ReadFollowersTweets(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Page parameter required", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Page parameter must be higher than 0", http.StatusBadRequest)
		return
	}

	res, status := db.ReadFollowersTweets(IDUser, page)
	if !status {
		http.Error(w, "Error while getting Tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
