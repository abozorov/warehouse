package controller

import (
	"net/http"
	"warehouse/internal/errs"
	"warehouse/internal/models"
	"warehouse/internal/service"
	"warehouse/utils"

	"github.com/gin-gonic/gin"
)

// SignIn godoc
// @Summary Авторизация пользователя
// @Description Вход по логину и паролю
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body models.UserSignIn true "Учётные данные"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /sign-in [post]
func SignIn(c *gin.Context) {
	// получить идентификатор и пароль
	var u models.UserSignIn
	if err := c.ShouldBindJSON(&u); err != nil {
		HandleError(c, err)
		return
	}

	// отправить в бд запрос на проверку есть ли такой
	// пользователь с таким паролем
	user, err := service.GetUserByUsernameAndPassword(u.Username, u.Password)
	if err != nil {
		HandleError(c, err)
		return
	}

	if !user.Active {
		HandleError(c, errs.ErrUserDeactive)
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": token,
	})
}
