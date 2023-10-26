package endpoints

import (
	"bytes"
	"emailn/internal/contract"
	internalmock "emailn/internal/test/mock"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignsPost_ShouldSaveNewCampaign(t *testing.T) {
	assert := assert.New(t)
	body := contract.NewCampaign{
		Name:    "name 1",
		Content: "content 1",
		Emails:  []string{"email1@email.com", "email2email.com"},
	}
	service := new(internalmock.CampaignServiceMock)
	service.On("Create", mock.MatchedBy(func(request contract.NewCampaign) bool {
		if request.Name == body.Name &&
			request.Content == body.Content &&
			request.Emails[0] == body.Emails[0] {
			return true
		} else {
			return false
		}
	})).Return("id", nil)
	handler := Handler{CampaignService: service}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/", &buf)
	rr := httptest.NewRecorder()

	json, status, err := handler.CampaignsPost(rr, req)

	assert.Equal(201, status)
	assert.Nil(err)
	assert.NotNil(json)
}

func Test_CampaignsPost_ShouldInformErrorWhenExist(t *testing.T) {
	assert := assert.New(t)
	body := contract.NewCampaign{
		Name:    "name 1",
		Content: "content 1",
		Emails:  []string{"email1@email.com", "email2email.com"},
	}
	service := new(internalmock.CampaignServiceMock)
	service.On("Create", mock.Anything).Return("", fmt.Errorf("error"))
	handler := Handler{CampaignService: service}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/", &buf)
	rr := httptest.NewRecorder()

	_, _, err := handler.CampaignsPost(rr, req)

	assert.NotNil(err)
}
