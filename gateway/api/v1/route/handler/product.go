package handlers

import (
	"net/http"
	"strconv"

	"github.com/achmad-dev/simple-ecommerce/gateway/internal/response"
	"github.com/achmad-dev/simple-ecommerce/gateway/service"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) FetchAllProducts(c echo.Context) error {
	products, err := h.productService.FetchAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewErrorResponse("Failed to fetch products", err))
	}
	return c.JSON(http.StatusOK, response.NewSuccessResponse(products))
}

func (h *ProductHandler) FetchProductsPaginated(c echo.Context) error {
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit <= 0 {
		limit = 10 // default limit
	}

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil || offset < 0 {
		offset = 0 // default offset
	}

	products, err := h.productService.FetchProductsPaginated(limit, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewErrorResponse("Failed to fetch products", err))
	}
	return c.JSON(http.StatusOK, response.NewSuccessResponse(products))
}

func (h *ProductHandler) FetchProductByName(c echo.Context) error {
	name := c.QueryParam("name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse("Product name is required", nil))
	}

	products, err := h.productService.FetchProductByName(name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewErrorResponse("Failed to fetch products", err))
	}
	return c.JSON(http.StatusOK, response.NewSuccessResponse(products))
}
