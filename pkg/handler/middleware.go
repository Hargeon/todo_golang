package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	headersParts := strings.Split(header, " ")
	if len(headersParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	// parse token
	userId, err := h.service.Authorization.ParseToken(headersParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth token")
		return
	}
	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "user id not found")
		return 0, errors.New("user id not found")
	}
	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "user id has invalid type")
		return 0, errors.New("user id has invalid type")
	}
	return idInt, nil
}