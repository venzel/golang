package services

import (
	"campaign/modules/campaign/dtos"
	"campaign/modules/campaign/repositories"
)

type CampaignService struct {
	Repository repositories.CampaignRepository
}

func (s *CampaignService) Create(dto *dtos.CreateCampaignDto) *dtos.ResponseCampaignDto {
	return &dtos.ResponseCampaignDto{}
}
