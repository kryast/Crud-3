package main

import (
	"log"

	"github.com/kryast/Crud-3.git/database"
	"github.com/kryast/Crud-3.git/handlers"
	"github.com/kryast/Crud-3.git/models"
	"github.com/kryast/Crud-3.git/repositories"
	"github.com/kryast/Crud-3.git/router"
	"github.com/kryast/Crud-3.git/services"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Gagal koneksi DB:", err)
	}
	db.AutoMigrate(&models.Customer{})

	repo := repositories.NewCustomerRepository(db)
	svc := services.NewCustomerService(repo)
	handler := handlers.NewCustomerHandler(svc)

	r := router.SetupRouter(handler)
	r.Run(":8080")
}
