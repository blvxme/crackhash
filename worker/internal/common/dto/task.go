package dto

type TaskRequest struct {
	RequestId  string `json:"requestId"`
	Alphabet   string `json:"alphabet"`
	Hash       string `json:"hash"`
	MaxLength  int    `json:"maxLength"`
	PartNumber int    `json:"partNumber"`
	PartCount  int    `json:"partCount"`
}

type TaskResponse struct {
	RequestId string   `json:"requestId"`
	Data      []string `json:"data"`
}
