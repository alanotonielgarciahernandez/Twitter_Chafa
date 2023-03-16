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
	router.HandleFunc("/view-profile", middlewares.CheckDB(middlewares.ValidateJWT(routers.ViewProfile))).Methods("GET")
	router.HandleFunc("/modify-profile", middlewares.CheckDB(middlewares.ValidateJWT(routers.ModifyProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middlewares.CheckDB(middlewares.ValidateJWT(routers.RecordTweet))).Methods("POST")
	router.HandleFunc("/tweets", middlewares.CheckDB(middlewares.ValidateJWT(routers.ReadTweets))).Methods("GET")
	router.HandleFunc("/delete-tweet", middlewares.CheckDB(middlewares.ValidateJWT(routers.DeleteTweet))).Methods("DELETE")

	router.HandleFunc("/upload-avatar", middlewares.CheckDB(middlewares.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/upload-banner", middlewares.CheckDB(middlewares.ValidateJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/get-avatar", middlewares.CheckDB(middlewares.ValidateJWT(routers.GetAvatar))).Methods("GET")
	router.HandleFunc("/get-banner", middlewares.CheckDB(middlewares.ValidateJWT(routers.GetBanner))).Methods("GET")

	router.HandleFunc("/create-relation", middlewares.CheckDB(middlewares.ValidateJWT(routers.CreateRelation))).Methods("POST")
	router.HandleFunc("/delete-relation", middlewares.CheckDB(middlewares.ValidateJWT(routers.DeleteRelation))).Methods("DELETE")
	router.HandleFunc("/check-relation", middlewares.CheckDB(middlewares.ValidateJWT(routers.CheckRelation))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
