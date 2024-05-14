package seeds

import (
	"fmt"

	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm"
)

func SeedProductInfoTypes(db *gorm.DB) {
	productInfoTypes := []models.ProductInfoType{
		{
			Id:           "atm_buyer_types_modal",
			Description:  "Modal colectivos en atm",
			DisplayOrder: 1,
		},
		{
			Id:           "buyer_types_top",
			Description:  "Información superior a colectivos",
			DisplayOrder: 2,
		},
		{
			Id:           "buyer_types_bottom",
			Description:  "Información inferior colectivos",
			DisplayOrder: 3,
		},
		{
			Id:           "calendar_bottom",
			Description:  "Información inferior a calendario",
			DisplayOrder: 4,
		},
		{
			Id:           "calendar_info",
			Description:  "Información en el calendario",
			DisplayOrder: 5,
		},
		{
			Id:           "calendar_top",
			Description:  "Información superior a calendario",
			DisplayOrder: 6,
		},
		{
			Id:           "disabled",
			Description:  "Desactivado",
			DisplayOrder: 7,
		},
		{
			Id:           "email_bottom",
			Description:  "Información inferior al email",
			DisplayOrder: 8,
		},
		{
			Id:           "email_info_introduction",
			Description:  "Información en email informativo de introduccion",
			DisplayOrder: 9,
		},
		{
			Id:           "email_info_reservation_bottom",
			Description:  "Información inferior en email de confirmación",
			DisplayOrder: 10,
		},
		{
			Id:           "email_info_reservation_top",
			Description:  "Información superior en email de confirmación",
			DisplayOrder: 11,
		},
		{
			Id:           "email_top",
			Description:  "Información superior al email",
			DisplayOrder: 12,
		},
		{
			Id:           "final_purchase_info",
			Description:  "Información al final de la compra",
			DisplayOrder: 13,
		},
		{
			Id:           "general_info",
			Description:  "Información del Producto (Ficha Web)",
			DisplayOrder: 14,
		},
		{
			Id:           "icon",
			Description:  "Mostrar icono con producto",
			DisplayOrder: 15,
		},
		{
			Id:           "pah_info",
			Description:  "Print At Home",
			DisplayOrder: 18,
		},
		{
			Id:           "passbook_conditions",
			Description:  "Condiciones de Passbook",
			DisplayOrder: 19,
		},
		{
			Id:           "popup_terms_addtocart",
			Description:  "Información de términos de producto al añadir al carro",
			DisplayOrder: 20,
		},
		{
			Id:           "popup_terms_product",
			Description:  "Información de términos de producto al entrar a comprar",
			DisplayOrder: 21,
		},
		{
			Id:           "promotions_bottom",
			Description:  "Información inferior a promociones",
			DisplayOrder: 22,
		},
		{
			Id:           "promotions_top",
			Description:  "Información superior a promociones",
			DisplayOrder: 23,
		},
		{
			Id:           "purchase_start_info",
			Description:  "Info antes del inicio de la venta",
			DisplayOrder: 24,
		},
		{
			Id:           "reservation_form_top",
			Description:  "Información superior a formulario reservas",
			DisplayOrder: 25,
		},
		{
			Id:           "reservation_form_bottom",
			Description:  "Información inferior a formulario reservas",
			DisplayOrder: 26,
		},
		{
			Id:           "services_bottom",
			Description:  "Información inferior a servicios",
			DisplayOrder: 27,
		},
		{
			Id:           "services_top",
			Description:  "Información superior a servicios",
			DisplayOrder: 28,
		},
		{
			Id:           "sessions_bottom",
			Description:  "Información inferior a sessiones",
			DisplayOrder: 29,
		},
		{
			Id:           "sessions_top",
			Description:  "Información superior a la selección de evento",
			DisplayOrder: 30,
		},
		{
			Id:           "ticket_info",
			Description:  "Información extra para ticket",
			DisplayOrder: 31,
		},
		{
			Id:           "ticket_form_top",
			Description:  "Información superior a formulario por ticket",
			DisplayOrder: 32,
		},
		{
			Id:           "ticket_form_bottom",
			Description:  "Información inferior a formulario por ticket",
			DisplayOrder: 33,
		},
		{
			Id:           "top_product_list",
			Description:  "top_product_list",
			DisplayOrder: 34,
		},
		{
			Id:           "under_product_list",
			Description:  "Información inferior a lista de productos",
			DisplayOrder: 35,
		},
		{
			Id:           "under_product",
			Description:  "Información inferior a producto",
			DisplayOrder: 36,
		},
		{
			Id:           "under_special_product_list",
			Description:  "Información inferior a lista de productos especiales",
			DisplayOrder: 37,
		},
	}

	result := db.Create(&productInfoTypes)
	if result.Error != nil {
		fmt.Println("Error occurred while seeding product info types:", result.Error)
	}
}
