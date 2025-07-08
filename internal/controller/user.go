package controller

import (
	"errors"
	"net/http"
	"strconv"
	"warehouse/internal/errs"
	"warehouse/internal/models"
	"warehouse/internal/service"
	"warehouse/logger"

	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @Summary Создать пользователя
// @Description Создается пользователя в базе
// @Tags users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param user body models.User true "пользователь"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Router /users [post]
func CreateUser(c *gin.Context) {
	// парсим в JSON
	var u models.User

	if err := c.ShouldBindJSON(&u); err != nil {
		HandleError(c, err)
		return
	}

	// проверяем валидации
	if u.Username == "" || u.Password == "" || u.FullName == "" {
		logger.Error.
			Printf("[controller] CreateUser(): username, password or full name is not filled in: %s\n",
				errs.ErrBadRequestBody.Error(),
			)
		err := errors.Join(errors.New("username, password or full name is not filled in: "),
			errs.ErrBadRequestBody,
		)
		HandleError(c, err)
		return
	}

	// создаем юзера
	user, err := service.CreateUser(u)
	if err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusCreated, user)
}

// GetUserByID godoc
// @Summary Получить пользователя по ID
// @Description Возвращает пользователя по его ID
// @Tags users
// @Security ApiKeyAuth
// @Produce json
// @Param id path int true "ID пользователя"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /users/{id} [get]
func GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err = errors.Join(err, errors.New("invalid id"), errs.ErrValidationFailed)
		HandleError(c, err)
		return
	}

	user, err := service.GetUserByID(id)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetAllUsers godoc
// @Summary Получить всех пользователей
// @Description Возвращает список всех пользователей в бд
// @Tags users
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {array} models.User
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Router /users [get]
func GetAllUsers(c *gin.Context) {
	users, err := service.GetAllUsers()
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

// UpdateUserByID godoc
// @Summary Изменить данные пользователя
// @Description Изменение данные пользователя
// @Tags users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "ID рользователя"
// @Param operation body models.User true "пользователь"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /users/{id} [patch]
func UpdateUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err = errors.Join(err, errors.New("invalid id"), errs.ErrValidationFailed)
		HandleError(c, err)
		return
	}

	var u models.User
	if err = c.ShouldBindJSON(&u); err != nil {
		HandleError(c, err)
		return
	}

	if err = service.UpdateUserByID(id, u); err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "user data updated successfully",
	})
}

func UserStats(c *gin.Context) {
	// достать id(params)
	idStr := c.Param("id")
	_, err := strconv.Atoi(idStr)
	if err != nil {
		err = errors.Join(err, errors.New("invalid id"), errs.ErrValidationFailed)
		HandleError(c, err)
		return
	}
	
	filter, err := GetFilterQuery(c)
	if err != nil {
		HandleError(c, err)
		return
	}
	filter.ID = idStr

	// указать какие поля должны возврашаться и с какой таблицы
	// логика фильтрации и отправка запроса
	stats, err := service.GetStats("u", filter)
	if err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusOK, stats)
}
