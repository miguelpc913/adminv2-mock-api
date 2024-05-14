package seeds

import (
	"fmt"

	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm"
)

func SeedProduct(db *gorm.DB) {
	products := []models.Product{
		{Name: "Visita guiada"},
		{Name: "Visita abierta"},
		{Name: "Programa escolar"},
		{Name: "Entrada"},
		{Name: "Guia reserva previa"},
	}

	result := db.Create(products)
	if result.Error != nil {
		fmt.Print("There has been an error in the seed of products")
	}
}
