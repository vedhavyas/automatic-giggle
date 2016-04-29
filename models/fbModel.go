package models

import "time"

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

type RAPSuccess struct {
	PaymentRequest struct {
			       ID string `json:"id"`
			       Phone string `json:"phone"`
			       Email string `json:"email"`
			       BuyerName string `json:"buyer_name"`
			       Amount string `json:"amount"`
			       Purpose string `json:"purpose"`
			       Status string `json:"status"`
			       SendSms bool `json:"send_sms"`
			       SendEmail bool `json:"send_email"`
			       SmsStatus string `json:"sms_status"`
			       EmailStatus string `json:"email_status"`
			       Shorturl interface{} `json:"shorturl"`
			       Longurl string `json:"longurl"`
			       RedirectURL string `json:"redirect_url"`
			       Webhook string `json:"webhook"`
			       CreatedAt time.Time `json:"created_at"`
			       ModifiedAt time.Time `json:"modified_at"`
			       AllowRepeatedPayments bool `json:"allow_repeated_payments"`
		       } `json:"payment_request"`
	Success bool `json:"success"`
}

type RAPFailure struct {
	Message struct {
			Amount []string `json:"amount"`
			Phone []string `json:"phone"`
		} `json:"message"`
	Success bool `json:"success"`
}