package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/service"
)

type ImageHandler struct {
	service *service.ImageService
}

func NewImageHandler(service *service.ImageService) *ImageHandler {
	return &ImageHandler{
		service: service,
	}
}

// Create image godoc
// @Summary      Create new image
// @Description  Upload and create new image
// @Tags         Images
// @Accept       multipart/form-data
// @Produce      json
// @Param        image formData file true "Image file"
// @Success      201 {object} dto.ResponseImage
// @Failure      400 {object} dto.ResponseImage
// @Failure      500 {object} dto.ResponseImage
// @Router       /admin/images [post]
func (i *ImageHandler) CreateImage(ctx *gin.Context) {
	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseImage{
			Success: false,
			Message: "image file is required",
			Error:   err.Error(),
		})
		return
	}

	const maxSize = 1 << 20
	if file.Size > maxSize {
		ctx.JSON(http.StatusBadRequest, dto.ResponseImage{
			Success: false,
			Message: "image size must be less than 1MB",
		})
		return
	}
	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), filepath.Base(file.Filename))

	path := "./images/"
	fullPath := path + filename

	if err := ctx.SaveUploadedFile(file, fullPath); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseImage{
			Success: false,
			Message: "failed to save image",
			Error:   err.Error(),
		})
		return
	}

	var req dto.ImageRequest
	req.ImagePath = fullPath

	result, err := i.service.CreateImage(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseImage{
			Success: false,
			Message: "internal server error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.ResponseImage{
		Success: true,
		Message: "success create image",
		Result:  result,
	})
}

// Get all images godoc
// @Summary      Get all images
// @Description  Retrieve all images
// @Tags         Images
// @Produce      json
// @Success      200 {object} dto.ResponseImages
// @Failure      500 {object} dto.ResponseImages
// @Router       /admin/images [get]
func (i *ImageHandler) GetImages(ctx *gin.Context) {

	result, err := i.service.GetImages(ctx.Request.Context())

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseImages{
			Success: false,
			Message: "internal server error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseImages{
		Success: true,
		Message: "success get all images",
		Result:  result,
	})
}

// Get image by id godoc
// @Summary      Get image by id
// @Description  Retrieve image by id
// @Tags         Images
// @Produce      json
// @Param        id path int true "Image ID"
// @Success      200 {object} dto.ResponseImage
// @Failure      400 {object} dto.ResponseImage
// @Failure      404 {object} dto.ResponseImage
// @Failure      500 {object} dto.ResponseImage
// @Router       /admin/images/{id} [get]
func (i *ImageHandler) GetImageById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseImage{
			Success: false,
			Message: "bad request",
			Error:   err.Error(),
		})
		return
	}

	result, err := i.service.GetImageById(ctx.Request.Context(), id)

	if err != nil && err.Error() == "id not found" {
		ctx.JSON(http.StatusNotFound, dto.ResponseImage{
			Success: false,
			Message: "id not found",
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseImage{
			Success: false,
			Message: "internal server error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseImage{
		Success: true,
		Message: "success get data by id",
		Result:  result,
	})
}

// Update image godoc
// @Summary      Update image
// @Description  Update image by id
// @Tags         Images
// @Accept       multipart/form-data
// @Produce      json
// @Param        id path int true "Image ID"
// @Param        image formData file true "Image file"
// @Success      200 {object} dto.ResponseImage
// @Failure      400 {object} dto.ResponseImage
// @Failure      404 {object} dto.ResponseImage
// @Failure      500 {object} dto.ResponseImage
// @Router       /admin/images/{id} [patch]
func (i *ImageHandler) UpdateImage(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseImage{
			Success: false,
			Message: "bad request",
			Error:   err.Error(),
		})
		return
	}

	// Ambil data image lama
	oldImage, err := i.service.GetImageById(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.ResponseImage{
			Success: false,
			Message: "image not found",
			Error:   err.Error(),
		})
		return
	}

	// Ambil file baru
	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseImage{
			Success: false,
			Message: "image file is required",
			Error:   err.Error(),
		})
		return
	}

	if file.Size > 1<<20 {
		ctx.JSON(http.StatusBadRequest, dto.ResponseImage{
			Success: false,
			Message: "image size must be less than 1MB",
		})
		return
	}

	// Buat nama file unik
	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)

	path := "./images/"
	fullPath := path + filename

	// Simpan file baru
	if err := ctx.SaveUploadedFile(file, fullPath); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseImage{
			Success: false,
			Message: "failed to save image",
			Error:   err.Error(),
		})
		return
	}

	// Hapus file lama
	if oldImage.ImagePath != "" {
		_ = os.Remove(oldImage.ImagePath)
	}

	// Update database
	req := dto.ImageRequest{
		ImagePath: fullPath,
	}

	result, err := i.service.UpdateImage(ctx.Request.Context(), req, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseImage{
			Success: false,
			Message: "internal server error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseImage{
		Success: true,
		Message: "success update image",
		Result:  result,
	})
}

// Delete image by id godoc
// @Summary      Delete image by id
// @Description  Retrieve image by id
// @Tags         Images
// @Produce      json
// @Param        id path int true "Image ID"
// @Success      200 {object} dto.ResponseImage
// @Failure      400 {object} dto.ResponseImage
// @Failure      404 {object} dto.ResponseImage
// @Failure      500 {object} dto.ResponseImage
// @Router       /admin/images/{id} [delete]
func (i *ImageHandler) DeleteImageById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseImage{
			Success: false,
			Message: "bad request",
			Error:   err.Error(),
		})
		return
	}

	err = i.service.DeleteImageById(ctx.Request.Context(), id)

	if err != nil && err.Error() == "id not found" {
		ctx.JSON(http.StatusNotFound, dto.ResponseImage{
			Success: false,
			Message: "id not found",
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseImage{
			Success: false,
			Message: "internal server error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseImage{
		Success: true,
		Message: "success delete data by id",
	})
}
