package models

type FBModel struct {
	Object string `json:"object"`
	Entry []struct {
		ID int64 `json:"id"`
		Time int64 `json:"time"`
		Messaging []struct {
			Sender struct {
				       ID int64 `json:"id"`
			       } `json:"sender"`
			Recipient struct {
				       ID int64 `json:"id"`
			       } `json:"recipient"`
			Timestamp int64 `json:"timestamp"`
			Message struct {
				       Mid string `json:"mid"`
				       Seq int `json:"seq"`
				       Text string `json:"text"`
			       } `json:"message"`
		} `json:"messaging"`
	} `json:"entry"`
}

type MessageModel struct {
	Text string `json:"text"`
}

type FBRecipient struct {
	ID string `json:"id"`
}

type FBPayload struct {
	Recipient FBRecipient `json:"recipient"`
	Message MessageModel `json:"message"`
}

type RAP struct {
	Phone string `json:"phone"`
	Purpose string `json:"purpose"`
	Amount string `json:"amount"`
	SendSMS bool `json:"send_sms"`

}