package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Brolervi/Twitter_Chafa/db"
	"github.com/Brolervi/Twitter_Chafa/models"
)

// ModifyProfile, modifies the user profile
func ModifyProfile(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Incorrect Data "+err.Error(), 400)
		return
	}

	var status bool

	status, err = db.ModifyLog(t, IDUser)
	if err != nil {
		http.Error(w, "Error while trying modifying log. Try again "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "Can't modify user log"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
