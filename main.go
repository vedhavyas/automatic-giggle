package main

import (
	"log"
	"net/http"
	"github.com/Instamojo/Instamojo++/httphandlers"
)

func main() {
	go log.Fatal(http.ListenAndServe(":3000", httphandlers.Router))
}
