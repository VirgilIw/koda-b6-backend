package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// GetUsers godoc
// @Summary      Get users
// @Description  Get users
// @Tags         Users
// @Produce      json
// @Success      200  {object}  dto.Response
// @Failure      400  {object}  dto.Response
// @Failure      500  {object}  dto.Response
// @Security     BearerAuth
// @Router       /admin/users [get]
func (u *UserHandler) GetUsers(ctx *gin.Context) {
	datas, err := u.service.GetUsers(ctx.Request.Context())

	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "bad request",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "success get users",
		Result:  datas,
	})
}

// @Summary      Get current user
// @Description  Get logged-in user profile
// @Tags         Users
// @Produce      json
// @Success      200  {object}  dto.ResponseOneData
// @Failure      400  {object}  dto.Response
// @Failure      500  {object}  dto.Response
// @Security     BearerAuth
// @Router       /admin/users/me [get]
func (u *UserHandler) GetUserById(ctx *gin.Context) {
	userId, ok := ctx.Get("userID")
	if !ok {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "user id not found in token",
			Error:   "invalid token",
		})
		return
	}

	id, ok := userId.(int)
	if !ok {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "invalid user id format",
			Error:   "type assertion failed",
		})
		return
	}

	user, err := u.service.GetUserById(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "internal server error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseOneData{
		Success: true,
		Message: "success get data by id",
		Result:  user,
	})
}

// UpdateProfile godoc
// @Summary      Update profile
// @Description  Update the logged-in user's profile, including optional picture upload
// @Tags         Users
// @Accept       multipart/form-data
// @Produce      json
// @Param        fullname   formData  string  false  "Full name"
// @Param        email      formData  string  false  "Email"
// @Param        password   formData  string  false  "Password"
// @Param        phone      formData  string  false  "Phone number"
// @Param        address    formData  string  false  "Address"
// @Param        role       formData  string  false  "Role"
// @Param        picture    formData  file    false  "Profile picture"
// @Success      200        {object}  dto.ResponseOneData
// @Failure      400        {object}  dto.ResponseOneData
// @Failure      500        {object}  dto.ResponseOneData
// @Security     BearerAuth
// @Router       /admin/users [patch]
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	idRaw, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Message: "unauthorized",
		})
		return
	}

	userId, ok := idRaw.(int)
	if !ok {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "invalid user id",
		})
		return
	}

	var req dto.UpdateUserRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "bad request: " + err.Error(),
		})
		return
	}

	var picture *string
	file, err := c.FormFile("picture")
	if err == nil {
		// bikin nama unik
		filename := fmt.Sprintf("%d_%s", time.Now().Unix(), filepath.Base(file.Filename))
		path := "./images/"
		fullPath := path + filename

		// save file (sekali aja)
		if err := c.SaveUploadedFile(file, fullPath); err != nil {
			c.JSON(http.StatusInternalServerError, dto.Response{
				Success: false,
				Message: "failed to upload picture",
			})
			return
		}

		// simpan filename ke DB
		picture = &filename
	}

	updatedUser, err := h.service.UpdateProfile(
		c.Request.Context(),
		userId,
		req,
		picture,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "failed to update profile",
			Error:   err.Error(),
		})
		return
	}

	respUser := dto.Users{
		Id:        updatedUser.Id,
		FullName:  updatedUser.FullName,
		Email:     updatedUser.Email,
		Picture:   updatedUser.Picture,
		Phone:     updatedUser.Phone,
		Address:   updatedUser.Address,
		Role:      updatedUser.Role,
		CreatedAt: updatedUser.CreatedAt,
		UpdatedAt: updatedUser.UpdatedAt,
	}

	c.JSON(http.StatusOK, dto.ResponseOneData{
		Success: true,
		Message: "profile updated",
		Result:  respUser,
	})
}

// DeleteUser godoc
// @Summary      Delete users
// @Description  Delete users
// @Tags         Users
// @Produce      json
// @Param        id    path 	int true  "User Id"
// @Success      200  {object}  dto.Response
// @Failure      400  {object}  dto.Response
// @Failure      500  {object}  dto.Response
// @Security     BearerAuth
// @Router       /admin/users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "bad request",
		})
		return
	}

	if err := h.service.DeleteUser(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "bad request",
		})
		return
	}
	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "success delete user",
	})
}

// CreateUser godoc
// @Summary      Create a new user
// @Description  Create user with optional picture upload
// @Tags         Users
// @Accept       multipart/form-data
// @Produce      json
// @Param        full_name   formData  string  true   "Full Name"
// @Param        email       formData  string  true   "Email"
// @Param        password    formData  string  true   "Password"
// @Param        phone       formData  string  false  "Phone Number"
// @Param        address     formData  string  false  "Address"
// @Param        role        formData  string  false  "Role"
// @Param        picture     formData  file    false  "Profile Picture"
// @Success      201 {object} dto.Response
// @Failure      400 {object} dto.Response
// @Failure      500 {object} dto.Response
// @Router       /admin/users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	// Tangkap file gambar
	file, err := c.FormFile("picture")
	if err != nil && err != http.ErrMissingFile {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "failed to get picture",
			Error:   err.Error(),
		})
		return
	}

	// Tangkap field lain
	fullName := c.PostForm("fullname")
	email := c.PostForm("email")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	address := c.PostForm("address")
	role := c.PostForm("role")

	// Mapping ke DTO / Model
	req := dto.CreateUserRequest{
		FullName: fullName,
		Email:    email,
		Password: password,
		Phone:    phone,
		Address:  address,
		Role:     role,
	}

	if file != nil {
		path := "./images/"

		if err := c.SaveUploadedFile(file, path+file.Filename); err != nil {
			c.JSON(http.StatusInternalServerError, dto.Response{
				Success: false,
				Message: "failed to save picture",
				Error:   err.Error(),
			})
			return
		}
		req.Picture = file.Filename
	}

	// Panggil service
	_, err = h.service.CreateUser(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "failed to create user",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Message: "create user success",
	})
}
