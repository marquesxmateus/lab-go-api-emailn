package endpoints

import (
	"emailn/internal/contract"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignsPost(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var request contract.NewCampaign
	err := render.DecodeJSON(r.Body, &request)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return nil, 500, err
	}

	id, err := h.CampaignService.Create(request)
	return map[string]string{"id": id}, 201, err
}
