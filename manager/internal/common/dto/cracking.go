package dto

type CrackingRequest struct {
	Hash      string `json:"hash"`
	MaxLength int    `json:"maxLength"`
}

type CrackingResponse struct {
	RequestId string `json:"requestId"`
}
