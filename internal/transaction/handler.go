package transaction

import (
	"errors"

	"github.com/aliyilmazdev/personal-budget-tracker/pkg/constant"
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
		return c.JSON(presenter.SuccessResponse(t, "transactions listed"))
	}
}

 func CreateTransactionHandler(service Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		t := new(Transaction)

		
		if err := c.BodyParser(t); err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(presenter.ErrorResponse(err))
		}

		if t.TransactionType != string(constant.TransactionTypeExpense){
			c.Status(fiber.StatusBadRequest)
			return c.JSON(presenter.ErrorResponse(errors.New("transaction type must be either expense or income")))
		}

		t.ID = uuid.New()
		service.Create(t)

		c.Status(fiber.StatusCreated)
		return c.JSON(presenter.SuccessResponse(t, "transaction created"))
	}	
 }

 func DeleteTransactionHandler(service Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		if err := service.Delete(id); err != nil {
			c.Status(fiber.StatusNotFound)
			return c.JSON(presenter.ErrorResponse(err))
		}

		c.Status(fiber.StatusOK)
		return c.JSON(presenter.SuccessResponse(nil,"transaction deleted"))
	}
 }