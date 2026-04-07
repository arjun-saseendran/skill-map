package handlers

import (
	"github.com/arjun-saseendran/skill-map/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	groupName   string
	userService service.UserService
}

func NewUserHandlerFrom(userService service.UserService) *UserHandler {
	return &UserHandler{"api/users", userService}
}

func (handler *UserHandler) RegisterEndpoints(r *gin.Engine) {
	userGroup := r.Group(handler.groupName)

	userGroup.GET("", handler.ListUser)
}

func (handler *UserHandler) ListUser(ctx *gin.Context) {
	allUsers, err := handler.userService.List()

}
