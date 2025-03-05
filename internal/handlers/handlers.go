package handlers

import (
	"net/http"
	"playground/internal/models"

	"github.com/gin-gonic/gin"
)

// @Summary CreateUser
// @Tag user
// @ID create-user
// @Accept json
// @Product json
// @Success 200 {string} string ok
// @Router /users [post]
func CreateUser(ctx *gin.Context) {
	var user models.User
	ctx.BindJSON(&user)

	user.Create(&user)
	ctx.JSON(http.StatusCreated, user)
}

// @Summary GetUsers
// @Tag user
// @ID find-users
// @Product json
// @Success 200 {string} string ok
// @Router /users [get]
func FindAllUsers(ctx *gin.Context) {
	var user models.User

	users, err := user.FindAll()
	if err != nil {
		ctx.AbortWithStatus(500)
		return
	}

	ctx.JSON(http.StatusOK, users)

}
