package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Brolervi/Twitter_Chafa/db"
)

// UserList, reads user list
func UserList(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Page parameter required higher than 0", http.StatusBadRequest)
		return
	}

	pag := int64(pagTemp)

	res, status := db.ReadAllUsers(IDUser, pag, search, typeUser)
	if !status {
		http.Error(w, "Error while reading users", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
