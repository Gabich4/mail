package models

type UserReceivers struct {
	Username  string `json:"username"`
	Receivers string `json:"receivers"`
}

type MailingList map[string]UserReceivers
