package mailsender

import (
	"mailsender/logic"
	"mailsender/models"
	"mailsender/repository"
)

// Mailsender is a struct for implementing
// service logic.
type Mailsender struct {
	repository repository.Repository
}

// New creates and returns new Mailsender instance
// or errors if any.
func New(r repository.Repository) logic.Logic {
	ms := new(Mailsender)
	ms.repository = r
	return ms
}

// InsertMessage inserts msg to a connected repository.
func (m *Mailsender) InsertMessage(req models.IncomingSendRequest) (string, error) {
	var msg models.Message
	msg.Text = req.Message
	msg.Receivers = req.Receivers
	msg.Status = models.WaitingStatus

	id, err := m.repository.CreateMessage(msg)
	if err != nil {
		return "", err
	}

	return id, nil
}

// UpdateMessage sets the found by id message's status.
func (m *Mailsender) UpdateMessage(id, status string) error {
	if err := m.repository.UpdateMessage(id, status); err != nil {
		return err
	}

	return nil
}
