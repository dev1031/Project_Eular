package routes

import (
	"Project_Eular/Authentication"
	"Project_Eular/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.GetLandingPage).Methods("GET", "OPTIONS")
	r.HandleFunc("/signup", controller.SignUp).Methods("POST", "OPTIONS")
	r.HandleFunc("/login", controller.Login).Methods("POST", "OPTIONS")
	r.Handle("/post_question", Authentication.AuthFunc(http.HandlerFunc(controller.InsertQuestions))).Methods("POST", "OPTIONS")
	r.Handle("/post_comment", Authentication.AuthFunc(http.HandlerFunc(controller.Comments))).Methods("POST", "OPTIONS")
	r.Handle("/profile/{id}", Authentication.AuthFunc(http.HandlerFunc(controller.Profile))).Methods("GET", "OPTIONS")
	r.Handle("/imageUpload/{id}", Authentication.AuthFunc(http.HandlerFunc(controller.ImageUpload))).Methods("POST", "OPTIONS")
	return r
}

// "password" :"john",
// "email" :"john@gmail.com",
// "username" :"John"
