package models

type SendRequestStatus struct {
	Receivers []string `json:"receivers"`
	Message   string   `json:"message"`
	Status    string   `json:"status"`
}
