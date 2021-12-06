package models

// IncomingSendRequest is a native representation
// of incoming request to mailsender.
type IncomingSendRequest struct {
	Receivers []string `json:"receivers"`
	Message   string   `json:"message"`
}

// OutgoingStatusRequest is a native representaiton
// of outgoint request to profile service.
type OutgoingStatusRequest struct {
	Receivers []string `json:"receivers"`
	Message   string   `json:"message"`
	Status    string   `json:"status"`
}
