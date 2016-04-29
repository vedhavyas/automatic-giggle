package httphandlers

import (
	"net/http"
	"fmt"
	"github.com/Instamojo/Instamojo++/models"
	"encoding/json"
	"bytes"
	"io/ioutil"
	"strconv"
	"strings"
)

var pageToken = "EAAIOkn4DvqABAAaBaiZAXbxHHzEcYsEfrBBFM0QRAZC8ZCAZChKXMtlcK8wTSegMcJxv03fWRbPlOBF8mz460dB2HcJA8M8Ut8FYZBt0TiZBrZBwVRJZB0T1kxyGSLbj0P9tQsnoR3qwswDx0rlhdJQbQoZAc8WYidLUDyrDYmnFJ9QZDZD"

var s = []string{}
var read bool = false

func fbHook(w http.ResponseWriter, r *http.Request) {
	fbModel := &models.FBModel{}
	err := json.NewDecoder(r.Body).Decode(fbModel)
	if err!= nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusOK)
	}
	dat, _ :=json.MarshalIndent(fbModel,"", " ")
	fmt.Println(string(dat))
	text := fbModel.Entry[0].Messaging[0].Message.Text
	text = strings.ToLower(text)
	if text == "" {
		return
	}

	if text == "rap"{
		s = nil
		read = true
		sendMessage(fbModel, "Enter Phone Number")
		return
	} else if read {
		if text == "cancel"{
			s = nil
			read = false
			sendMessage(fbModel, "Okay, Done :-D")
			sendMessage(fbModel, "Type \"rap\" to start a Request a Payment flow!")
			return
		}
		s = append(s, text)
		if len(s) == 1 {
			sendMessage(fbModel, "Enter purpose")
			return
		} else if len(s) == 2{
			sendMessage(fbModel, "Enter amount")
			return
		} else {
			sendMessage(fbModel, "Sending RAP")
			err := requestPayment(s)
			read = false
			if err != nil{
				sendMessage(fbModel, "Failed to Send RAP :(")
			}else {
				sendMessage(fbModel, "RAP Sent :-D")
				sendMessage(fbModel, "Type \"rap\" to start a Request a Payment flow. Again!")
			}
			w.WriteHeader(http.StatusOK)
			return
		}
	} else {
		s = nil
		read = false
		sendMessage(fbModel, "I do not have that feature yet. Wait for it bitch :p")
		sendMessage(fbModel, "Type \"rap\" to start a Request a Payment flow!")
		return
	}

}

func requestPayment(s []string) error {
	url := "https://www.instamojo.com/api/1.1/payment-requests/"
	apikey := "4cb7fe8523302dc1a78dbddddcb7c6c1"
	authToken := "c7c46720938f41b1c82de53d4d8c745e"
	rap := models.RAP{Phone:s[0], Purpose:s[1], Amount:s[2], SendSMS:true}
	data, _ := json.Marshal(rap)
	fmt.Println(string(data))
	req, _ := http.NewRequest("POST", url, bytes.NewReader(data))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Api-Key", apikey)
	req.Header.Add("X-Auth-Token", authToken)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		return err
	}

	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return nil
}

func sendMessage(fbModel *models.FBModel, text string) {
	url := "https://graph.facebook.com/v2.6/me/messages?access_token="+pageToken
	message := models.MessageModel{Text:text}
	recipient := models.FBRecipient{ID:strconv.FormatInt(fbModel.Entry[0].Messaging[0].Sender.ID, 10)}
	payload := models.FBPayload{Message:message, Recipient: recipient}
	data, _ := json.Marshal(payload)
	fmt.Println(string(data))
	req, _ := http.NewRequest("POST", url, bytes.NewReader(data))
	client := http.Client{}
	req.Header.Add("Content-Type", "application/json")
	resp, err :=client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}


