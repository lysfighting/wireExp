package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wireExp/internal/api/service"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Index(ctx *gin.Context) {
	text := h.userService.Index(ctx)
	ctx.JSON(http.StatusOK, text)
	return
}
