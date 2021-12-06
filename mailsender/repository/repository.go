package repository

import "mailsender/models"

type Repository interface {
	// CreateMessage inserts msg into repository.
	CreateMessage(msg models.Message) (string, error)

	// UpdateMessage updates message in the repository.
	UpdateMessage(id, status string) error
}
