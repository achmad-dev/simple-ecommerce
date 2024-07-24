package handlers

import (
	"net/http"
	"strconv"

	"github.com/achmad-dev/simple-ecommerce/gateway/internal/response"
	"github.com/achmad-dev/simple-ecommerce/gateway/service"
	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func (h *OrderHandler) CreateOrder(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("userID"))
	if (err != nil) || userID <= 0 {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid user ID", err))
	}

	var request struct {
		CartIDs         []int `json:"cart_ids"`
		PaymentMethodID int   `json:"payment_method_id"`
	}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid request", err))
	}

	if request.PaymentMethodID <= 0 {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse(" Payment Method ID are required", nil))
	}

	err = h.orderService.CreateOrder(userID, request.PaymentMethodID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewErrorResponse("Failed to create order", err))
	}

	return c.JSON(http.StatusOK, response.NewSuccessResponse("Order created successfully"))
}

func (h *OrderHandler) PayOrder(c echo.Context) error {
	orderID, err := strconv.Atoi(c.Param("orderID"))
	if err != nil || orderID <= 0 {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid order ID", err))
	}

	err = h.orderService.PayOrder(orderID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewErrorResponse("Failed to pay order", err))
	}

	return c.JSON(http.StatusOK, response.NewSuccessResponse("Order paid successfully"))
}

func (h *OrderHandler) GetOrdersByUserID(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil || userID <= 0 {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid user ID", err))
	}

	orders, err := h.orderService.GetOrdersByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewErrorResponse("Failed to fetch orders", err))
	}

	return c.JSON(http.StatusOK, response.NewSuccessResponse(orders))
}
