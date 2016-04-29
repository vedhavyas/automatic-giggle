package main

import (
	"log"
	"net/http"
	"github.com/Instamojo/Instamojo++/httphandlers"
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

func main() {

	port := "3000"

	if os.Getenv("PORT") != ""{
		port = os.Getenv("PORT")
	}
	go func() {
		log.Fatal(http.ListenAndServe(":"+port, httphandlers.Router))
	}()
	fmt.Println("listening on port - "+port)
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGTERM)
	<-sigs
	fmt.Println("SIGTERM, time to shutdown")
}
