package contact

import "errors"

type Contact struct {
	Emails []string
}

func CreateContact(emails *[]string) (*Contact, error) {
	message, valid := validate(emails)

	if !valid {
		return nil, errors.New(message)
	}

	return &Contact{
		Emails: *emails,
	}, nil
}

func validate(emails *[]string) (string, bool) {
	if message, valid := validateEmails(emails); !valid {
		return message, false
	}

	return "", true
}

func validateEmails(emails *[]string) (string, bool) {
	if len(*emails) <= 0 {
		return "emails is required", false
	}

	return "", true
}
