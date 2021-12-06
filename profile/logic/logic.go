package logic

import "profile/models"

type Logic interface {
	Create(username, receivers string) error

	Read(username string) (string, error)

	Update(username, receivers string) error

	Delete(username string) error

	ReadAll() ([]models.UserReceivers, error)
}
