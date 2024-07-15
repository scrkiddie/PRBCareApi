package route

import (
	"github.com/gofiber/fiber/v3"
	"github.com/spf13/viper"
	"prbcare_be/internal/controller"
)

type RouteConfig struct {
	App                      *fiber.App
	AdminSuperController     *controller.AdminSuperController
	AdminSuperMiddleware     fiber.Handler
	AdminPuskesmasController *controller.AdminPuskesmasController
	AdminPuskesmasMiddleware fiber.Handler
	Config                   *viper.Viper
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/api/admin-super/login", c.AdminSuperController.Login)
	c.App.Post("/api/admin-puskesmas/login", c.AdminPuskesmasController.Login)
}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use("/api/admin-super/current", c.AdminSuperMiddleware)
	c.App.Patch("/api/admin-super/current/password", c.AdminSuperController.PasswordUpdate)

	c.App.Use("/api/admin-puskesmas/current", c.AdminPuskesmasMiddleware)
	c.App.Get("/api/admin-puskesmas/current", c.AdminPuskesmasController.Current)
	c.App.Patch("/api/admin-puskesmas/current", c.AdminPuskesmasController.ProfileUpdate)
	c.App.Patch("/api/admin-puskesmas/current/password", c.AdminPuskesmasController.PasswordUpdate)

	c.App.Use("/api/admin-puskesmas", c.AdminSuperMiddleware)
	c.App.Get("/api/admin-puskesmas", c.AdminPuskesmasController.List)
	c.App.Get("/api/admin-puskesmas/:id", c.AdminPuskesmasController.Get)
	c.App.Post("/api/admin-puskesmas", c.AdminPuskesmasController.Create)
	c.App.Patch("/api/admin-puskesmas/:id", c.AdminPuskesmasController.Update)
	c.App.Delete("/api/admin-puskesmas/:id", c.AdminPuskesmasController.Delete)
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}
