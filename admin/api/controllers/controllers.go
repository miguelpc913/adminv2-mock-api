package controllers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	AdminMiddleware "github.com/tiqueteo/adminv2-mock-api/api/middleware"
	"github.com/tiqueteo/adminv2-mock-api/api/services"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	sm := services.NewServiceManager(db)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true, // Set to true if you want to allow credentials
		MaxAge:           300,  // Maximum value not ignored by any of the major browsers
	}))

	r.Post("/login", sm.Login)
	r.Route("/products", func(r chi.Router) {
		r.Use(AdminMiddleware.CheckJTW)
		r.Get("/", sm.GetProducts)
	})
	r.Route("/venueCapacities", func(r chi.Router) {
		r.Use(AdminMiddleware.CheckJTW)
		r.Get("/", sm.GetVenues)
	})
	r.Route("/salesGroups", func(r chi.Router) {
		r.Use(AdminMiddleware.CheckJTW)
		r.Get("/", sm.GetSalesGroups)
	})
	r.Route("/buyerTypes", func(r chi.Router) {
		r.Use(AdminMiddleware.CheckJTW)
		r.Get("/", sm.GetBuyerTypes)
	})
	r.Route("/productTags", func(r chi.Router) {
		r.Use(AdminMiddleware.CheckJTW)
		r.Get("/", sm.GetProductTags)
	})
	r.Route("/productInfos", func(r chi.Router) {
		r.Use(AdminMiddleware.CheckJTW)
		r.Get("/", sm.GetProductInfos)
		r.Get("/{id}", sm.GetProductInfoById)
		r.Put("/{id}/identity", sm.PutProductInfoIdentity)
		r.Put("/{id}/configurations", sm.PutProductInfoSettings)
		r.Put("/{id}/salesGroups", sm.PutProductInfoSalesGroups)
		r.Put("/{id}/products", sm.PutProductInfoProducts)
		r.Put("/{id}/venues", sm.PutProductInfoVenue)
		r.Post("/", sm.PostProductInfos)
		r.Get("/types", sm.GetProductInfoType)
		r.Put("/order", sm.PutOrderProductInfos)
	})
	r.Route("/recommendationRules", func(r chi.Router) {
		r.Use(AdminMiddleware.CheckJTW)
		r.Get("/", sm.GetRecommendationRules)
		r.Post("/", sm.PostRecommendationRule)
		r.Get("/{id}", sm.GetRecommendationRuleById)
		r.Put("/orderPriority", sm.PutOrderRecommendationRules)
		r.Put("/{id}/identity", sm.PutRecommendationRuleIdentity)
		r.Put("/{id}/general", sm.PutRecommendationRuleGeneral)
		r.Put("/{id}/validities", sm.PutRecommendationRuleValidities)
		r.Put("/{id}/displays", sm.PutRecommendationDisplay)
		r.Put("/{id}/salesGroups", sm.PutRecommendationSalesGroups)
		r.Put("/{id}/buyerTypes", sm.PutRecommendationBuyerTypes)
	})
	r.Route("/promotions", func(r chi.Router) {
		r.Use(AdminMiddleware.CheckJTW)
		r.Get("/", sm.GetPromotions)
		r.Get("/{id}", sm.GetPromotionById)
		r.Post("/", sm.PostPromotion)
		r.Post("/promotionalCodes/validate", sm.PostValidateCode)
		r.Put("/{id}/identity", sm.PutPromotionIdentity)
		r.Put("/{id}/general", sm.PutPromotionGeneral)
		r.Put("/{id}/validities", sm.PutPromotionValidities)
		r.Put("/{id}/advancedSettings", sm.PutPromotionAdvancedSettings)
		r.Put("/{id}/salesGroups", sm.PutPromotionSalesGroups)
		r.Put("/{id}/buyerTypes", sm.PutPromotionBuyerTypes)
		r.Put("/{id}/products", sm.PutPromotionProducts)
	})
	return r
}
