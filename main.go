package main

import (
	"Project_Eular/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/jesseokeya/go-httplogger"
)

func main() {

	r := routes.Router()
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", httplogger.Golog(r)))
}
