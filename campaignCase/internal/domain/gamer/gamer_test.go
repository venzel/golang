package gamer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateGamer(t *testing.T) {
    assert := assert.New(t);

    name := "Tiago"
    winners := 10

    emails := []string{"tiago@gmail.com", "marcos@gmail.com"}

    contacts := Contact{
        Emails: emails,
    }

    gamer := CreateGamer(name, winners, contacts)

    assert.Equal(gamer.ID, "1")
    assert.Equal(gamer.Name, name)
    assert.Equal(gamer.Winners, winners)
}