package transaction

import (
	"github.com/gofiber/fiber/v2"
)

func SetupTransactionRoutes(router fiber.Router, service Service) {
	t := router.Group("/transactions")
	
	t.Get("/", GetTransactionsHandler(service))
	t.Post("/", CreateTransactionHandler(service))
	//t.Get("/:transactionId", func(c *fiber.Ctx) error {})
	//t.Delete("/:transactionId", func(c *fiber.Ctx) error {})
}