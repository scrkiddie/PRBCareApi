package middleware

import (
	"github.com/gofiber/fiber/v3"
	"log"
	"prbcare_be/internal/model"
	"prbcare_be/internal/service"
)

func AdminSuperAuth(adminSuperService *service.AdminSuperService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		request := &model.VerifyUserRequest{Token: ctx.Get("Authorization", "NOT_FOUND")}
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
