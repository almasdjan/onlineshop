package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		NewErrorResponce(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		NewErrorResponce(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	//parse token
	userId, err := h.services.Authorization.ParseToken((headerParts[1]))
	if err != nil {
		NewErrorResponce(c, http.StatusUnauthorized, err.Error())
	}

	isAdmin, err := h.services.Authorization.IsAdmin(userId)
	if err != nil {
		NewErrorResponce(c, http.StatusUnauthorized, err.Error())
	}

	c.Set(userCtx, userId)
	c.Set("isAdmin", isAdmin)

}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		NewErrorResponce(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}
	idInt, ok := id.(int)
	if !ok {
		NewErrorResponce(c, http.StatusInternalServerError, "user id is of invalid type")
		return 0, errors.New("user id not found")
	}
	return idInt, nil

}

func (h *Handler) checkAdmin(c *gin.Context) {
	isAdmin, ok := c.Get("isAdmin")
	if !ok {
		NewErrorResponce(c, http.StatusInternalServerError, "user role not found")
		return
	}
	isAdminBool, ok := isAdmin.(bool)
	if !ok {
		NewErrorResponce(c, http.StatusInternalServerError, "user role is of invalid type")
		return
	}

	if !isAdminBool {
		NewErrorResponce(c, http.StatusInternalServerError, "user is not admin")
		return
	}

	c.Next()

}
