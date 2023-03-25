package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Brolervi/Twitter_Chafa/db"
	"github.com/Brolervi/Twitter_Chafa/models"
)

// ModifyProfile, Gets the information to be modified in the user profile
func ModifyProfile(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error in data received "+err.Error(), http.StatusBadRequest)
		return
	}

	var status bool

	status, err = db.ModifyUser(t, IDUser)
	if err != nil {
		http.Error(w, "Error while trying to modify the profile. Try again "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "Can't modify user profile "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
