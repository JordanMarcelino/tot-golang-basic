package controller

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"tot_golang/internal/entity"
	"tot_golang/internal/models"
	"tot_golang/internal/models/converter"
	"tot_golang/internal/repository"
)

type UserController struct {
	Repository *repository.UserRepository
}

func NewUserController(repository *repository.UserRepository) *UserController {
	return &UserController{Repository: repository}
}

func (c *UserController) Create(ctx *fiber.Ctx) error {
	request := new(models.UserCreate)

	if err := ctx.BodyParser(request); err != nil {
		return fiber.ErrBadRequest
	}

	user := new(entity.User)
	user.Name = request.Name
	user.Division = request.Division

	if err := c.Repository.Create(user); err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(models.WebResponse{
		Info: models.Info{
			Code:    200,
			Message: "Success create user",
		},
		Data: converter.UserToResponse(user),
	})
}

func (c *UserController) Update(ctx *fiber.Ctx) error {
	request := new(models.UserUpdate)
	id, _ := strconv.Atoi(ctx.Params("userId"))
	request.ID = id

	if err := ctx.BodyParser(request); err != nil {
		return fiber.ErrBadRequest
	}

	user := new(entity.User)
	user.ID = request.ID
	user.Name = request.Name
	user.Division = request.Division

	if err := c.Repository.Update(user); err != nil {
		return fiber.ErrNotFound
	}

	return ctx.JSON(models.WebResponse{
		Info: models.Info{
			Code:    200,
			Message: "Success update user",
		},
		Data: converter.UserToResponse(user),
	})
}

func (c *UserController) Delete(ctx *fiber.Ctx) error {
	request := new(models.UserDelete)
	id, _ := strconv.Atoi(ctx.Params("userId"))
	request.ID = id

	user := new(entity.User)
	user.ID = request.ID

	if err := c.Repository.Delete(user); err != nil {
		return fiber.ErrNotFound
	}

	return ctx.JSON(models.WebResponse{
		Info: models.Info{
			Code:    200,
			Message: "Success delete user",
		},
		Data: nil,
	})
}

func (c *UserController) FindById(ctx *fiber.Ctx) error {
	request := new(models.UserGet)
	id, _ := strconv.Atoi(ctx.Params("userId"))
	request.ID = id

	user := new(entity.User)
	user.ID = request.ID

	if err := c.Repository.FindById(user); err != nil {
		return fiber.ErrNotFound
	}

	return ctx.JSON(models.WebResponse{
		Info: models.Info{
			Code:    200,
			Message: "Success get user",
		},
		Data: converter.UserToResponse(user),
	})
}

func (c *UserController) FindAll(ctx *fiber.Ctx) error {
	users, err := c.Repository.FindAll()

	if err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(models.WebResponse{
		Info: models.Info{
			Code:    200,
			Message: "Success get all user",
		},
		Data: converter.UsersToResponse(users),
	})
}
