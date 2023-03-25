package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Brolervi/Twitter_Chafa/db"
	"github.com/Brolervi/Twitter_Chafa/models"
)

// SignUp, Gets information to create a user
func SignUp(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error in data received"+err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Password must be longer than 6 characters", http.StatusBadRequest)
		return
	}

	_, founded, _ := db.CheckUserAlreadyExists(t.Email)
	if founded {
		http.Error(w, "Email already taken", http.StatusBadRequest)
		return
	}

	_, status, err := db.InsertUser(t)
	if err != nil {
		http.Error(w, "Error while trying to register "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "Unable to register", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
