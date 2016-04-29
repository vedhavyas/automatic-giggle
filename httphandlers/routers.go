package httphandlers

import "github.com/gorilla/mux"

var Router = mux.NewRouter()

func init() {
	Router.HandleFunc("/spreadsheet/init/{id}", getSpreadSheetAuthUrl).Methods("GET")
	Router.HandleFunc("/spreadsheet/init/", saveSpreadSheetAuthToken).Methods("POST")
	Router.HandleFunc("/spreadsheet/post/", addToSpreadSheet).Methods("POST")
	Router.HandleFunc("/dropbox/init/", saveDropBoxAuthToken).Methods("POST")
	Router.HandleFunc("/dropbox/post/", pushToDropBox).Methods("POST")
}