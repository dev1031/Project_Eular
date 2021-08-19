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
	r.Handle("/login", Authentication.AuthFunc(http.HandlerFunc(controller.Login))).Methods("POST", "OPTIONS")
	r.HandleFunc("/insertQues", controller.InsertQuestions).Methods("POST")
	return r
}
