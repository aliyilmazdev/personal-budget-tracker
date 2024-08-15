package transaction

import (
	"github.com/gofiber/fiber/v2"
)

func SetupTransactionRoutes(router fiber.Router, service Service) {
	t := router.Group("/transactions")
	
	t.Get("/", GetTransactionsHandler(service))
	t.Post("/", CreateTransactionHandler(service))
	t.Delete("/:id", DeleteTransactionHandler(service))
	t.Get("/:id", GetByIDTransactionHandler(service))
}