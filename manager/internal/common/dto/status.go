package dto

type Status string

const (
	StatusInProgress Status = "IN_PROGRESS"
	StatusReady      Status = "READY"
	StatusError      Status = "ERROR"
)

type RequestStatus struct {
	Status Status   `json:"status"`
	Data   []string `json:"data"`
}
