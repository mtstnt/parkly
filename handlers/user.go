package handlers

import "github.com/gin-gonic/gin"

type UserService interface {
	GetAllUsers()
}

type User struct {
	userService UserService
}

func NewUser(userService UserService) *User {
	return &User{
		userService,
	}
}

func (h *User) GetAllUsers(ctx *gin.Context) {
	h.userService.GetAllUsers()
}
