package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Brolervi/Twitter_Chafa/db"
	"github.com/Brolervi/Twitter_Chafa/jwt"
	"github.com/Brolervi/Twitter_Chafa/models"
)

// SignIn, Gets the user's login information
func SignIn(w http.ResponseWriter, r *http.Request) {

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Email and/or Password incorrect "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is required ", http.StatusBadRequest)
		return
	}

	doc, exists := db.TryLogin(t.Email, t.Password)
	if !exists {
		http.Error(w, "Email and/or Password incorrect ", http.StatusBadRequest)
		return
	}

	jwtKey, err := jwt.GenJWT(doc)
	if err != nil {
		http.Error(w, "JWT Error "+err.Error(), http.StatusBadRequest)
		return
	}

	resp := models.LoginResponse{
		Token: jwtKey,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expTime,
	})
}
