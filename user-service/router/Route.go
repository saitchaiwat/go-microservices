package router

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"user-service/handler"
	"user-service/repository"
	"user-service/service"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserController(userService)

	roleRepository := repository.NewRoleRepository(db)
	roleService := service.NewRoleService(roleRepository)
	roleHandler := handler.NewRoleController(roleService)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("/api")

	user := api.Group("/user")
	user.Get("/", userHandler.GetAllUser)
	user.Post("/", userHandler.CreateUser)

	role := api.Group("/role")
	role.Get("size/:size/page/:page", roleHandler.GetAllRole)
	role.Get("/:id", roleHandler.GetRoleById)
	role.Post("/", roleHandler.CreateRole)
	role.Patch("/:id", roleHandler.UpdateRole)
	role.Delete("/:id", roleHandler.DeleteRole)
}
