package routes

import (
	"Project_Eular/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.GetLandingPage).Methods("GET")
	r.HandleFunc("/signup", controller.SignUp).Methods("POST")
	return r
}
