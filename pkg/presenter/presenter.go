package presenter

import "github.com/gofiber/fiber/v2"

func SuccessResponse(response, msg interface{}) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data": response,
		"message": msg,
	}
}

func ErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"message": err.Error(),
	}
}