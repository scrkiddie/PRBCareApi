package controller

import (
	"github.com/go-playground/mold/v4"
	"github.com/gofiber/fiber/v3"
	"log"
	"prbcare_be/internal/constant"
	"prbcare_be/internal/middleware"
	"prbcare_be/internal/model"
	"prbcare_be/internal/service"
	"strconv"
)

type ObatController struct {
	ObatService *service.ObatService
	Modifier    *mold.Transformer
}

func NewObatController(obatService *service.ObatService, modifier *mold.Transformer) *ObatController {
	return &ObatController{obatService, modifier}
}
func (c *ObatController) List(ctx fiber.Ctx) error {
	auth := middleware.GetAuth(ctx)
	request := new(model.ObatListRequest)
	if auth.Role == constant.RoleAdminApotek {
		request.IdAdminApotek = auth.ID
	}
	response, err := c.ObatService.List(ctx.Context(), request)
	if err != nil {
		log.Println(err)
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": response})
}

func (c *ObatController) Get(ctx fiber.Ctx) error {
	auth := middleware.GetAuth(ctx)
	request := new(model.ObatGetRequest)
	if auth.Role == constant.RoleAdminApotek {
		request.IdAdminApotek = auth.ID
	}
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}
	request.ID = id
	response, err := c.ObatService.Get(ctx.Context(), request)
	if err != nil {
		log.Println(err)
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": response})
}

func (c *ObatController) Create(ctx fiber.Ctx) error {
	auth := middleware.GetAuth(ctx)
	request := new(model.ObatCreateRequest)

	if err := ctx.Bind().JSON(request); err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}

	if auth.Role == constant.RoleAdminApotek {
		request.IdAdminApotek = auth.ID
	}

	if err := c.Modifier.Struct(ctx.UserContext(), request); err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	if err := c.ObatService.Create(ctx.UserContext(), request); err != nil {
		log.Println(err.Error())
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"data": "Obat berhasil dibuat"})
}

func (c *ObatController) Update(ctx fiber.Ctx) error {
	auth := middleware.GetAuth(ctx)
	request := new(model.ObatUpdateRequest)
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}
	request.ID = id
	if err := ctx.Bind().JSON(request); err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}
	if auth.Role == constant.RoleAdminApotek {
		request.IdAdminApotek = auth.ID
		request.CurrentAdminApotek = true
	}

	if err := c.Modifier.Struct(ctx.UserContext(), request); err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	if err := c.ObatService.Update(ctx.UserContext(), request); err != nil {
		log.Println(err.Error())
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"data": "Obat berhasil diupdate"})
}

func (c *ObatController) Delete(ctx fiber.Ctx) error {
	auth := middleware.GetAuth(ctx)
	request := new(model.ObatDeleteRequest)
	if auth.Role == constant.RoleAdminApotek {
		request.IdAdminApotek = auth.ID
	}
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}
	request.ID = id

	if err := c.ObatService.Delete(ctx.UserContext(), request); err != nil {
		log.Println(err.Error())
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"data": "Obat berhasil dihapus"})
}
