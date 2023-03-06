package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Brolervi/Twitter_Chafa/db"
	"github.com/Brolervi/Twitter_Chafa/models"
)

// SignUp, creates user in DB
func SignUp(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error in data received"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Password must be long than 6 characters", 400)
		return
	}

	_, founded, _ := db.CheckUserAlreadyExists(t.Email)
	if founded {
		http.Error(w, "Email already taken", 400)
		return
	}

	_, status, err := db.SetUser(t)
	if err != nil {
		http.Error(w, "Error while trying to sign up "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Could not sign up", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
