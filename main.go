package main

import (
	"log"
	"net/http"
	"github.com/Instamojo/Instamojo++/httphandlers"
)

func main() {
	log.Fatal(http.ListenAndServe(":8090", httphandlers.Router))
}
