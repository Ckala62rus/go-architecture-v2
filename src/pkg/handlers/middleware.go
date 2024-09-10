package handlers

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"practice/pkg/utils"
	"strings"
)

const (
	authorizationHeader       = "Authorization"
	userCtx                   = "userId"
	AuthenticationTokenHeader = "TokenHeader"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	tokenHeader := headerParts[1]

	userId, err := h.services.Authorization.ParseToken(tokenHeader)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	if userId == 0 {
		newErrorResponse(c, http.StatusUnauthorized, "User unauthorized")
		return
	}

	rdb := utils.RedisDb
	ctx := context.Background()
	token, err := rdb.GetToken(ctx, tokenHeader)

	if len(token) == 0 {
		newErrorResponse(c, http.StatusUnauthorized, "Token has expired")
	}

	c.Set(userCtx, userId)
	c.Set(AuthenticationTokenHeader, tokenHeader)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id is of invalid type")
		return 0, errors.New("user id not found")
	}

	return idInt, nil
}

func getAuthenticationHeader(c *gin.Context) (string, error) {
	token, ok := c.Get(AuthenticationTokenHeader)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "not authentication token in header")
		return "", errors.New("authentication token not found in header")
	}

	tokenPerform, ok := token.(string)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "tokenPerform is of invalid type")
		return "", errors.New("tokenPerform is of invalid type")
	}
	return tokenPerform, nil
}
