package httphandlers

import (
	"net/http"
	"fmt"
	"github.com/Instamojo/Instamojo++/models"
	"encoding/json"
	"strings"
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

func sendMessage(fbModel models.FBModel, text string) {
	url := "https://graph.facebook.com/v2.6/me/messages?access_token="+pageToken
	message := models.MessageModel{Text:text}
	recipient := models.FBRecipient{ID:fbModel.Entry[0].ID}
	payload := models.FBPayload{Message:message, Recipient: recipient}
	req, _ := http.NewRequest("POST", url, json.Marshal(payload))
	client := http.Client{}
	_, err :=client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
}


