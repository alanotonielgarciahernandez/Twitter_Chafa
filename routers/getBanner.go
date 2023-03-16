package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/Brolervi/Twitter_Chafa/db"
)

// GetAvatar, Sends banner to HTTP
func GetBanner(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Need to send ID", http.StatusBadRequest)
		return
	}

	profile, err := db.SearchProfile(ID)
	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	openFile, err := os.Open("uploads/banner/" + profile.Banner)
	if err != nil {
		http.Error(w, "Image not found "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, openFile)
	if err != nil {
		http.Error(w, "Error while copying image "+err.Error(), http.StatusBadRequest)
	}
}
