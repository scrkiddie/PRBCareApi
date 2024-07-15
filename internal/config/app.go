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

	adminSuperService := service.NewAdminSuperService(config.DB, adminSuperRepository, config.Validate, config.Config)
	adminPuskesmasService := service.NewAdminPuskesmasService(config.DB, adminPuskesmasRepository, config.Validate, config.Config)

	adminSuperController := controller.NewAdminSuperController(adminSuperService)
	adminPuskesmasController := controller.NewAdminPuskesmasController(adminPuskesmasService, config.Modifier)

	adminSuperMiddleware := middleware.AdminSuperAuth(adminSuperService)
	adminPuskesmasMiddleware := middleware.AdminPuskesmasAuth(adminPuskesmasService)

	route := route.RouteConfig{
		config.App,
		adminSuperController,
		adminSuperMiddleware,
		adminPuskesmasController,
		adminPuskesmasMiddleware,
		config.Config,
	}
	route.Setup()

}
