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
