package api

import (
	"github.com/FianGumilar/auth-api-token/domain"
	"github.com/FianGumilar/auth-api-token/dto"
	"github.com/FianGumilar/auth-api-token/utils"
	"github.com/gofiber/fiber/v2"
)

type authApi struct {
	userService domain.UserService
}

func NewAuth(app *fiber.App, userService domain.UserService, authMid fiber.Handler) {
	api := &authApi{
		userService: userService,
	}

	app.Post("/token/generate", api.GetToken)
	app.Get("/token/validate", authMid, api.ValidateToken)
}

func (a authApi) GetToken(ctx *fiber.Ctx) error {
	var req dto.AutRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON("Error at body parser")
	}

	token, err := a.userService.Authenticate(ctx.Context(), req)
	if err != nil {
		return ctx.SendStatus(utils.GetHttpStatus(err))
	}
	return ctx.Status(200).JSON(token)
}

func (a authApi) ValidateToken(ctx *fiber.Ctx) error {
	user := ctx.Locals("x-user")

	return ctx.Status(200).JSON(user)

}
