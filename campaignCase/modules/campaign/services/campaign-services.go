package services

import (
	"campaign/modules/campaign/dtos"
	"campaign/modules/campaign/repositories"
)

type CampaignService struct {
	Repository repositories.CampaignRepository
}

func (s *CampaignService) Create(dto *dtos.CreateCampaignDto) *dtos.ResponseCampaignDto {
	create, _ := s.Repository.Create(dto)

	// contact, _ := contact.CreateContact(&dto.Emails)

	// campaign, _ := campaign.CreateCampaign(&dto.Name, &dto.Content, contact)

	return create
}
