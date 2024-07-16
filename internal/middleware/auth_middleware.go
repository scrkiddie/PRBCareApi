package middleware

import (
	"github.com/gofiber/fiber/v3"
	"log"
	"prbcare_be/internal/model"
	"prbcare_be/internal/service"
)

func AdminSuperAuth(adminSuperService *service.AdminSuperService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		request := &model.VerifyAdminSuperRequest{Token: ctx.Get("Authorization", "NOT_FOUND")}
		log.Printf("Authorization : %s", request.Token)

		auth, err := adminSuperService.Verify(ctx.UserContext(), request)
		if err != nil {
			log.Printf(err.Error())
			return fiber.ErrUnauthorized
		}

		log.Printf("user : %+v", auth.ID)
		ctx.Locals("auth", auth)
		return ctx.Next()
	}
}
func GetAdminSuperAuth(ctx fiber.Ctx) *model.Auth {
	return ctx.Locals("auth").(*model.Auth)
}

func AdminPuskesmasAuth(adminPuskesmasService *service.AdminPuskesmasService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		request := &model.VerifyAdminPuskesmasRequest{Token: ctx.Get("Authorization", "NOT_FOUND")}
		log.Printf("Authorization : %s", request.Token)

		auth, err := adminPuskesmasService.Verify(ctx.UserContext(), request)
		if err != nil {
			log.Printf(err.Error())
			return fiber.ErrUnauthorized
		}

		log.Printf("user : %+v", auth.ID)
		ctx.Locals("auth", auth)
		return ctx.Next()
	}
}
func GetAdminPuskesmasAuth(ctx fiber.Ctx) *model.Auth {
	return ctx.Locals("auth").(*model.Auth)
}

func AdminApotekAuth(adminApotekService *service.AdminApotekService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		request := &model.VerifyAdminApotekRequest{Token: ctx.Get("Authorization", "NOT_FOUND")}
		log.Printf("Authorization : %s", request.Token)

		auth, err := adminApotekService.Verify(ctx.UserContext(), request)
		if err != nil {
			log.Printf(err.Error())
			return fiber.ErrUnauthorized
		}

		log.Printf("user : %+v", auth.ID)
		ctx.Locals("auth", auth)
		return ctx.Next()
	}
}
func GetAdminApotekAuth(ctx fiber.Ctx) *model.Auth {
	return ctx.Locals("auth").(*model.Auth)
}

func AdminSuperOrPuskesmasAuth(adminSuperService *service.AdminSuperService, adminPuskesmasService *service.AdminPuskesmasService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		token := ctx.Get("Authorization", "NOT_FOUND")
		requestSuper := &model.VerifyAdminSuperRequest{Token: token}
		requestPuskesmas := &model.VerifyAdminPuskesmasRequest{Token: token}

		log.Printf("Authorization : %s", token)

		authSuper, errSuper := adminSuperService.Verify(ctx.UserContext(), requestSuper)
		if errSuper == nil {
			log.Printf("Authenticated as AdminSuper: %+v", authSuper.ID)
			ctx.Locals("auth", authSuper)
			return ctx.Next()
		}

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

func GetAdminSuperOrPuskesmasAuth(ctx fiber.Ctx) *model.Auth {
	return ctx.Locals("auth").(*model.Auth)
}
