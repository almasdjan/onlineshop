package handler

import (
	"net/http"
	"onlineshop/models"

	"github.com/gin-gonic/gin"
)

// @Summary SignUp
// @Tags auth
// @Accept json
// @Produce json
// @Param input body models.User true "account info"
// @Success 200 {object} map[string]any
// @Failure 400 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /auth/signup [post]
func (h *Handler) signup(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	/*
		c.JSON(http.StatusOK, map[string]interface{}{
			"id": id,
		})
	*/

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
		"id":    id,
	})
}

// @Summary LogIn
// @Tags auth
// @Accept json
// @Produce json
// @Param input body models.Login true "email password"
// @Success 200 {object} map[string]any
// @Failure 400 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /auth/login [post]
func (h *Handler) login(c *gin.Context) {
	var input models.Login

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
