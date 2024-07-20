package middleware

import (
	"github.com/gofiber/fiber/v3"
	"log"
	"prb_care_api/internal/model"
	"prb_care_api/internal/service"
	"strings"
)

func AdminSuperAuth(adminSuperService *service.AdminSuperService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		tokenWithBearer := ctx.Get("Authorization")
		if tokenWithBearer == "" {
			return fiber.ErrUnauthorized
		}

		parts := strings.Split(tokenWithBearer, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return fiber.ErrUnauthorized
		}
		token := parts[1]

		request := &model.VerifyAdminSuperRequest{Token: token}
		log.Printf("Token: %s", request.Token)

		auth, err := adminSuperService.Verify(ctx.UserContext(), request)
		if err != nil {
			log.Printf(err.Error())
			return fiber.ErrUnauthorized
		}

		log.Printf("user: %+v", auth.ID)
		ctx.Locals("auth", auth)
		return ctx.Next()
	}
}

func AdminPuskesmasAuth(adminPuskesmasService *service.AdminPuskesmasService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		tokenWithBearer := ctx.Get("Authorization")
		if tokenWithBearer == "" {
			return fiber.ErrUnauthorized
		}

		parts := strings.Split(tokenWithBearer, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return fiber.ErrUnauthorized
		}
		token := parts[1]

		request := &model.VerifyAdminPuskesmasRequest{Token: token}
		log.Printf("Token: %s", request.Token)

		auth, err := adminPuskesmasService.Verify(ctx.UserContext(), request)
		if err != nil {
			log.Printf(err.Error())
			return fiber.ErrUnauthorized
		}

		log.Printf("user: %+v", auth.ID)
		ctx.Locals("auth", auth)
		return ctx.Next()
	}
}

func AdminApotekAuth(adminApotekService *service.AdminApotekService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		tokenWithBearer := ctx.Get("Authorization")
		if tokenWithBearer == "" {
			return fiber.ErrUnauthorized
		}

		parts := strings.Split(tokenWithBearer, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return fiber.ErrUnauthorized
		}
		token := parts[1]
		log.Printf("Token: %s", token)

		request := &model.VerifyAdminApotekRequest{Token: token}
		auth, err := adminApotekService.Verify(ctx.UserContext(), request)
		if err != nil {
			log.Printf(err.Error())
			return fiber.ErrUnauthorized
		}

		log.Printf("user: %+v", auth.ID)
		ctx.Locals("auth", auth)
		return ctx.Next()
	}
}

func AdminSuperOrPuskesmasAuth(adminSuperService *service.AdminSuperService, adminPuskesmasService *service.AdminPuskesmasService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		tokenWithBearer := ctx.Get("Authorization")
		if tokenWithBearer == "" {
			return fiber.ErrUnauthorized
		}

		parts := strings.Split(tokenWithBearer, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return fiber.ErrUnauthorized
		}
		token := parts[1]
		log.Printf("Token: %s", token)

		requestSuper := &model.VerifyAdminSuperRequest{Token: token}
		authSuper, errSuper := adminSuperService.Verify(ctx.UserContext(), requestSuper)
		if errSuper == nil {
			log.Printf("Authenticated as AdminSuper: %+v", authSuper.ID)
			ctx.Locals("auth", authSuper)
			return ctx.Next()
		}

		requestPuskesmas := &model.VerifyAdminPuskesmasRequest{Token: token}
		authPuskesmas, errPuskesmas := adminPuskesmasService.Verify(ctx.UserContext(), requestPuskesmas)
		if errPuskesmas == nil {
			log.Printf("Authenticated as AdminPuskesmas: %+v", authPuskesmas.ID)
			ctx.Locals("auth", authPuskesmas)
			return ctx.Next()
		}

		log.Printf("Unauthorized: %v, %v", errSuper, errPuskesmas)
		return fiber.ErrUnauthorized
	}
}

