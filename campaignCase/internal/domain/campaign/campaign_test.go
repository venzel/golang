package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name    = "Tiago"
	content = "Body"
	emails  = []string{"tiago@gmail.com", "marcos@gmail.com"}
)

func Test_CreateCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := CreateCampaign(name, content, emails)

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
}

func Test_CreateCampaign_IdNotNil(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := CreateCampaign(name, content, emails)

	assert.NotNil(campaign.ID)
}

func Test_CreateCampaign_CreatedMostBeNow(t *testing.T) {
	assert := assert.New(t)

	datetime := time.Now().Add(-time.Minute)

	campaign, _ := CreateCampaign(name, content, emails)

	assert.Greater(campaign.CreatedAt, datetime)
}

func Test_CreateCampaign_MostValidateName(t *testing.T) {
	assert := assert.New(t)

	_, err := CreateCampaign("", content, emails)

	assert.Equal("name is required", err.Error())
}
