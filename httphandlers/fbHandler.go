package httphandlers

import (
	"net/http"
	"fmt"
	"github.com/Instamojo/Instamojo++/models"
	"encoding/json"
	"bytes"
	"io/ioutil"
	"strconv"
)

var pageToken = "EAAIOkn4DvqABAAaBaiZAXbxHHzEcYsEfrBBFM0QRAZC8ZCAZChKXMtlcK8wTSegMcJxv03fWRbPlOBF8mz460dB2HcJA8M8Ut8FYZBt0TiZBrZBwVRJZB0T1kxyGSLbj0P9tQsnoR3qwswDx0rlhdJQbQoZAc8WYidLUDyrDYmnFJ9QZDZD"

func fbHook(w http.ResponseWriter, r *http.Request) {
	fbModel := &models.FBModel{}
	err := json.NewDecoder(r.Body).Decode(fbModel)
	dat, _ :=ioutil.ReadAll(r.Body)
	fmt.Println(string(dat))
	if err!= nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusOK)
	}

	text := fbModel.Entry[0].Messaging[0].Message.Text
	if text == ""{
		w.WriteHeader(http.StatusOK)
		return
	}
	sendMessage(fbModel, "hello, "+text)
	w.WriteHeader(http.StatusOK)
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
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}


