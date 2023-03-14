package campaign

import (
	"campaign/internal/domain/contact"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	empty    = ""
	name     = "Tiago"
	content  = "Body"
	emails   = []string{"tiago@gmail.com", "marcos@gmail.com"}
	contact_ = contact.Contact{
		Emails: emails,
	}
)

func Test_CreateCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := CreateCampaign(&name, &content, &contact_)

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
}

func Test_CreateCampaign_IdNotNil(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := CreateCampaign(&name, &content, &contact_)

	assert.NotNil(campaign.Id)
}

func Test_CreateCampaign_CreatedMostBeNow(t *testing.T) {
	assert := assert.New(t)

	datetime := time.Now().Add(-time.Minute)

	campaign, _ := CreateCampaign(&name, &content, &contact_)

	assert.Greater(campaign.CreatedAt, datetime)
}

func Test_CreateCampaign_MostValidateName(t *testing.T) {
	assert := assert.New(t)

	_, err := CreateCampaign(&empty, &content, &contact_)

	assert.Equal("name is required", err.Error())
}

func Test_CreateCampaign_MostValidateContent(t *testing.T) {
	assert := assert.New(t)

	_, err := CreateCampaign(&name, &empty, &contact_)

	assert.Equal("content is required", err.Error())
}
