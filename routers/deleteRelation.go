package routers

import (
	"net/http"

	"github.com/Brolervi/Twitter_Chafa/db"
	"github.com/Brolervi/Twitter_Chafa/models"
)

// DeleteRelation, Gets a relation ID to delete it
func DeleteRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Missing ID parameter", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UserID = IDUser
	t.UserRelationID = ID

	status, err := db.DeleteRelation(t)
	if err != nil {
		http.Error(w, "Error while removing relation "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "Can't remove relation", http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
}
