package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internalErrors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type respositoryMock struct {
	mock.Mock
}

func (r *respositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *respositoryMock) Update(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *respositoryMock) Get() ([]Campaign, error) {
	return nil, nil
}

func (r *respositoryMock) GetBy(id string) (*Campaign, error) {
	args := r.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Campaign), args.Error(1)
}

func (r *respositoryMock) Delete(campaign *Campaign) error {
	return nil
}

var (
	newCampaign = contract.NewCampaign{
		Name:    "Campaign X",
		Content: "Content X",
		Emails:  []string{"teste1@email.com", "teste2@email.com"},
	}
	service = ServiceImp{}
)

func Test_Create_CreateCampaign(t *testing.T) {

	repositoryMock := new(respositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(nil)
	service.Repository = repositoryMock

	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)
	newCampaign := newCampaign
	newCampaign.Name = ""

	_, err := service.Create(newCampaign)

	assert.NotNil(err)
	assert.Equal("name is required", err.Error())
}

func Test_Create_SaveCampaign(t *testing.T) {

	repositoryMock := new(respositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		return campaign.Name == newCampaign.Name &&
			campaign.Content == newCampaign.Content &&
			len(campaign.Contacts) == len(newCampaign.Emails)
	})).Return(nil)
	service.Repository = repositoryMock

	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(respositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(internalerrors.ErrInternal)
	service.Repository = repositoryMock

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(internalerrors.ErrInternal, err))
}

func Test_GetBy_ShouldReturnCampaign(t *testing.T) {
	assert := assert.New(t)
	campaign, _ := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	repositoryMock := new(respositoryMock)
	repositoryMock.On("GetBy", mock.MatchedBy(func(id string) bool {
		return id == campaign.ID
	})).Return(campaign, nil)
	service.Repository = repositoryMock

	campaignRetorned, _ := service.GetBy(campaign.ID)

	assert.Equal(campaign.ID, campaignRetorned.ID)
	assert.Equal(campaign.Name, campaignRetorned.Name)
	assert.Equal(campaign.Content, campaignRetorned.Content)
	assert.Equal(campaign.Status, campaignRetorned.Status)
}

func Test_GetBy_ShouldReturnErrorWhenSomethingWrongExist(t *testing.T) {
	assert := assert.New(t)
	campaign, _ := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	repositoryMock := new(respositoryMock)
	repositoryMock.On("GetBy", mock.Anything).Return(nil, errors.New("error"))
	service.Repository = repositoryMock

	_, err := service.GetBy(campaign.ID)

	assert.Equal(internalerrors.ErrInternal.Error(), err.Error())
}
