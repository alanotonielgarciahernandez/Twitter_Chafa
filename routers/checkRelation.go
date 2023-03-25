package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Brolervi/Twitter_Chafa/db"
	"github.com/Brolervi/Twitter_Chafa/models"
)

// CheckRelation, Gets a relation ID to check if exists
func CheckRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Missing ID parameter", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UserID = IDUser
	t.UserRelationID = ID

	var res models.RelationResponse

	status, err := db.CheckRelation(t)
	if err != nil || !status {
		res.Status = false
	} else {
		res.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