func PenggunaAuth(adminApotekService *service.PenggunaService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		tokenWithBearer := ctx.Get("Authorization")
		if tokenWithBearer == "" {
			return fiber.ErrUnauthorized
		}

		parts := strings.Split(tokenWithBearer, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return fiber.ErrUnauthorized
		}
		token := parts[1]
		log.Printf("Token: %s", token)

		request := &model.VerifyPenggunaRequest{Token: token}
		auth, err := adminApotekService.Verify(ctx.UserContext(), request)
		if err != nil {
			log.Printf(err.Error())
			return fiber.ErrUnauthorized
		}

		log.Printf("user: %+v", auth.ID)
		ctx.Locals("auth", auth)
		return ctx.Next()
	}
}
func AdminSuperOrPuskesmasOrApotekAuth(adminSuperService *service.AdminSuperService, adminPuskesmasService *service.AdminPuskesmasService, adminApotekService *service.AdminApotekService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		tokenWithBearer := ctx.Get("Authorization")
		if tokenWithBearer == "" {
			return fiber.ErrUnauthorized
		}

		parts := strings.Split(tokenWithBearer, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return fiber.ErrUnauthorized
		}

		token := parts[1]
		log.Printf("Token: %s", token)

		requestSuper := &model.VerifyAdminSuperRequest{Token: token}
		authSuper, errSuper := adminSuperService.Verify(ctx.UserContext(), requestSuper)
		if errSuper == nil {
			log.Printf("Authenticated as AdminSuper: %+v", authSuper.ID)
			ctx.Locals("auth", authSuper)
			return ctx.Next()
		}

		requestPuskesmas := &model.VerifyAdminPuskesmasRequest{Token: token}
		authPuskesmas, errPuskesmas := adminPuskesmasService.Verify(ctx.UserContext(), requestPuskesmas)
		if errPuskesmas == nil {
			log.Printf("Authenticated as AdminPuskesmas: %+v", authPuskesmas.ID)
			ctx.Locals("auth", authPuskesmas)
			return ctx.Next()
		}

		requestApotek := &model.VerifyAdminApotekRequest{Token: token}
		authApotek, errApotek := adminApotekService.Verify(ctx.UserContext(), requestApotek)
		if errApotek == nil {
			log.Printf("Authenticated as AdminApotek: %+v", authApotek.ID)
			ctx.Locals("auth", authApotek)
			return ctx.Next()
		}

		log.Printf("Unauthorized: AdminSuper error: %v, Puskesmas error: %v, Apotek error: %v", errSuper, errPuskesmas, errApotek)
		return fiber.ErrUnauthorized
	}
}
func AdminSuperOrApotekAuth(adminSuperService *service.AdminSuperService, adminApotekService *service.AdminApotekService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		tokenWithBearer := ctx.Get("Authorization")
		if tokenWithBearer == "" {
			return fiber.ErrUnauthorized
		}

		parts := strings.Split(tokenWithBearer, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return fiber.ErrUnauthorized
		}

		token := parts[1]
		log.Printf("Token: %s", token)

		requestSuper := &model.VerifyAdminSuperRequest{Token: token}
		authSuper, errSuper := adminSuperService.Verify(ctx.UserContext(), requestSuper)
		if errSuper == nil {
			log.Printf("Authenticated as AdminSuper: %+v", authSuper.ID)
			ctx.Locals("auth", authSuper)
			return ctx.Next()
		}

		requestApotek := &model.VerifyAdminApotekRequest{Token: token}
		authApotek, errApotek := adminApotekService.Verify(ctx.UserContext(), requestApotek)
		if errApotek == nil {
			log.Printf("Authenticated as AdminApotek: %+v", authApotek.ID)
			ctx.Locals("auth", authApotek)
			return ctx.Next()
		}

		log.Printf("Unauthorized: AdminSuper error: %v, Apotek error: %v", errSuper, errApotek)
		return fiber.ErrUnauthorized
	}
}
func AdminSuperOrPuskesmasOrPenggunaAuth(adminSuperService *service.AdminSuperService, adminPuskesmasService *service.AdminPuskesmasService, penggunaService *service.PenggunaService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		tokenWithBearer := ctx.Get("Authorization")
		if tokenWithBearer == "" {
			return fiber.ErrUnauthorized
		}

		parts := strings.Split(tokenWithBearer, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return fiber.ErrUnauthorized
		}

		token := parts[1]
		log.Printf("Token: %s", token)

		requestSuper := &model.VerifyAdminSuperRequest{Token: token}
		authSuper, errSuper := adminSuperService.Verify(ctx.UserContext(), requestSuper)
		if errSuper == nil {
			log.Printf("Authenticated as AdminSuper: %+v", authSuper.ID)
			ctx.Locals("auth", authSuper)
			return ctx.Next()
		}

		requestPuskesmas := &model.VerifyAdminPuskesmasRequest{Token: token}
		authPuskesmas, errPuskesmas := adminPuskesmasService.Verify(ctx.UserContext(), requestPuskesmas)
		if errPuskesmas == nil {
			log.Printf("Authenticated as AdminPuskesmas: %+v", authPuskesmas.ID)
			ctx.Locals("auth", authPuskesmas)
			return ctx.Next()
		}

		requestPengguna := &model.VerifyPenggunaRequest{Token: token}
		authPengguna, errPengguna := penggunaService.Verify(ctx.UserContext(), requestPengguna)
		if errPengguna == nil {
			log.Printf("Authenticated as Pengguna: %+v", authPengguna.ID)
			ctx.Locals("auth", authPengguna)
			return ctx.Next()
		}

		log.Printf("Unauthorized: AdminSuper error: %v, Puskesmas error: %v, Pengguna error: %v", errSuper, errPuskesmas, errPengguna)
		return fiber.ErrUnauthorized
	}
}
func AdminSuperOrPuskesmasOrApotekOrPenggunaAuth(adminSuperService *service.AdminSuperService, adminPuskesmasService *service.AdminPuskesmasService, adminApotekService *service.AdminApotekService, penggunaService *service.PenggunaService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		tokenWithBearer := ctx.Get("Authorization")
		if tokenWithBearer == "" {
			return fiber.ErrUnauthorized
		}

		parts := strings.Split(tokenWithBearer, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return fiber.ErrUnauthorized
		}
		token := parts[1]
		log.Printf("Token: %s", token)

		requestSuper := &model.VerifyAdminSuperRequest{Token: token}
		authSuper, errSuper := adminSuperService.Verify(ctx.UserContext(), requestSuper)
		if errSuper == nil {
			log.Printf("Authenticated as AdminSuper: %+v", authSuper.ID)
			ctx.Locals("auth", authSuper)
			return ctx.Next()
		}

		requestPuskesmas := &model.VerifyAdminPuskesmasRequest{Token: token}
		authPuskesmas, errPuskesmas := adminPuskesmasService.Verify(ctx.UserContext(), requestPuskesmas)
		if errPuskesmas == nil {
			log.Printf("Authenticated as AdminPuskesmas: %+v", authPuskesmas.ID)
			ctx.Locals("auth", authPuskesmas)
			return ctx.Next()
		}

		requestApotek := &model.VerifyAdminApotekRequest{Token: token}
		authApotek, errApotek := adminApotekService.Verify(ctx.UserContext(), requestApotek)
		if errApotek == nil {
			log.Printf("Authenticated as AdminApotek: %+v", authApotek.ID)
			ctx.Locals("auth", authApotek)
			return ctx.Next()
		}

		requestPengguna := &model.VerifyPenggunaRequest{Token: token}
		authPengguna, errPengguna := penggunaService.Verify(ctx.UserContext(), requestPengguna)
		if errPengguna == nil {
			log.Printf("Authenticated as Pengguna: %+v", authPengguna.ID)
			ctx.Locals("auth", authPengguna)
			return ctx.Next()
		}

		log.Printf("Unauthorized: AdminSuper error: %v, Puskesmas error: %v, Apotek error: %v, Pengguna error: %v", errSuper, errPuskesmas, errApotek, errPengguna)
		return fiber.ErrUnauthorized
	}
}
func GetAuth(ctx fiber.Ctx) *model.Auth {
	return ctx.Locals("auth").(*model.Auth)
}
