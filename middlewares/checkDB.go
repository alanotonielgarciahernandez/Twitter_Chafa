package middlewares

import (
	"net/http"

	"github.com/Brolervi/Twitter_Chafa/db"
)

// CheckDB, Checks if DB is active
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckDBConnection() == 0 {
			http.Error(w, "Lost connection with DB", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	}
}
