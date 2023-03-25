package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Brolervi/Twitter_Chafa/db"
)

// ViewProfile, Gets a user ID to extract his data
func ViewProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Missing ID parameter", http.StatusBadRequest)
		return
	}

	profile, err := db.SearchUser(ID)
	if err != nil {
		http.Error(w, "Error while searching user "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
