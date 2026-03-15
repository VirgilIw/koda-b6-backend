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

	forgotPwdRepo    *repository.ForgotPwdRepository
	forgotPwdService *service.ForgotPwdService

	authHandler *handler.AuthHandler

	productRepo    *repository.ProductRepository
	productService *service.ProductService
	productHandler *handler.ProductHandler

	filterRepo    *repository.SearchRepository
	filterService *service.SearchService
	filterHandler *handler.SearchHandler

	orderRepo    *repository.OrderRepository
	orderService *service.OrderService
	orderHandler *handler.OrderHandler

	categoryRepo    *repository.CategoriesRepository
	categoryService *service.CategoriesService
	categoryHandler *handler.CategoriesHandler

	sizesRepo    *repository.SizesRepository
	sizesService *service.SizesService
	sizesHandler *handler.SizesHandler

	landingRepo    *repository.LandingRepository
	landingService *service.LandingService
	landingHandler *handler.LandingHandler

	variantRepo    *repository.VariantRepository
	variantService *service.VariantService
	variantHandler *handler.VariantHandler

	imagesRepo    *repository.ImageRepository
	imagesService *service.ImageService
	imagesHandler *handler.ImageHandler
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

	c.forgotPwdRepo = repository.NewForgotPwdRepository(c.db)
	c.forgotPwdService = service.NewForgotPwdService(c.forgotPwdRepo, c.userRepo)

	c.authService = service.NewAuthService(c.userRepo)
	c.authHandler = handler.NewAuthHandler(c.authService, c.forgotPwdService)

	c.productRepo = repository.NewProductRepository(c.db, c.rdb)
	c.productService = service.NewProductService(c.productRepo)
	c.productHandler = handler.NewProductService(c.productService)

	c.orderRepo = repository.NewOrderRepository(c.db)
	c.orderService = service.NewOrderService(c.orderRepo)
	c.orderHandler = handler.NewOrderHandler(c.orderService)

	c.categoryRepo = repository.NewCategoriesRepository(c.db)
	c.categoryService = service.NewCategoriesService(c.categoryRepo)
	c.categoryHandler = handler.NewCategoriesHandler(c.categoryService)

	c.sizesRepo = repository.NewSizesRepository(c.db)
	c.sizesService = service.NewSizesService(c.sizesRepo)
	c.sizesHandler = handler.NewSizesHandler(c.sizesService)

	c.landingRepo = repository.NewLandingRepository(c.db)
	c.landingService = service.NewLandingService(c.landingRepo)
	c.landingHandler = handler.NewLandingService(c.landingService)

	c.variantRepo = repository.NewVariantRepository(c.db)
	c.variantService = service.NewVariantService(c.variantRepo)
	c.variantHandler = handler.NewVariantHandler(c.variantService)

	c.imagesRepo = repository.NewImageRepository(c.db)
	c.imagesService = service.NewImageService(c.imagesRepo)
	c.imagesHandler = handler.NewImageHandler(c.imagesService)

	c.filterRepo = repository.NewSearchRepository(c.db)
	c.filterService = service.NewSearchService(c.filterRepo, c.categoryRepo)
	c.filterHandler = handler.NewSearchHandler(c.filterService)
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

func (c *Container) SizesHandler() *handler.SizesHandler {
	return c.sizesHandler
}

func (c *Container) LandingHandler() *handler.LandingHandler {
	return c.landingHandler
}

func (c *Container) VariantHandler() *handler.VariantHandler {
	return c.variantHandler
}

func (c *Container) ImagesHandler() *handler.ImageHandler {
	return c.imagesHandler
}

func (c *Container) SearchHandler() *handler.SearchHandler {
	return c.filterHandler
}
