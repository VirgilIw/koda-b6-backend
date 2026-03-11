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

	productRepo    *repository.ProductRepository
	productService *service.ProductService
	productHandler *handler.ProductHandler

	orderRepo    *repository.OrderRepository
	orderService *service.OrderService
	orderHandler *handler.OrderHandler

	categoryRepo    *repository.CategoriesRepository
	categoryService *service.CategoriesService
	categoryHandler *handler.CategoriesHandler
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

	c.productRepo = repository.NewProductRepository(c.db, c.rdb)
	c.productService = service.NewProductService(c.productRepo)
	c.productHandler = handler.NewProductService(c.productService)

	c.orderRepo = repository.NewOrderRepository(c.db)
	c.orderService = service.NewOrderService(c.orderRepo)
	c.orderHandler = handler.NewOrderHandler(c.orderService)

	c.categoryRepo = repository.NewCategoriesRepository(c.db)
	c.categoryService = service.NewCategoriesService(c.categoryRepo)
	c.categoryHandler = handler.NewCategoriesHandler(c.categoryService)
}

func (c *Container) UserHandler() *handler.UserHandler {
	return c.userHandler
}

func (c *Container) AuthHandler() *handler.AuthHandler {
	return c.authHandler
}

func (c *Container) ProductHandler() *handler.ProductHandler {
	return c.productHandler
}

func (c *Container) OrderHandler() *handler.OrderHandler {
	return c.orderHandler
}

func (c *Container) CategoriesHandler() *handler.CategoriesHandler {
	return c.categoryHandler
}
