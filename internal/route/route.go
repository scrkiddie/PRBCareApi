package route

import (
	"github.com/gofiber/fiber/v3"
	"github.com/spf13/viper"
	"prbcare_be/internal/controller"
)

type RouteConfig struct {
	App                  *fiber.App
	AdminSuperController *controller.AdminSuperController
	AdminSuperMiddleware fiber.Handler

	AdminPuskesmasController *controller.AdminPuskesmasController
	AdminPuskesmasMiddleware fiber.Handler

	AdminApotekController           *controller.AdminApotekController
	AdminApotekMiddleware           fiber.Handler
	AdminSuperOrPuskesmasMiddleware fiber.Handler

	Config *viper.Viper
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/api/admin-super/login", c.AdminSuperController.Login)
	c.App.Post("/api/admin-puskesmas/login", c.AdminPuskesmasController.Login)
	c.App.Post("/api/admin-apotek/login", c.AdminApotekController.Login)
}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use("/api/admin-super/current", c.AdminSuperMiddleware)
	c.App.Patch("/api/admin-super/current/password", c.AdminSuperController.PasswordUpdate)

	c.App.Use("/api/admin-puskesmas/current", c.AdminPuskesmasMiddleware)
	c.App.Get("/api/admin-puskesmas/current", c.AdminPuskesmasController.Current)
	c.App.Patch("/api/admin-puskesmas/current", c.AdminPuskesmasController.CurrentProfileUpdate)
	c.App.Patch("/api/admin-puskesmas/current/password", c.AdminPuskesmasController.CurrentPasswordUpdate)

	c.App.Use("/api/admin-puskesmas", c.AdminSuperMiddleware)
	c.App.Get("/api/admin-puskesmas", c.AdminPuskesmasController.List)
	c.App.Get("/api/admin-puskesmas/:id", c.AdminPuskesmasController.Get)
	c.App.Post("/api/admin-puskesmas", c.AdminPuskesmasController.Create)
	c.App.Patch("/api/admin-puskesmas/:id", c.AdminPuskesmasController.Update)
	c.App.Delete("/api/admin-puskesmas/:id", c.AdminPuskesmasController.Delete)

	c.App.Use("/api/admin-apotek/current", c.AdminApotekMiddleware)
	c.App.Get("/api/admin-apotek/current", c.AdminApotekController.Current)
	c.App.Patch("/api/admin-apotek/current", c.AdminApotekController.CurrentProfileUpdate)
	c.App.Patch("/api/admin-apotek/current/password", c.AdminApotekController.CurrentPasswordUpdate)

	c.App.Use("/api/admin-apotek", c.AdminSuperOrPuskesmasMiddleware)
	c.App.Get("/api/admin-apotek", c.AdminApotekController.List)

	c.App.Use("/api/admin-apotek", c.AdminSuperMiddleware)
	c.App.Get("/api/admin-apotek/:id", c.AdminApotekController.Get)
	c.App.Post("/api/admin-apotek", c.AdminApotekController.Create)
	c.App.Patch("/api/admin-apotek/:id", c.AdminApotekController.Update)
	c.App.Delete("/api/admin-apotek/:id", c.AdminApotekController.Delete)
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}
