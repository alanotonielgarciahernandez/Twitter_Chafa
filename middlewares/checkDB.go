package middlewares

import (
	"net/http"

	"github.com/Brolervi/Twitter_Chafa/db"
)

// CheckDB, checks if DB is active
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.PingPong() == 0 {
			http.Error(w, "Lost connection with DB", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
