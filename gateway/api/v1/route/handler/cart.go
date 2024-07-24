package handlers

import (
	"net/http"
	"strconv"

	"github.com/achmad-dev/simple-ecommerce/gateway/internal/response"
	"github.com/achmad-dev/simple-ecommerce/gateway/service"
	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	cartService service.CartService
}

func NewCartHandler(cartService service.CartService) *CartHandler {
	return &CartHandler{cartService: cartService}
}

func (h *CartHandler) AddProductToCart(c echo.Context) error {
	userEmail := c.Get("email").(string)
	productID, err := strconv.Atoi(c.FormValue("productID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid product ID", err))
	}

	err = h.cartService.AddProductToCart(userEmail, productID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewErrorResponse("Failed to add product to cart", err))
	}

	return c.JSON(http.StatusOK, response.NewSuccessResponse("Product added to cart"))
}

func (h *CartHandler) RemoveProductFromCart(c echo.Context) error {
	userEmail := c.Get("email").(string)

	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid product ID", err))
	}

	err = h.cartService.RemoveProductFromCart(userEmail, productID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewErrorResponse("Failed to remove product from cart", err))
	}

	return c.JSON(http.StatusOK, response.NewSuccessResponse("Product removed from cart"))
}

func (h *CartHandler) GetCartByUserID(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid user ID", err))
	}

	cartItems, err := h.cartService.GetCartByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewErrorResponse("Failed to fetch cart items", err))
	}

	return c.JSON(http.StatusOK, response.NewSuccessResponse(cartItems))
}
