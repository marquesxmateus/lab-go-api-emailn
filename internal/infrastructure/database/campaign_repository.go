package database

import (
	"emailn/internal/domain/campaign"

	"gorm.io/gorm"
)

type CampaignRepository struct {
	Db *gorm.DB
}

func (r *CampaignRepository) Save(campaign *campaign.Campaign) error {
	tx := r.Db.Create(campaign)
	return tx.Error
}

func (r *CampaignRepository) Update(campaign *campaign.Campaign) error {
	tx := r.Db.Save(campaign)
	return tx.Error
}

func (r *CampaignRepository) Get() ([]campaign.Campaign, error) {
	var campaigns []campaign.Campaign
	r.Db.Find(&campaigns)
	return campaigns, nil
}

func (r *CampaignRepository) GetBy(id string) (*campaign.Campaign, error) {
	var campaign campaign.Campaign
	tx := r.Db.Preload("Contacts").First(&campaign, "id = ?", id)
	return &campaign, tx.Error
}

func (r *CampaignRepository) Delete(campaign *campaign.Campaign) error {
	tx := r.Db.Select("Contacts").Delete(campaign)
	return tx.Error
}
