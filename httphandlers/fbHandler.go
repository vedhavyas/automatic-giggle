package httphandlers

import (
	"net/http"
	"fmt"
	"github.com/Instamojo/Instamojo++/models"
	"encoding/json"
	"bytes"
	"io/ioutil"
)

var pageToken = "EAAIOkn4DvqABAAaBaiZAXbxHHzEcYsEfrBBFM0QRAZC8ZCAZChKXMtlcK8wTSegMcJxv03fWRbPlOBF8mz460dB2HcJA8M8Ut8FYZBt0TiZBrZBwVRJZB0T1kxyGSLbj0P9tQsnoR3qwswDx0rlhdJQbQoZAc8WYidLUDyrDYmnFJ9QZDZD"

func fbHook(w http.ResponseWriter, r *http.Request) {
	fbModel := &models.FBModel{}
	err := json.NewDecoder(r.Body).Decode(fbModel)
	if err!= nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusOK)
	}

	fmt.Println(fbModel.Entry[0].Messaging[0].Message.Text)
	sendMessage(fbModel, "Text received, echo: Hello,"+fbModel.Entry[0].Messaging[0].Message.Text)
}

func sendMessage(fbModel *models.FBModel, text string) {
	url := "https://graph.facebook.com/v2.6/me/messages?access_token="+pageToken
	message := models.MessageModel{Text:text}
	recipient := models.FBRecipient{ID:fbModel.Entry[0].Messaging[0].Sender.ID}
	payload := models.FBPayload{Message:message, Recipient: recipient}
	data, _ := json.Marshal(payload)
	fmt.Println(string(data))
	req, _ := http.NewRequest("POST", url, bytes.NewReader(data))
	client := http.Client{}
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


