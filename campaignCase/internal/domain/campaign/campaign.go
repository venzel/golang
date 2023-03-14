package campaign

import (
	"campaign/internal/domain/contact"
	"errors"
	"time"

	"github.com/rs/xid"
)

type Campaign struct {
	Id        string
	Name      string
	Content   string
	CreatedAt time.Time
	contact.Contact
}

func CreateCampaign(name *string, content *string, contact *contact.Contact) (*Campaign, error) {
	message, valid := validate(name, content, contact)

	if !valid {
		return nil, errors.New(message)
	}

	return &Campaign{
		Id:        xid.New().String(),
		Name:      *name,
		Content:   *content,
		CreatedAt: time.Now(),
		Contact:   *contact,
	}, nil
}

func validate(name *string, content *string, contact *contact.Contact) (string, bool) {
	if message, valid := validateName(name); !valid {
		return message, false
	}

	if message, valid := validateContent(content); !valid {
		return message, false
	}

	if message, valid := validateContact(contact); !valid {
		return message, false
	}

	return "", true
}

func validateName(name *string) (string, bool) {
	if *name == "" {
		return "name is required", false
	}

	return "", true
}

func validateContent(content *string) (string, bool) {
	if *content == "" {
		return "content is required", false
	}

	return "", true
}

func validateContact(contact *contact.Contact) (string, bool) {
	return "", true
}
