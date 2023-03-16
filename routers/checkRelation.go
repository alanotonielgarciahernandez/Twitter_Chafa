package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Brolervi/Twitter_Chafa/db"
	"github.com/Brolervi/Twitter_Chafa/models"
)

// CheckRelation, checks if there is a relation between 2 users
func CheckRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

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
