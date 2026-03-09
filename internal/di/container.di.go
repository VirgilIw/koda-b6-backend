package di

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/handler"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/repository"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/service"
)

type Container struct {
	db  *pgxpool.Pool
	rdb *redis.Client

	userRepo    *repository.UserRepository
	userService *service.UserService
	userHandler *handler.UserHandler

	authService *service.AuthService
	authHandler *handler.AuthHandler
}

func NewContainer(db *pgxpool.Pool, rdb *redis.Client) *Container {

	container := &Container{
		db:  db,
		rdb: rdb,
	}
	container.initDependencies()
	return container
}

// USER

func (c *Container) initDependencies() {
	c.userRepo = repository.NewUserRepository(c.db, c.rdb)
	c.userService = service.NewUserService(c.userRepo)
	c.userHandler = handler.NewUserHandler(c.userService)

	c.authService = service.NewAuthService(c.userRepo)
	c.authHandler = handler.NewAuthHandler(c.authService)
}

func (c *Container) UserHandler() *handler.UserHandler {
	return c.userHandler
}

func (c *Container) AuthHandler() *handler.AuthHandler {
	return c.authHandler
}
