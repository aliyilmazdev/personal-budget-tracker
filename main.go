package main

import (
	"github.com/aliyilmazdev/personal-budget-tracker/database"
	"github.com/aliyilmazdev/personal-budget-tracker/internal/transaction"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New()

	database.ConnectDB()

	api := app.Group("/api", logger.New())

	transactionRepository := transaction.NewRepository(database.DB)
	transactionService := transaction.NewService(transactionRepository)
	transaction.SetupTransactionRoutes(api, transactionService)


	app.Listen(":3000")
	 
}