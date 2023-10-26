package main

import (
	"emailn/internal/domain/campaign"
	"emailn/internal/endpoints"
	"emailn/internal/infrastructure/database"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	db, _ := database.NewDb()
	campaignService := campaign.ServiceImp{
		Repository: &database.CampaignRepository{Db: db},
	}
	handler := endpoints.Handler{
		CampaignService: &campaignService,
	}
	r.Post("/campaigns", endpoints.HandlerError(handler.CampaignsPost))
	r.Get("/campaigns/{id}", endpoints.HandlerError(handler.CampaignsGetById))
	r.Patch("/campaigns/cancel/{id}", endpoints.HandlerError(handler.CampaignsCancelPatch))
	r.Delete("/campaigns/{id}", endpoints.HandlerError(handler.CampaignsDelete))

	http.ListenAndServe(":3000", r)
}
