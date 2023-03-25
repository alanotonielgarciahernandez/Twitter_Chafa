package routers

import (
	"net/http"

	"github.com/Brolervi/Twitter_Chafa/db"
	"github.com/Brolervi/Twitter_Chafa/models"
)

// CreateRelation, Creates a relation model
func CreateRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Missing ID parameter", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UserID = IDUser
	t.UserRelationID = ID

	status, err := db.InsertRelation(t)
	if err != nil {
		http.Error(w, "Error while inserting data "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "Can't insert data", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
