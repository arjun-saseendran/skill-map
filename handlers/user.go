package handlers

import (
	"fmt"
	"net/http"

	"github.com/arjun-saseendran/skill-map/dto"
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
	userGroup.POST("", handler.CreateUser)
	userGroup.GET(":id/", handler.UserDetail)
	userGroup.DELETE(":id/", handler.DeleteUser)
}

func (handler *UserHandler) ListUser(ctx *gin.Context) {
	allUsers, err := handler.userService.List()
	if err != nil {
		dto.BadResponse(ctx, "failed to get users data")
		return
	}
	ctx.JSON(http.StatusOK, allUsers)
}

func (handler *UserHandler) CreateUser(ctx *gin.Context) {
	userData := dto.NewCreateUserInput()

	err := ctx.BindJSON(&userData)

	if err != nil {
		dto.BadResponse(ctx, "failed to bind user data")
		return
	}
	newUser, err := handler.userService.Create(userData)

	if err != nil {
		dto.BadResponse(ctx, "failed to create user")
		return
	}
	ctx.JSON(http.StatusOK, newUser)
}

func (handler *UserHandler) UserDetail(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")

	if !ok {
		fmt.Println("invalid user id")
		return
	}
	user, err := handler.userService.Get(id)
	if err != nil {
		dto.BadResponse(ctx, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (handler *UserHandler) DeleteUser(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")

	if !ok {
		dto.BadResponse(ctx, "invalid user id")
		return
	}
	err := handler.userService.Delete(id)
	if err != nil {
		dto.BadResponse(ctx, err.Error())
		return
	}
	dto.SuccessResponse(ctx, "deleted user")

}

func (handler *UserHandler) UpdateUser(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		dto.BadResponse(ctx, "invalid user id")
		return
	}
	userData := dto.NewUserUpdateInput()

	err := ctx.BindJSON(&userData)

	if err != nil {
		dto.BadResponse(ctx, "failed to bind user data")
		return
	}
	updatedUserData, err := handler.userService.Update(id, userData)
	if err != nil {
		dto.BadResponse(ctx, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, updatedUserData)
}
