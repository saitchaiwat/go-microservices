package handler

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"user-service/dto"
	"user-service/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserHandler {
	return UserHandler{userService: userService}
}

func (controller UserHandler) GetAllUser(c *fiber.Ctx) error {
	result, err := controller.userService.GetAllUser()

	if err != nil {
		c.Status(http.StatusNoContent).JSON(fiber.Map{"message": "No data available", "result": nil})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": err, "result": result})
}

func (controller UserHandler) CreateUser(c *fiber.Ctx) error {
	createUser := new(dto.CreateUser)

	if err := c.BodyParser(createUser); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err, "result": nil})
	}

	if err := controller.userService.CreateUser(createUser); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err, "result": nil})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Created user successfully", "result": createUser})
}
