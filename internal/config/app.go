package config

import (
	"github.com/go-playground/mold/v4"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"prbcare_be/internal/controller"
	"prbcare_be/internal/middleware"
	"prbcare_be/internal/repository"
	"prbcare_be/internal/route"
	"prbcare_be/internal/service"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Validate *validator.Validate
	Config   *viper.Viper
	Modifier *mold.Transformer
}

func Bootstrap(config *BootstrapConfig) {

	adminSuperRepository := repository.NewAdminSuperRepository()
	adminPuskesmasRepository := repository.NewAdminPuskesmasRepository()
	adminApotekRepository := repository.NewAdminApotekRepository()
	penggunaRepository := repository.NewPenggunaRepository()
	obatRepository := repository.NewObatRepository()
	pasienRepository := repository.NewPasienRepository()
	kontrolBalikRepository := repository.NewKontrolBalikRepository()

	adminSuperService := service.NewAdminSuperService(config.DB, adminSuperRepository, config.Validate, config.Config)
	adminPuskesmasService := service.NewAdminPuskesmasService(config.DB, adminPuskesmasRepository, config.Validate, config.Config)
	adminApotekService := service.NewAdminApotekService(config.DB, adminApotekRepository, config.Validate, config.Config)
	penggunaService := service.NewPenggunaService(config.DB, penggunaRepository, config.Validate, config.Config)
	obatService := service.NewObatService(config.DB, obatRepository, adminApotekRepository, config.Validate)
	pasienService := service.NewPasienService(config.DB, pasienRepository, adminPuskesmasRepository, penggunaRepository, kontrolBalikRepository, config.Validate)
	kontrolBalikService := service.NewKontrolBalikService(config.DB, kontrolBalikRepository, pasienRepository, config.Validate)

	adminSuperController := controller.NewAdminSuperController(adminSuperService)
	adminPuskesmasController := controller.NewAdminPuskesmasController(adminPuskesmasService, config.Modifier)
	adminApotekController := controller.NewAdminApotekController(adminApotekService, config.Modifier)
	penggunaController := controller.NewPenggunaController(penggunaService, config.Modifier)
	obatController := controller.NewObatController(obatService, config.Modifier)
	pasienController := controller.NewPasienController(pasienService, config.Modifier)
	kontrolBalikController := controller.NewKontrolBalikController(kontrolBalikService)

	adminSuperMiddleware := middleware.AdminSuperAuth(adminSuperService)
	adminPuskesmasMiddleware := middleware.AdminPuskesmasAuth(adminPuskesmasService)
	adminApotekMiddleware := middleware.AdminApotekAuth(adminApotekService)
	penggunaMiddleware := middleware.PenggunaAuth(penggunaService)
	adminSuperOrPuskesmasMiddleware := middleware.AdminSuperOrPuskesmasAuth(adminSuperService, adminPuskesmasService)
	adminSuperOrApotekMiddleware := middleware.AdminSuperOrApotekAuth(adminSuperService, adminApotekService)
	adminSuperOrPuskesmasOrApotekMiddleware := middleware.AdminSuperOrPuskesmasOrApotekAuth(adminSuperService, adminPuskesmasService, adminApotekService)
	adminSuperOrPuskesmasOrPengguna := middleware.AdminSuperOrPuskesmasOrPenggunaAuth(adminSuperService, adminPuskesmasService, penggunaService)

	route := route.RouteConfig{
		config.App,
		adminSuperController,
		adminSuperMiddleware,
		adminPuskesmasController,
		adminPuskesmasMiddleware,
		adminApotekController,
		adminApotekMiddleware,
		adminSuperOrPuskesmasMiddleware,
		penggunaController,
		penggunaMiddleware,
		adminSuperOrPuskesmasOrApotekMiddleware,
		adminSuperOrApotekMiddleware,
		obatController,
		pasienController,
		adminSuperOrPuskesmasOrPengguna,
		kontrolBalikController,
		config.Config,
	}
	route.Setup()

}
