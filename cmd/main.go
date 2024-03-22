package main

import (
	"github.com/gofiber/fiber/v2"
	"tot_golang/internal/controller"
	"tot_golang/internal/models"
	"tot_golang/internal/repository"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: NewErrorHandler(),
	})
	userRepository := repository.NewUserRepository()
	userController := controller.NewUserController(userRepository)

	api := app.Group("/api/v1")
	userApi := api.Group("/users")

	userApi.Get("", userController.FindAll)
	userApi.Post("", userController.Create)
	userApi.Get("/:userId", userController.FindById)
	userApi.Post("/:userId", userController.Update)
	userApi.Delete("/:userId", userController.Delete)

	app.Listen(":5050")
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError

		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		return ctx.Status(code).JSON(models.WebResponse{
			Info: models.Info{
				Code:    code,
				Message: err.Error(),
			},
			Data: nil,
		})
	}
}
