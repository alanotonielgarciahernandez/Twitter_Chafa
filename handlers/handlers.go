package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Brolervi/Twitter_Chafa/middlewares"
	"github.com/Brolervi/Twitter_Chafa/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Managers, sets the port
// Handler, listens to the server
func Managers() {
	router := mux.NewRouter()

	router.HandleFunc("/sign-up", middlewares.CheckDB(routers.SignUp)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
