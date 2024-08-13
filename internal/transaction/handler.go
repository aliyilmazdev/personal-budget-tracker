package transaction

import (
	"github.com/aliyilmazdev/personal-budget-tracker/pkg/presenter"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetTransactionsHandler(service Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		t, err := service.GetAll()

		if err != nil {
			c.Status(fiber.StatusNotFound)
			return c.JSON(presenter.ErrorResponse(err))
		}

		c.Status(fiber.StatusOK)
		return c.JSON(presenter.SuccessResponse(t))
	}
}

 func CreateTransactionHandler(service Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		t := new(Transaction)

		if err := c.BodyParser(t); err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(presenter.ErrorResponse(err))
		}

		t.ID = uuid.New()

		service.Create(t)

		c.Status(fiber.StatusCreated)
		return c.JSON(presenter.SuccessResponse(t))
	}	
 }