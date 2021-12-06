package models

// IncomingRequest - request from client.
type IncomingRequest struct {
	Receivers  []string            `json:"receivers"`
	TemplateId string              `json:"template_id"`
	Parameters []TemplateParameter `json:"parameters"`
}

// MailsenderRequest - request to mailsender.
type MailsenderRequest struct {
	Receivers []string `json:"receivers"`
	Message   string   `json:"message"`
}
