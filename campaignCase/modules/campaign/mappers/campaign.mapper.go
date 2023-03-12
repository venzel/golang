package mappers

import (
	"campaign/internal/domain/campaign"
	"campaign/internal/domain/contact"
	"campaign/modules/campaign/dtos"
)

func ToEntity(dto *dtos.CreateCampaignDto) *campaign.Campaign {
	contacts := contact.Contact{
		Emails: dto.Emails,
	}

	campaign, _ := campaign.CreateCampaign(dto.Name, dto.Name, contacts)

	return campaign
}

func ToDto(campaign *campaign.Campaign) *dtos.ResponseCampaignDto {
	return &dtos.ResponseCampaignDto{
		Id:   campaign.Id,
		Name: campaign.Name,
	}
}
