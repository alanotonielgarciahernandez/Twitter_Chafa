package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Brolervi/Twitter_Chafa/db"
)

// ReadTweets, read the Tweets
func ReadTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Need to send ID", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Need to send page", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Need to send page with value higher than 0", http.StatusBadRequest)
		return
	}

	pag := int64(page)

	res, correct := db.ReadTweet(ID, pag)
	if !correct {
		http.Error(w, "Error while reading tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
