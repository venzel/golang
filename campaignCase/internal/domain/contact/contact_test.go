package contact

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	emails = []string{
		"tiago@gmail.com",
	}
)

func Test_CreateContact(t *testing.T) {
	assert := assert.New(t)

	contact, _ := CreateContact(&emails)

	assert.Equal(len(contact.Emails), 1)
}
