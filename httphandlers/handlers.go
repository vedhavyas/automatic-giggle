package httphandlers

import (
	"net/http"
	"io/ioutil"
	"log"
	"golang.org/x/oauth2/google"
	"fmt"
	"golang.org/x/oauth2"
	"os"
	"encoding/json"
	"github.com/gorilla/mux"
	"google.golang.org/api/script/v1"
	"golang.org/x/net/context"
	"github.com/tj/go-dropy"
	"github.com/tj/go-dropbox"
	"time"
)

var config *oauth2.Config
var ctx context.Context
var client *http.Client
var scriptId string = "MTXyn5GPII68lkqWo_hdyxNowMYC71MR5"

func init() {
	ctx = context.Background()
	clientFile, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		log.Fatal("Unable to read client secret file - "+err.Error())
	}

	config, err = google.ConfigFromJSON(clientFile, "https://www.googleapis.com/auth/drive", "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
}

func getSpreadSheetAuthUrl(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	userId, _ := vars["id"]
	if userId == ""{
		log.Fatal("user_id is missing in the request")
	}

	userConfigFile := "spreadsheets-configs/"+userId
	_, err := ioutil.ReadFile(userConfigFile)
	if err != nil {
		os.Remove(userConfigFile)
	}

	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	data, _ := json.MarshalIndent(spreadsheetResponse{authURL}, "", " ")
	fmt.Fprint(w, string(data))
}

func saveSpreadSheetAuthToken(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("auth_token")
	if code == ""{
		log.Fatal("Auth token is nil")
	}

	userId := r.FormValue("user_id")
	if userId == "" {
		log.Fatal("user id is nil")
	}

	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}

	fileName := "spreadsheets-configs/"+userId
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(tok)
	w.WriteHeader(http.StatusOK)
}

func tokenFromFile(userId string) (*oauth2.Token, error) {
	f, err := os.Open("spreadsheets-configs/"+userId)
	if err != nil {
		return nil, err
	}
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return t, err
}

func getRequestResponse(req script.ExecutionRequest)  *script.Operation{
	srv, err := script.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve script Client %v", err)
	}

	resp, err := srv.Scripts.Run(scriptId, &req).Do()
	if err != nil {
		log.Fatalf("Unable to execute Apps Script function. %v", err)
		return nil
	}

	if resp.Error != nil {
		err := resp.Error.Details[0].(map[string]interface{})
		fmt.Printf("Script error message: %s\n", err["errorMessage"]);

		if (err["scriptStackTraceElements"] != nil) {
			fmt.Printf("Script error stacktrace:\n")
			for _, trace := range err["scriptStackTraceElements"].([]interface{}) {
				t := trace.(map[string]interface{})
				fmt.Printf("\t%s: %d\n", t["function"], int(t["lineNumber"].(float64)))
			}
		}

		return nil
	}

	return resp
}

func addToSpreadSheet(w http.ResponseWriter, r *http.Request) {
	//userId := r.FormValue("user_id")
	//if userId == ""{
	//	log.Fatal("user_id is nil")
	//}

	//hardcode the userID for now
	userId := "123456"
	tok, _ := tokenFromFile(userId)
	client = config.Client(ctx, tok)

	//get url of the target sheet
	req := script.ExecutionRequest{Function:"getUrlOfSheet"}
	resp := getRequestResponse(req)
	if resp == nil{
		log.Fatal("response is nil")
	}

	result := resp.Response.(map[string]interface{})
	spreadSheetUrl := result["result"].(string)

	//Append the row to the spread sheet
	paymentID := r.FormValue("payment_id")
	purpose := r.FormValue("purpose")
	phone := r.FormValue("buyer_phone")
	email := r.FormValue("buyer")
	name := r.FormValue("buyer_name")
	amount := r.FormValue("amount")
	currency := r.FormValue("currency")
	fees := r.FormValue("fees")
	status := r.FormValue("status")
	params := []interface{}{spreadSheetUrl, paymentID, email, purpose, phone, name, amount, fees, status, currency}
	req = script.ExecutionRequest{Function:"addRowToFile", Parameters:params}
	resp = getRequestResponse(req)
	if resp == nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func saveDropBoxAuthToken(w http.ResponseWriter, r *http.Request)  {
	userId := r.FormValue("user_id")
	if userId == "" {
		log.Fatal("user id is nil")
	}
	token := r.FormValue("auth_token")
	if token == "" {
		log.Fatal("dropbox auth token is null")
	}

	fileName := "dropbox-configs/"+userId
	_, err := os.Open(fileName)
	if err != nil {
		os.Remove(fileName)
	}

	configFile := dropBoxConfigFile{token}
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(configFile)
	w.WriteHeader(http.StatusOK)
}

func getDropBoxToken(userId string) (string, error) {
	file, err := os.Open("dropbox-configs/"+userId)
	if err != nil {
		return "", err
	}

	configFile := &dropBoxConfigFile{}
	json.NewDecoder(file).Decode(configFile)
	return configFile.Token, nil
}

func pushToDropBox(w http.ResponseWriter, r *http.Request)  {
	userId := r.FormValue("user_id")
	if userId == "" {
		log.Fatal("user id nil")
	}
	token, err := getDropBoxToken(userId)
	if err != nil {
		log.Fatal(err)
	}

	client := dropy.New(dropbox.New(dropbox.NewConfig(token)))

	file, err := os.Open("MOJO5c25001M40477703.pdf")
	if err != nil {
		log.Fatal(err)
	}
	currentTime := time.Now().Format("2006-01-02:15:04:05")
	fileName := fmt.Sprintf("/Instamojo-payout-%v.pdf", currentTime)
	err = client.Upload(fileName, file)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

type spreadsheetResponse struct {
	Url string `json:"url"`

}

type dropBoxConfigFile struct {
	Token string `json:"auth_token"`
}


