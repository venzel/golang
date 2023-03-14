package mappers

import (
	"campaign/internal/domain/campaign"
	"campaign/internal/domain/contact"
	"campaign/modules/campaign/dtos"
	"errors"
)

type CampaignMapper struct{}

func (c *CampaignMapper) ToEntity(dto *dtos.CreateCampaignDto) (*campaign.Campaign, error) {
	contact, err := contact.CreateContact(&dto.Emails)

	if err != nil {
		return nil, errors.New("ocorreu um erro na criação do contato")
	}

	campaign, err := campaign.CreateCampaign(&dto.Name, &dto.Name, contact)

	if err != nil {
		return nil, errors.New("ocorreu um erro na criação da campanha")
	}

	return campaign, nil
}

func (c *CampaignMapper) ToDto(campaign *campaign.Campaign) *dtos.ResponseCampaignDto {
	return &dtos.ResponseCampaignDto{
		Id:   campaign.Id,
		Name: campaign.Name,
	}
}
