package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/Brolervi/Twitter_Chafa/db"
)

// GetAvatar, Sends user's banner image to HTTP
func GetBanner(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Missing ID parameter", http.StatusBadRequest)
		return
	}

	profile, err := db.SearchUser(ID)
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
		http.Error(w, "Error while copying the image "+err.Error(), http.StatusBadRequest)
	}
}
