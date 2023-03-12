package repositories

import "campaign/modules/campaign/dtos"

type CampaignRepository interface {
	Create(dto *dtos.CreateCampaignDto) *dtos.ResponseCampaignDto
}
