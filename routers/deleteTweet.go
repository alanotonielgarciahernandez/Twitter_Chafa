package routers

import (
	"net/http"

	"github.com/Brolervi/Twitter_Chafa/db"
)

// DeleteTweet, Gets the ID of the Tweet to delete
func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Missing ID parameter", http.StatusBadRequest)
		return
	}

	err := db.DeleteTweet(ID, IDUser)
	if err != nil {
		http.Error(w, "Error while trying to delete the Tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
