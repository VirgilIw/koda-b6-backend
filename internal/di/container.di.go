package di

import (
	"github.com/jackc/pgx/v5"

	"github.com/virgiIw/koda-b6-coffeshopdb/internal/handler"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/repository"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/service"
)

type Container struct {
	userRepo    *repository.UserRepository
	userService *service.UserService
	userHandler *handler.UserHandler

	authService *service.AuthService
	authHandler *handler.AuthHandler
}

func NewContainer(db *pgx.Conn) *Container {

	c := &Container{}

	// USER
	c.userRepo = repository.NewUserRepository(db)
	c.userService = service.NewUserService(c.userRepo)
	c.userHandler = handler.NewUserHandler(c.userService)

	c.authService = service.NewAuthService(c.userRepo)
	c.authHandler = handler.NewAuthHandler(c.authService)

	return c
}

func (c *Container) UserHandler() *handler.UserHandler {
	return c.userHandler
}

func (c *Container) AuthHandler() *handler.AuthHandler {
	return c.authHandler
}
