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
	router.HandleFunc("/sign-in", middlewares.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/view-profile", middlewares.CheckDB(middlewares.ValidateJWT(routers.ViewProfile))).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
