package routes

import (
	"Project_Eular/Authentication"
	"Project_Eular/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.GetLandingPage).Methods("GET")
	r.HandleFunc("/signup", controller.SignUp).Methods("POST")
	r.HandleFunc("/login", controller.Login).Methods("POST")
	//r.Handle("/login", Authentication.AuthFunc(http.HandlerFunc(controller.Login))).Methods("POST", "OPTIONS")
	r.Handle("/post_question", Authentication.AuthFunc(http.HandlerFunc(controller.InsertQuestions))).Methods("POST", "OPTIONS")
	r.Handle("/post_comment", Authentication.AuthFunc(http.HandlerFunc(controller.Comments))).Methods("POST", "OPTIONS")
	return r
}

// "password" :"john",
// "email" :"john@gmail.com",
// "username" :"John"
