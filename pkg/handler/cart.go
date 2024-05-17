package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary add to cart
// @Security ApiKeyAuth
// @Tags cart
// @Accept json
// @Produce json
// @Param id path string true "product id"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /api/products/{id} [post]
func (h *Handler) addToCart(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, "user id not found")
		return
	}

	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Cart.Add(userId, productId)
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary decrease the quantity by one
// @Security ApiKeyAuth
// @Tags cart
// @Accept json
// @Produce json
// @Param id path string true "product id"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /api/cart/minus/{id} [put]
func (h *Handler) minus(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, "user id not found")
		return
	}

	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Cart.Minus(userId, productId); err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary increase the quantity by one
// @Security ApiKeyAuth
// @Tags cart
// @Accept json
// @Produce json
// @Param id path string true "product id"
// @Success 200 {integer} integer
// @Failure 400 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /api/cart/plus/{id} [put]
func (h *Handler) plus(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, "user id not found")
		return
	}

	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Cart.Plus(userId, productId); err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary clear the cart
// @Security ApiKeyAuth
// @Tags cart
// @Accept json
// @Produce json
// @Success 200 {integer} integer
// @Failure 400 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /api/cart [delete]
func (h *Handler) deleteAllFromCart(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, "user id not found")
		return
	}

	if err := h.services.Cart.Delete(userId); err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary All in cart
// @Security ApiKeyAuth
// @Tags cart
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 400 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /api/cart [get]
func (h *Handler) getAllFromCart(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, "user id not found")
		return
	}

	products, err := h.services.Cart.GetAllFromCart(userId)
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	var totalAmount float64
	if products != nil {
		totalAmount, err = h.services.Cart.GetTotalAmout(userId)
		if err != nil {
			NewErrorResponce(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Products":     products,
		"Total amount": totalAmount,
	})
}
