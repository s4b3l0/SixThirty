package domain

type Message struct {
	Id          int64
	MessageText string `json:"message_text"`
	MessageTime string `json:"message_time"`
}
