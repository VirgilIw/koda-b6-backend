package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/service"
)

type CategoriesHandler struct {
	service *service.CategoriesService
}

func NewCategoriesHandler(service *service.CategoriesService) *CategoriesHandler {
	return &CategoriesHandler{
		service: service,
	}
}

// Get all category godoc
// @Summary      get all category
// @Description  get all category
// @Tags         Categories
// @Produce      json
// @Success      200      {object}  dto.ResponseCategories
// @Failure      400      {object}  dto.ResponseCategories
// @Failure      401      {object}  dto.ResponseCategories
// @Router       /admin/categories [get]
func (c *CategoriesHandler) GetCategories(ctx *gin.Context) {

	categories, err := c.service.GetCategories(ctx.Request.Context())

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseCategories{
			Success: false,
			Message: "internal server error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseCategories{
		Success: true,
		Message: "success get all categories",
		Result:  categories,
	})
}

// Get category by Id godoc
// @Summary      get category by id
// @Description  get category by id
// @Tags         Categories
// @Accept       json
// @Produce      json
// @Param        id  	   path    int  true  "category id"
// @Success      200      {object}  dto.ResponseCategory
// @Failure      400      {object}  dto.ResponseCategory
// @Failure      401      {object}  dto.ResponseCategory
// @Router       /admin/categories/{id} [get]
func (c *CategoriesHandler) GetCategoryByID(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, dto.ResponseCategory{
			Success: false,
			Message: "bad request",
			Error:   "invalid id",
		})
		return
	}

	categories, err := c.service.GetCategoryById(ctx.Request.Context(), id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseCategory{
			Success: false,
			Message: "internal server error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseCategory{
		Success: true,
		Message: "success get category",
		Result:  categories,
	})
}

// Create category by Id godoc
// @Summary      create category
// @Description  create category
// @Tags         Categories
// @Accept       json
// @Produce      json
// @Param 		 request   body dto.CategoryRequest true "create category"
// @Success      201      {object}  dto.ResponseCategory
// @Failure      400      {object}  dto.ResponseCategory
// @Failure      500      {object}  dto.ResponseCategory
// @Router       /admin/categories [post]
func (c *CategoriesHandler) CreateCategory(ctx *gin.Context) {
	var catName dto.CategoryRequest

	if err := ctx.ShouldBindJSON(&catName); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseCategory{
			Success: false,
			Message: "bad request",
			Error:   err.Error(),
		})
		return
	}

	newCategory, err := c.service.CreateCategory(ctx.Request.Context(), catName)
	// log.Println(newCategory)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseCategory{
			Success: false,
			Message: "internal server error",
			Error:   "failed to create category",
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.ResponseCategory{
		Success: true,
		Message: "success create new category",
		Result:  newCategory,
	})
}

// Update category by Id godoc
// @Summary      update category
// @Description  update category
// @Tags         Categories
// @Accept       json
// @Produce      json
// @Param 		 id  path 	int true "category id"
// @Param 		 request   body dto.CategoryRequest true "update category"
// @Success      200      {object}  dto.ResponseCategory
// @Failure      400      {object}  dto.ResponseCategory
// @Failure      500      {object}  dto.ResponseCategory
// @Router       /admin/categories/{id} [patch]
func (c *CategoriesHandler) UpdateCategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseCategory{
			Success: false,
			Message: "invalid category id",
		})
		return
	}

	// Bind body JSON
	var req dto.CreateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseCategory{
			Success: false,
			Message: "bad request",
			Error:   err.Error(),
		})
		return
	}

	req.Id = id

	updatedCategory, err := c.service.UpdateCategory(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseCategory{
			Success: false,
			Message: "failed to update category",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseCategory{
		Success: true,
		Message: "success update category",
		Result:  updatedCategory,
	})
}

// Delete category by Id godoc
// @Summary      Delete category
// @Description  Delete category
// @Tags         Categories
// @Accept       json
// @Produce      json
// @Param 		 id  path 	int true "category id"
// @Success      200      {object}  dto.ResponseCategory
// @Failure      400      {object}  dto.ResponseCategory
// @Failure      500      {object}  dto.ResponseCategory
// @Router       /admin/categories/{id} [delete]
func (c *CategoriesHandler) DeleteCategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseCategory{
			Success: false,
			Message: "invalid category id",
		})
		return
	}

	if err := c.service.DeleteCategory(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseCategory{
			Success: false,
			Message: "failed to delete category",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseCategory{
		Success: true,
		Message: "success delete category",
	})
}
