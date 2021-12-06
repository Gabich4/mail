package profile

import (
	"errors"
	"profile/logic"
	"profile/models"
)

type Profile struct {
	ml models.MailingList
}

func New() logic.Logic {
	ml := make(models.MailingList, 0)
	p := new(Profile)
	p.ml = ml
	return p
}

func (p *Profile) Create(username, receivers string) error {
	if _, ok := p.ml[username]; ok {
		return errors.New("user already exists")
	}

	p.ml[username] = models.UserReceivers{
		Username:  username,
		Receivers: receivers,
	}

	return nil
}

func (p *Profile) Read(username string) (string, error) {
	if u, ok := p.ml[username]; !ok {
		return "", errors.New("user not found")
	} else {
		return u.Receivers, nil
	}
}

func (p *Profile) Update(username, receivers string) error {
	if _, ok := p.ml[username]; !ok {
		return errors.New("user not found")
	}

	p.ml[username] = models.UserReceivers{
		Username:  username,
		Receivers: receivers,
	}
	return nil
}

func (p *Profile) Delete(username string) error {
	if _, ok := p.ml[username]; !ok {
		return errors.New("user not found")
	}

	delete(p.ml, username)
	return nil
}

func (p *Profile) ReadAll() ([]models.UserReceivers, error) {
	if len(p.ml) == 0 {
		return nil, errors.New("mailing list is empty")
	}

	ur := make([]models.UserReceivers, 0, len(p.ml))
	for _, receivers := range p.ml {
		ur = append(ur, receivers)
	}

	return ur, nil
}
