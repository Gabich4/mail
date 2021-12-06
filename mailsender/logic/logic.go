package logic

import "mailsender/models"

// Logic is a base interface for
// API layer logic.
type Logic interface {
	// InsertMessage inserts message to the connected repository.
	InsertMessage(inReq models.IncomingSendRequest) (string, error)

	// UpdateMessage updates message status in the connected repository.
	UpdateMessage(id string, status string) error
}
