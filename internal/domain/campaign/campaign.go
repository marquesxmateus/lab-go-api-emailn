package campaign

import (
	internalerrors "emailn/internal/internalErrors"
	"time"

	"github.com/rs/xid"
)

const (
	Canceled string = "Canceled"
	Deleted  string = "Deleted"
	Pending  string = "Pending"
	Started  string = "Started"
	Finished string = "Finished"
)

type Contact struct {
	ID         string `gorm:"size:50"`
	CampaignID string `gorm:"size:50"`
	Email      string `validate:"required,email" gorm:"size:500"`
}

type Campaign struct {
	ID        string    `validate:"required" gorm:"size:50"`
	Name      string    `validate:"required,min=5,max=24" gorm:"size:100"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"required,min=5,max=1024" gorm:"size:1024"`
	Contacts  []Contact `validate:"min=1,dive"`
	Status    string    `gorm:"size:20"`
}

func (c *Campaign) Cancel() {
	c.Status = Canceled
}

func (c *Campaign) Delete() {
	c.Status = Deleted
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	contacts := make([]Contact, len(emails))
	for i, email := range emails {
		contacts[i] = Contact{
			ID:    xid.New().String(),
			Email: email,
		}
	}

	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedOn: time.Now(),
		Content:   content,
		Contacts:  contacts,
		Status:    Pending,
	}

	err := internalerrors.ValidateStruct(campaign)
	if err != nil {
		return nil, err
	}

	return campaign, nil
}
