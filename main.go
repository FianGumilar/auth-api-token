package main

import (
	"github.com/FianGumilar/auth-api-token/internal/api"
	"github.com/FianGumilar/auth-api-token/internal/component"
	"github.com/FianGumilar/auth-api-token/internal/config"
	"github.com/FianGumilar/auth-api-token/internal/repository"
	"github.com/FianGumilar/auth-api-token/internal/service"
	"github.com/FianGumilar/auth-api-token/middleware"
	"github.com/gofiber/fiber/v2"
)

func main() {
	conf := config.Get()

	dbConnection := component.GetConnectionDatabase(conf)

	cacheRepository := component.GetCacheConnection()
	userRepository := repository.NewUserRepository(dbConnection)

	userService := service.NewUserService(userRepository, cacheRepository)

	authMiddleware := middleware.Authenticate(userService)

	app := fiber.New()
	api.NewAuth(app, userService, authMiddleware)

	app.Listen(conf.Srv.Host + ":" + conf.Srv.Port)

}
