package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/service"
)

type CategorieHandler struct {
	service *service.CategorieService
}

func NewCategoriesHandler(service *service.CategorieService) *CategorieHandler {
	return &CategorieHandler{
		service: service,
	}
}

// Get all category godoc
// @Summary      get all category
// @Description  get all category
// @Tags         Category
// @Produce      json
// @Success      200      {object}  dto.ResponseCategories
// @Failure      400      {object}  dto.ResponseCategories
// @Failure      401      {object}  dto.ResponseCategories
// @Router       /categories [get]
func (c *CategorieHandler) GetCategories(ctx *gin.Context) {

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
