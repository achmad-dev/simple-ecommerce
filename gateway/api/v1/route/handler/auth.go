package handlers

import (
	"net/http"

	"github.com/achmad-dev/simple-ecommerce/gateway/dto"
	"github.com/achmad-dev/simple-ecommerce/gateway/internal/response"
	"github.com/achmad-dev/simple-ecommerce/gateway/internal/utils"
	"github.com/achmad-dev/simple-ecommerce/gateway/service"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	userService service.UserService
	jwtHelper   utils.JwtHelper
}

func NewAuthHandler(userService service.UserService, jwtHelper utils.JwtHelper) *AuthHandler {
	return &AuthHandler{userService: userService, jwtHelper: jwtHelper}
}

func (h *AuthHandler) Register(c echo.Context) error {
	var userDto dto.AuthUserDto
	if err := c.Bind(&userDto); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid input", err))
	}

	err := h.userService.RegisterUser(userDto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewErrorResponse("Failed to register user", err))
	}

	return c.JSON(http.StatusOK, response.NewSuccessResponse(map[string]string{
		"message": "User registered successfully",
	}))
}

func (h *AuthHandler) Login(c echo.Context) error {
	var userDto dto.AuthUserDto
	if err := c.Bind(&userDto); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid input", err))
	}

	user, err := h.userService.GetUserByEmail(userDto.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.NewErrorResponse("Invalid email or password", err))
	}

	ok := utils.CheckPasswordHash(userDto.Password, user.PasswordHash)
	if !ok {
		return c.JSON(http.StatusUnauthorized, response.NewErrorResponse("Invalid email or password", err))
	}

	token, err := h.jwtHelper.GenerateToken(user.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewErrorResponse("Failed to generate token", err))
	}

	return c.JSON(http.StatusOK, response.NewSuccessResponse(map[string]string{
		"token": token,
	}))
}
