package route

import (
	"github.com/gofiber/fiber/v3"
	"github.com/spf13/viper"
	"prbcare_be/internal/controller"
)

type RouteConfig struct {
	App                  *fiber.App
	AdminSuperController *controller.AdminSuperController
	AuthMiddleware       fiber.Handler
	Config               *viper.Viper
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/api/admin-super/login", c.AdminSuperController.Login)
}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use("/api/admin-super/current", c.AuthMiddleware)
	c.App.Patch("/api/admin-super/current/password", c.AdminSuperController.PasswordUpdate)
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}
