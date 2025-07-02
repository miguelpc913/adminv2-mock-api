package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/tiqueteo/adminv2-mock-api/api/helpers"
	AdminMiddleware "github.com/tiqueteo/adminv2-mock-api/api/middleware"
	"github.com/tiqueteo/adminv2-mock-api/api/services"

	dbHelpers "github.com/tiqueteo/adminv2-mock-api/db/utils"
)

// #################################################################
// based on this example
// https://github.com/go-chi/cors/blob/master/_example/main.go#L79
// #################################################################
func AllowOriginFunc(r *http.Request, origin string) bool {
	allowedOrigins := []string{"http://localhost:8000", "https://admin-micro-qa.clorian.com"}
	for _, o := range allowedOrigins {
		if origin == o {
			return true
		}
	}

	// Get the client's IP address
	clientIP := r.RemoteAddr

	// Print the log in red color
	red := "\033[31m"
	reset := "\033[0m"
	log.Printf("%sDisallowed origin: %s, Client IP: %s%s", red, origin, clientIP, reset)

	return false
}

func Init() *chi.Mux {
	db, err := dbHelpers.InitDB(true)
	if err != nil {
		panic("There has been an error connecting to the datababase")
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	sm := services.NewServiceManager(db)

	r.Use(cors.Handler(cors.Options{
		AllowOriginFunc:  AllowOriginFunc,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, Go!")
	})

	r.Post("/login", sm.Login)
	r.Route("/products", func(r chi.Router) {
		r.Use(AdminMiddleware.CheckJTW)
		r.Use(AdminMiddleware.RecoverMiddleware)
		r.Get("/", sm.GetProducts)
	})
	r.Route("/venueCapacities", func(r chi.Router) {
		r.Use(AdminMiddleware.CheckJTW)
		r.Use(AdminMiddleware.RecoverMiddleware)
		r.Get("/", sm.GetVenueCapacities)
	})
	r.Route("/venues", func(r chi.Router) {
		r.Use(AdminMiddleware.CheckJTW)
		r.Use(AdminMiddleware.RecoverMiddleware)
		r.Get("/", sm.GetVenues)
	})
	r.Route("/salesGroups", func(r chi.Router) {
		r.Use(AdminMiddleware.CheckJTW)
		r.Use(AdminMiddleware.RecoverMiddleware)
		r.Get("/", sm.GetSalesGroups)
	})
	r.Route("/buyerTypes", func(r chi.Router) {
		r.Use(AdminMiddleware.CheckJTW)
		r.Use(AdminMiddleware.RecoverMiddleware)
		r.Get("/", sm.GetBuyerTypes)
		r.Get("/{id}", sm.GetBuyerTypeById)
		r.Get("/verifierAlerts", sm.GetVAB)
		r.Get("/verifierAlerts/{id}", sm.GetVABById)
		r.Put("/verifierAlerts/{id}", sm.PutVAB)
		r.Post("/verifierAlerts", sm.PostVAB)
		r.Delete("/verifierAlerts", sm.DeleteVAB)
	})
	r.Route("/extras", func(r chi.Router) {
		r.Use(AdminMiddleware.CheckJTW)
		r.Use(AdminMiddleware.RecoverMiddleware)
		r.Get("/", sm.GetExtras)
	})
	r.Route("/productTags", func(r chi.Router) {
		r.Use(AdminMiddleware.CheckJTW)
		r.Use(AdminMiddleware.RecoverMiddleware)
		r.Get("/", sm.GetProductTags)
	})
	r.Route("/verifiers", func(r chi.Router) {
		r.Use(AdminMiddleware.CheckJTW)
		r.Use(AdminMiddleware.RecoverMiddleware)
		r.Get("/", sm.GetVerifiers)
	})
	r.Route("/productInfos", func(r chi.Router) {
		r.Use(AdminMiddleware.CheckJTW)
		r.Use(AdminMiddleware.RecoverMiddleware)
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
		r.Use(AdminMiddleware.RecoverMiddleware)
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
		r.Use(AdminMiddleware.RecoverMiddleware)
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
		r.Get("/verifierAlerts", sm.GetVAP)
		r.Get("/verifierAlerts/{id}", sm.GetVAPById)
		r.Put("/verifierAlerts/{id}", sm.PutVAP)
		r.Post("/verifierAlerts", sm.PostVAP)
		r.Delete("/verifierAlerts", sm.DeleteVAP)
	})
	r.Route("/affiliateItems", func(r chi.Router) {
		r.Use(AdminMiddleware.CheckJTW)
		r.Use(AdminMiddleware.RecoverMiddleware)
		r.Get("/", sm.GetAffiliateItems)
		r.Get("/{id}", sm.GetAffiliateItemById)
		r.Post("/", sm.PostAffiliateItem)
		r.Put("/{id}", sm.PutAffiliateItem)
	})
	r.Route("/affiliateAgreements", func(r chi.Router) {
		r.Use(AdminMiddleware.CheckJTW)
		r.Use(AdminMiddleware.RecoverMiddleware)
		r.Get("/", sm.GetAffiliateAgreement)
		r.Get("/{id}", sm.GetAffiliateAgreementById)
		r.Post("/", sm.PostAffiliateAgreement)
		r.Put("/orderPriority", sm.PutOrderAffiliateAgreement)
		r.Put("/{id}/general", sm.PutAffiliateAgreementGeneral)
		r.Put("/{id}/validities", sm.PutAffiliateAgreementValities)
		r.Put("/{id}/products", sm.PutAffiliateAgreementProducts)
		r.Put("/{id}/buyerTypes", sm.PutAffiliateAgreementBuyerTypes)
	})
	r.Route("/affiliates", func(r chi.Router) {
		r.Use(AdminMiddleware.CheckJTW)
		r.Use(AdminMiddleware.RecoverMiddleware)
		r.Get("/", sm.GetAffiliates)
		r.Get("/{id}", sm.GetAffiliateById)
		r.Put("/{id}", sm.PutAffiliate)
		r.Put("/{id}/affiliateAgreements", sm.PutAffiliateAgreements)
	})
	r.Route("/boxOffices", func(r chi.Router) {
		r.Use(AdminMiddleware.CheckJTW)
		r.Use(AdminMiddleware.RecoverMiddleware)
		r.Get("/", sm.GetBO)
		r.Post("/", sm.PostBO)
		r.Get("/{id}", sm.GetBOById)
		r.Put("/{id}/basicConfigurations", sm.PutBOBasicConfigurations)
		r.Put("/{id}/cashCount", sm.PutBOCashCount)
		r.Put("/{id}/presentations", sm.PutBOPresentations)
		r.Put("/{id}/functionalities", sm.PutBOFunctionalities)
		r.Put("/{id}/languages", sm.PutBOLanguages)
		r.Put("/{id}/printSettings", sm.PutBOPrintSettings)
		r.Put("/{id}/paymentSettings", sm.PutBOPaymentSettings)
		r.Put("/{id}/validations", sm.PutBOValidations)
		r.Put("/{id}/advancedSettings", sm.PutBOAdvancedSettings)
		r.Put("/{id}/salesGroups", sm.PutBOSalesGroups)
		r.Put("/{id}/products", sm.PutBOProducts)
	})

	r.Route("/pricings", func(r chi.Router) {
		r.Use(AdminMiddleware.CheckJTW)
		r.Use(AdminMiddleware.RecoverMiddleware)
		r.Get("/", sm.GetPricings)
		r.Post("/", sm.PostBasePricing)
		r.Get("/{id}", sm.GetPricingById)
		r.Post("/{id}/advancedPricings", sm.PostSpecficPricing)
		r.Put("/{id}/configurations", sm.PutSpecificPricingsConfiguration)
		r.Put("/{id}/advancedPricings/priorities", sm.PutPricingsPriorities)
		r.Put("/tariffs", sm.PutPricingsTariffs)

	})

	r.Route("/buyerTypeRules", func(r chi.Router) {
		r.Use(AdminMiddleware.CheckJTW)
		r.Use(AdminMiddleware.RecoverMiddleware)
		r.Get("/", sm.GetBuyerTypesRules)
		r.Get("/{id}", sm.GetBuyerTypeRulesId)
		r.Post("/", sm.PostBuyerTypeRules)
		r.Put("/{id}/identity", sm.PutBuyerTypeRulesIdentity)
		r.Put("/{id}/configurations", sm.PutBuyerTypeRuleConfiguration)
		r.Put("/{id}/products", sm.PutBuyerTypeRulesProductSet)
	})

	r.Route("/appUsers", func(r chi.Router) {
		r.Use(AdminMiddleware.CheckJTW)
		r.Use(AdminMiddleware.RecoverMiddleware)
		r.Get("/", sm.GetAppUsers)
		r.Get("/{id}", sm.GetAppUserById)
		r.Post("/", sm.PostAppUser)
		r.Post("/validate", sm.PostValidateUser)
		r.Put("/{id}/identity", sm.PutAppUserIdentity)
		r.Put("/{id}/reports", sm.PutReportSet)
		r.Put("/{id}/pointOfSales", sm.PutPointOfSaleSet)
	})

	r.Route("/pointOfSales", func(r chi.Router) {
		r.Use(AdminMiddleware.CheckJTW)
		r.Use(AdminMiddleware.RecoverMiddleware)
		r.Get("/", sm.GetPO)
	})

	r.Route("/reports", func(r chi.Router) {
		r.Use(AdminMiddleware.CheckJTW)
		r.Use(AdminMiddleware.RecoverMiddleware)
		r.Get("/", sm.GetReports)
	})

	r.Route("/bulkActions", func(r chi.Router) {
		r.Use(AdminMiddleware.CheckJTW)
		r.Use(AdminMiddleware.RecoverMiddleware)
		r.Post("/validate", sm.PostBulkActionsValidate)
		r.Post("/", sm.PostBulkActionsExecute)
	})

	r.Route("/restartDb", func(r chi.Router) {
		r.Use(AdminMiddleware.CheckJTW)
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			var dbname = os.Getenv("MYSQL_DATABASE")
			tx := db.Exec("DROP DATABASE " + dbname + ";")
			if tx.Error != nil {
				helpers.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": tx.Error.Error()})
				return
			}
			tx = db.Exec("CREATE DATABASE " + dbname + ";")
			if tx.Error != nil {
				helpers.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": tx.Error.Error()})
				return
			}
			newDb, err := dbHelpers.InitDB(true)
			if err != nil {
				helpers.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
				return
			}
			*db = *newDb
			*sm = *services.NewServiceManager(db)
			helpers.WriteJSON(w, http.StatusOK, map[string]string{"success": "Db has been restarted"})
		})
	})

	return r
}
