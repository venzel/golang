package services

import (
	"campaign/modules/campaign/dtos"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Create(dto *dtos.CreateCampaignDto) (*dtos.ResponseCampaignDto, error) {
	args := r.Called(dto)
	return &dtos.ResponseCampaignDto{}, args.Error(0)
}

func Test_CreateCampaign(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(repositoryMock)

	service := CampaignService{Repository: repositoryMock}

	createDto := dtos.CreateCampaignDto{
		Name:    "tiago",
		Content: "body",
		Emails:  []string{"tiago@gmail.com"},
	}

	responseDto := service.Create(&createDto)

	assert.Equal(responseDto.Name, createDto.Name)
	assert.Equal(responseDto.Content, createDto.Content)
}
