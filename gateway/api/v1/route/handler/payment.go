package handlers

import (
	"net/http"

	"github.com/achmad-dev/simple-ecommerce/gateway/dto"
	"github.com/achmad-dev/simple-ecommerce/gateway/internal/response"
	"github.com/achmad-dev/simple-ecommerce/gateway/service"
	"github.com/labstack/echo/v4"
)

type PaymentHandler struct {
	paymentService service.PaymentService
	userService    service.UserService
}

func NewPaymentHandler(paymentService service.PaymentService, userService service.UserService) *PaymentHandler {
	return &PaymentHandler{paymentService: paymentService, userService: userService}
}

func (h *PaymentHandler) CreatePaymentMethod(c echo.Context) error {
	userEmail := c.Get("email").(string)
	var paymentDto dto.PaymentMethodDto
	if err := c.Bind(&paymentDto); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid input", err))
	}

	// userID := utils.GetUserIDFromContext(c)
	user, err := h.userService.GetUserByEmail(userEmail)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewErrorResponse("Failed to create payment method", err))
	}
	paymentDto.UserID = user.Id
	err = h.paymentService.CreatePaymentMethod(paymentDto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewErrorResponse("Failed to create payment method", err))
	}

	return c.JSON(http.StatusOK, response.NewSuccessResponse(map[string]string{
		"message": "Payment method created successfully",
	}))
}
