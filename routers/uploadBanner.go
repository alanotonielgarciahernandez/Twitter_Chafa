package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Brolervi/Twitter_Chafa/db"
	"github.com/Brolervi/Twitter_Chafa/models"
)

// Upload Banner, Gets a banner image an uploads it to the user's profile in the DB
func UploadBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("banner")
	if err != nil {
		http.Error(w, "Error while searching banner image "+err.Error(), http.StatusBadRequest)
		return
	}
	var extension = strings.Split(handler.Filename, ".")[1]
	var dir string = "uploads/banners/" + IDUser + "." + extension

	f, err := os.OpenFile(dir, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error while uploading the image "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error while copying the image "+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Banner = IDUser + "." + extension
	status, err = db.ModifyUser(user, IDUser)
	if err != nil || !status {
		http.Error(w, "Error while saving banner to DB "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
