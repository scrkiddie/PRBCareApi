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

	AdminApotekController *controller.AdminApotekController
	AdminApotekMiddleware fiber.Handler

	AdminSuperOrPuskesmasMiddleware fiber.Handler

	PenggunaController *controller.PenggunaController
	PenggunaMiddleware fiber.Handler

	AdminSuperOrPuskesmasOrApotekMiddleware fiber.Handler
	AdminSuperOrApotekMiddleware            fiber.Handler
	ObatController                          *controller.ObatController

	PasienController                          *controller.PasienController
	AdminSuperOrPuskesmasOrPenggunaMiddleware fiber.Handler

	KontrolBalikController *controller.KontrolBalikController

	AdminSuperOrPuskesmasOrApotekOrPenggunaAuth fiber.Handler
	PengambilanObatController                   *controller.PengambilanObatController

	Config *viper.Viper
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/api/admin-super/login", c.AdminSuperController.Login)
	c.App.Post("/api/admin-puskesmas/login", c.AdminPuskesmasController.Login)
	c.App.Post("/api/admin-apotek/login", c.AdminApotekController.Login)
	c.App.Post("/api/pengguna/login", c.PenggunaController.Login)
	c.App.Post("/api/pengguna", c.PenggunaController.Create)
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

	c.App.Use("/api/pengguna/current", c.PenggunaMiddleware)
	c.App.Get("/api/pengguna/current", c.PenggunaController.Current)
	c.App.Patch("/api/pengguna/current", c.PenggunaController.CurrentProfileUpdate)
	c.App.Patch("/api/pengguna/current/password", c.PenggunaController.CurrentPasswordUpdate)
	c.App.Patch("/api/pengguna/current/perangkat", c.PenggunaController.CurrentTokenPerangkatUpdate)

	c.App.Use("/api/pengguna", c.AdminSuperOrPuskesmasMiddleware)
	c.App.Get("/api/pengguna", c.PenggunaController.List)

	c.App.Use("/api/pengguna", c.AdminSuperMiddleware)
	c.App.Get("/api/pengguna/:id", c.PenggunaController.Get)
	c.App.Post("/api/pengguna", c.PenggunaController.Create)
	c.App.Patch("/api/pengguna/:id", c.PenggunaController.Update)
	c.App.Delete("/api/pengguna/:id", c.PenggunaController.Delete)

	c.App.Use("/api/obat", c.AdminSuperOrPuskesmasOrApotekMiddleware)
	c.App.Get("/api/obat", c.ObatController.List)

	c.App.Use("/api/obat", c.AdminSuperOrApotekMiddleware)
	c.App.Get("/api/obat/:id", c.ObatController.Get)
	c.App.Post("/api/obat", c.ObatController.Create)
	c.App.Patch("/api/obat/:id", c.ObatController.Update)
	c.App.Delete("/api/obat/:id", c.ObatController.Delete)

	c.App.Use("/api/pasien", c.AdminSuperOrPuskesmasOrPenggunaMiddleware)
	c.App.Get("/api/pasien", c.PasienController.Search)

	c.App.Use("/api/pasien", c.AdminSuperOrPuskesmasMiddleware)
	c.App.Get("/api/pasien/:id", c.PasienController.Get)
	c.App.Post("/api/pasien", c.PasienController.Create)
	c.App.Patch("/api/pasien/:id", c.PasienController.Update)
	c.App.Delete("/api/pasien/:id", c.PasienController.Delete)
	c.App.Patch("/api/pasien/:id/selesai", c.PasienController.Selesai)

	c.App.Use("/api/kontrol-balik", c.AdminSuperOrPuskesmasOrPenggunaMiddleware)
	c.App.Get("/api/kontrol-balik", c.KontrolBalikController.Search)

	c.App.Use("/api/kontrol-balik", c.AdminSuperOrPuskesmasMiddleware)
	c.App.Get("/api/kontrol-balik/:id", c.KontrolBalikController.Get)
	c.App.Post("/api/kontrol-balik", c.KontrolBalikController.Create)
	c.App.Patch("/api/kontrol-balik/:id", c.KontrolBalikController.Update)
	c.App.Delete("/api/kontrol-balik/:id", c.KontrolBalikController.Delete)
	c.App.Patch("/api/kontrol-balik/:id/selesai", c.KontrolBalikController.Selesai)
	c.App.Patch("/api/kontrol-balik/:id/batal", c.KontrolBalikController.Batal)

	c.App.Use("/api/pengambilan-obat", c.AdminSuperOrPuskesmasOrApotekOrPenggunaAuth)
	c.App.Get("/api/pengambilan-obat", c.PengambilanObatController.Search)

	c.App.Use("/api/pengambilan-obat", c.AdminSuperOrPuskesmasMiddleware)
	c.App.Get("/api/pengambilan-obat/:id", c.PengambilanObatController.Get)
	c.App.Post("/api/pengambilan-obat", c.PengambilanObatController.Create)
	c.App.Patch("/api/pengambilan-obat/:id", c.PengambilanObatController.Update)
	c.App.Delete("/api/pengambilan-obat/:id", c.PengambilanObatController.Delete)
	c.App.Patch("/api/pengambilan-obat/:id/batal", c.PengambilanObatController.Batal)

	c.App.Use("/api/pengambilan-obat/:id/diambil", c.AdminSuperOrApotekMiddleware)
	c.App.Patch("/api/pengambilan-obat/:id/diambil", c.PengambilanObatController.Diambil)

}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}
