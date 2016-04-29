package main

import (
	"log"
	"net/http"
	"github.com/Instamojo/Instamojo++/httphandlers"
	"os"
)

func main() {

	port := "3000"

	if os.Getenv("PORT") != ""{
		port = os.Getenv("PORT")
	}
	log.Fatal(http.ListenAndServe(":"+port, httphandlers.Router))
}
