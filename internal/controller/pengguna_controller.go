package controller

import (
	"github.com/go-playground/mold/v4"
	"github.com/gofiber/fiber/v3"
	"log"
	"prb_care_api/internal/middleware"
	"prb_care_api/internal/model"
	"prb_care_api/internal/service"
	"strconv"
)

type PenggunaController struct {
	PenggunaService *service.PenggunaService
	Modifier        *mold.Transformer
}

func NewPenggunaController(apotekService *service.PenggunaService, modifier *mold.Transformer) *PenggunaController {
	return &PenggunaController{apotekService, modifier}
}

func (c *PenggunaController) Login(ctx fiber.Ctx) error {
	request := new(model.PenggunaLoginRequest)
	if err := ctx.Bind().JSON(request); err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}
	response, err := c.PenggunaService.Login(ctx.Context(), request)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": response.Token})
}

func (c *PenggunaController) Current(ctx fiber.Ctx) error {
	auth := middleware.GetAuth(ctx)
	request := new(model.PenggunaGetRequest)
	request.ID = auth.ID
	response, err := c.PenggunaService.Current(ctx.Context(), request)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": response})
}

func (c *PenggunaController) CurrentProfileUpdate(ctx fiber.Ctx) error {
	auth := middleware.GetAuth(ctx)
	request := new(model.PenggunaProfileUpdateRequest)
	request.ID = auth.ID
	if err := ctx.Bind().JSON(request); err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}
	if err := c.Modifier.Struct(ctx.UserContext(), request); err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}
	if err := c.PenggunaService.CurrentProfileUpdate(ctx.UserContext(), request); err != nil {
		log.Println(err.Error())
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": "Pengguna berhasil diupdate"})
}

func (c *PenggunaController) CurrentTokenPerangkatUpdate(ctx fiber.Ctx) error {
	auth := middleware.GetAuth(ctx)
	request := new(model.PenggunaTokenPerangkatUpdateRequest)
	request.ID = auth.ID
	if err := ctx.Bind().JSON(request); err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}
	if err := c.PenggunaService.CurrentTokenPerangkatUpdate(ctx.UserContext(), request); err != nil {
		log.Println(err.Error())
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": "Token perangkat pengguna berhasil diupdate"})
}

func (c *PenggunaController) CurrentPasswordUpdate(ctx fiber.Ctx) error {
	auth := middleware.GetAuth(ctx)
	request := new(model.PenggunaPasswordUpdateRequest)
	request.ID = auth.ID
	if err := ctx.Bind().JSON(request); err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}
	if err := c.PenggunaService.CurrentPasswordUpdate(ctx.UserContext(), request); err != nil {
		log.Println(err.Error())
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": "Password berhasil diganti"})
}

func (c *PenggunaController) List(ctx fiber.Ctx) error {
	response, err := c.PenggunaService.List(ctx.Context())
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": response})
}

func (c *PenggunaController) Get(ctx fiber.Ctx) error {
	request := new(model.PenggunaGetRequest)
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}
	request.ID = id
	response, err := c.PenggunaService.Get(ctx.Context(), request)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": response})
}

func (c *PenggunaController) Create(ctx fiber.Ctx) error {
	request := new(model.PenggunaCreateRequest)

	if err := ctx.Bind().JSON(request); err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}

	if err := c.Modifier.Struct(ctx.UserContext(), request); err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	if err := c.PenggunaService.Create(ctx.UserContext(), request); err != nil {
		log.Println(err.Error())
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"data": "Pengguna berhasil dibuat"})
}

func (c *PenggunaController) Update(ctx fiber.Ctx) error {
	request := new(model.PenggunaUpdateRequest)
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

	if err := c.Modifier.Struct(ctx.UserContext(), request); err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	if err := c.PenggunaService.Update(ctx.UserContext(), request); err != nil {
		log.Println(err.Error())
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"data": "Pengguna berhasil diupdate"})
}

func (c *PenggunaController) Delete(ctx fiber.Ctx) error {
	request := new(model.PenggunaDeleteRequest)
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}
	request.ID = id

	if err := c.PenggunaService.Delete(ctx.UserContext(), request); err != nil {
		log.Println(err.Error())
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"data": "Pengguna berhasil dihapus"})
}
