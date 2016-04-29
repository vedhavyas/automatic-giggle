package main

import (
	"log"
	"net/http"
	"github.com/Instamojo/Instamojo++/httphandlers"
	"fmt"
	"os/signal"
	"os"
	"syscall"
)

func main() {
	go func() {
		log.Fatal(http.ListenAndServe(":3000", httphandlers.Router))
	}()
	fmt.Println("listening now ...")
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGTERM)
	<-sigs
	fmt.Println("SIGTERM, time to shutdown")
}
