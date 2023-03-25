package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Brolervi/Twitter_Chafa/db"
)

// ReadTweets, Gets a user's Tweets
func ReadTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Missing ID parameter", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Page parameter required", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Page parameter value must be higher than 0", http.StatusBadRequest)
		return
	}

	pag := int64(page)

	res, correct := db.GetTweets(ID, pag)
	if !correct {
		http.Error(w, "Error while getting tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
