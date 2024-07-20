package controller

import (
	"github.com/gofiber/fiber/v3"
	"log"
	"prb_care_api/internal/constant"
	"prb_care_api/internal/middleware"
	"prb_care_api/internal/model"
	"prb_care_api/internal/service"
	"strconv"
)

type PengambilanObatController struct {
	PengambilanObatService *service.PengambilanObatService
}

func NewPengambilanObatController(pengambilanObatService *service.PengambilanObatService) *PengambilanObatController {
	return &PengambilanObatController{pengambilanObatService}
}

func (c *PengambilanObatController) Search(ctx fiber.Ctx) error {
	auth := middleware.GetAuth(ctx)
	request := new(model.PengambilanObatSearchRequest)
	if auth.Role == constant.RoleAdminPuskesmas {
		request.IdAdminPuskesmas = auth.ID
	} else if auth.Role == constant.RoleAdminApotek {
		request.IdAdminApotek = auth.ID
	} else if auth.Role == constant.RolePengguna {
		request.IdPengguna = auth.ID
	}
	request.Status = ctx.Query("status")
	response, err := c.PengambilanObatService.Search(ctx.Context(), request)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": response})
}

func (c *PengambilanObatController) Get(ctx fiber.Ctx) error {
	auth := middleware.GetAuth(ctx)
	request := new(model.PengambilanObatGetRequest)
	if auth.Role == constant.RoleAdminPuskesmas {
		request.IdAdminPuskesmas = auth.ID
	}
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}
	request.ID = id
	response, err := c.PengambilanObatService.Get(ctx.Context(), request)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": response})
}

func (c *PengambilanObatController) Create(ctx fiber.Ctx) error {
	auth := middleware.GetAuth(ctx)
	request := new(model.PengambilanObatCreateRequest)
	if err := ctx.Bind().JSON(request); err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}
	if auth.Role == constant.RoleAdminPuskesmas {
		request.IdAdminPuskesmas = auth.ID
	}
	if err := c.PengambilanObatService.Create(ctx.Context(), request); err != nil {
		log.Println(err.Error())
		return err
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"data": "Pengambilan obat berhasil dibuat"})
}

func (c *PengambilanObatController) Update(ctx fiber.Ctx) error {
	auth := middleware.GetAuth(ctx)
	request := new(model.PengambilanObatUpdateRequest)
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
	if auth.Role == constant.RoleAdminPuskesmas {
		request.IdAdminPuskesmas = auth.ID
	}
	if err := c.PengambilanObatService.Update(ctx.Context(), request); err != nil {
		log.Println(err.Error())
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": "Pengambilan obat berhasil diperbarui"})
}

func (c *PengambilanObatController) Delete(ctx fiber.Ctx) error {
	auth := middleware.GetAuth(ctx)
	request := new(model.PengambilanObatDeleteRequest)
	if auth.Role == constant.RoleAdminPuskesmas {
		request.IdAdminPuskesmas = auth.ID
	}
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}
	request.ID = id
	if err := c.PengambilanObatService.Delete(ctx.Context(), request); err != nil {
		log.Println(err.Error())
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": "Pengambilan obat berhasil dihapus"})
}

func (c *PengambilanObatController) Batal(ctx fiber.Ctx) error {
	auth := middleware.GetAuth(ctx)
	request := new(model.PengambilanObatBatalRequest)
	if auth.Role == constant.RoleAdminPuskesmas {
		request.IdAdminPuskesmas = auth.ID
	}
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}
	request.ID = id
	if err := c.PengambilanObatService.Batal(ctx.Context(), request); err != nil {
		log.Println(err.Error())
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": "Pengambilan obat berhasil ditandai batal"})
}

func (c *PengambilanObatController) Diambil(ctx fiber.Ctx) error {
	auth := middleware.GetAuth(ctx)
	request := new(model.PengambilanObatDiambilRequest)
	if auth.Role == constant.RoleAdminApotek {
		request.IdAdminApotek = auth.ID
	}
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}
	request.ID = id
	if err := c.PengambilanObatService.Diambil(ctx.Context(), request); err != nil {
		log.Println(err.Error())
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": "Pengambilan obat berhasil ditandai selesai"})
}
