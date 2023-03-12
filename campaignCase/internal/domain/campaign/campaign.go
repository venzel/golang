package campaign

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string
}

type Campaign struct {
	ID        string
	Name      string
	Content   string
	CreatedAt time.Time
	Contacts  []Contact
}

func CreateCampaign(name string, content string, emails []string) (*Campaign, error) {
	contacts := make([]Contact, len(emails))

	message, err := validate(&name, &content, &emails)

	if err {
		return nil, errors.New(message)
	}

	for index, value := range emails {
		contacts[index].Email = value
	}

	return &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedAt: time.Now(),
		Contacts:  contacts,
	}, nil
}

func validate(name *string, content *string, emails *[]string) (string, bool) {
	if *name == "" {
		return "name is required", true
	}

	return "", false
}
