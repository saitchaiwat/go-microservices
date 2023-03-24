package handler

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"user-service/dto"
	"user-service/helper"
	"user-service/service"
)

type RoleHandler struct {
	roleService service.RoleService
}

func NewRoleController(roleService service.RoleService) RoleHandler {
	return RoleHandler{roleService: roleService}
}

func (controller RoleHandler) GetAllRole(c *fiber.Ctx) error {
	param := struct {
		Size int `params:"size"`
		Page int `params:"page"`
	}{}

	c.ParamsParser(&param)

	result, err := controller.roleService.GetAllRole(param.Size, param.Page)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "result": nil})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Get all role successfully", "result": result})
}

func (controller RoleHandler) GetRoleById(c *fiber.Ctx) error {
	param := struct {
		ID int `params:"id"`
	}{}

	c.ParamsParser(&param)

	result, err := controller.roleService.GetRoleByID(param.ID)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "result": nil})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Get role successfully", "result": result})
}

func (controller RoleHandler) CreateRole(c *fiber.Ctx) error {
	createRole := new(dto.CreateRole)

	if err := c.BodyParser(createRole); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error(), "result": nil})
	}

	if err := helper.ValidateStruct(createRole); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err, "result": nil})
	}

	if err := controller.roleService.CreateRole(createRole); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "result": nil})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Created role successfully", "result": createRole})
}

func (controller RoleHandler) UpdateRole(c *fiber.Ctx) error {
	param := struct {
		ID int `params:"id"`
	}{}

	c.ParamsParser(&param)

	updateRole := new(dto.UpdateRole)

	c.BodyParser(updateRole)

	if err := controller.roleService.UpdateRole(param.ID, updateRole); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "result": nil})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Updated role successfully", "result": nil})
}

func (controller RoleHandler) DeleteRole(c *fiber.Ctx) error {
	param := struct {
		ID int `params:"id"`
	}{}

	c.ParamsParser(&param)

	if err := controller.roleService.DeleteRole(param.ID); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "result": nil})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Deleted role successfully", "result": nil})
}
