package handler

import (
	"net/http"
	"onlineshop/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary make order
// @Security ApiKeyAuth
// @Tags order
// @Accept json
// @Produce json
// @Param address body models.Address true "address"
// @Success 200 {integer} integer
// @Failure 400 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /api/cart/order [post]
func (h *Handler) createOrder(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, "user id not found")
		return
	}

	totalAmount, err := h.services.Cart.GetTotalAmout(userId)
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	var address models.Address
	if err := c.BindJSON(&address); err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	products, err := h.services.Cart.GetAllFromCart(userId)
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error()+"asdasd")
		return
	}

	id, err := h.services.Order.Create(userId, totalAmount, address, products)
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.services.Cart.Delete(userId); err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary order history
// @Security ApiKeyAuth
// @Tags order
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 400 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /api/cart/orders [get]
func (h *Handler) getOrderForUser(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, "user id not found")
		return
	}

	orders, err := h.services.Order.GetAllForUser(userId)
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	orderMap := make(map[int]*models.OrderWithProducts)
	for _, order := range orders {
		if _, found := orderMap[order.Id]; !found {

			orderMap[order.Id] = &models.OrderWithProducts{
				Id:          order.Id,
				TotalAmount: order.TotalAmount,
				Status:      order.Status,
				CreatedAt:   order.CreatedAt,
				Address:     order.Address,
			}

		}

		product := models.GetProductsFromCart{
			Id:       order.ProductId,
			Name:     order.Name,
			Image:    order.Image,
			Price:    order.Price,
			Quantity: order.Quantity,
		}
		orderMap[order.Id].Products = append(orderMap[order.Id].Products, product)

	}

	// Convert map to slice
	result := make([]models.OrderWithProducts, 0, len(orderMap))
	for _, v := range orderMap {
		result = append(result, *v)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"orders": result,
	})

}

// @Summary all orders
// @Security ApiKeyAuth
// @Tags Admin
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 400 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /api/admin/orders [get]
func (h *Handler) getAllOrders(c *gin.Context) {
	orders, err := h.services.Order.GetAll()
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	orderMap := make(map[int]*models.OrderWithProducts)
	for _, order := range orders {
		if _, found := orderMap[order.Id]; !found {

			orderMap[order.Id] = &models.OrderWithProducts{
				Id:          order.Id,
				TotalAmount: order.TotalAmount,
				Status:      order.Status,
				CreatedAt:   order.CreatedAt,
				UserId:      order.UserId,
				Address:     order.Address,
			}

		}

		product := models.GetProductsFromCart{
			Id:       order.ProductId,
			Name:     order.Name,
			Image:    order.Image,
			Price:    order.Price,
			Quantity: order.Quantity,
		}
		orderMap[order.Id].Products = append(orderMap[order.Id].Products, product)

	}

	// Convert map to slice
	result := make([]models.OrderWithProducts, 0, len(orderMap))
	for _, v := range orderMap {
		result = append(result, *v)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"orders": result,
	})

}

// @Summary edit order status
// @Security ApiKeyAuth
// @Tags Admin
// @Accept json
// @Produce json
// @Param id path string true "order id"
// @Param status body models.Status true "new status"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /api/admin/orders/{id} [put]
func (h *Handler) updateOrderStatus(c *gin.Context) {
	orderId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	var status models.Status
	if err = c.BindJSON(&status); err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	if err = h.services.Order.Update(orderId, status); err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
