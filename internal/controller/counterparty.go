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

// CreateCounterparty godoc
// @Summary Создать контрагента
// @Description Создает контрагента в базе
// @Tags counterparties
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param user body models.PostCounterparty true "контрагент"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /counterparties [post]
func CreateCounterparty(c *gin.Context) {
	// парсим в JSON
	var postCp models.PostCounterparty
	if err := c.ShouldBindJSON(&postCp); err != nil {
		HandleError(c, err)
		return
	}

	// проверяем валидации
	if postCp.Name == "" ||
		postCp.Contact == "" ||
		postCp.Phone == "" ||
		postCp.Email == "" {
		logger.Error.
			Printf("[controller] CreateCounterparty(): type in not provider/recipient or name, contact, phone or email is not filled: %s\n",
				errs.ErrBadRequestBody.Error(),
			)
		err := errors.Join(errors.New("type in not provider/recipient or name, contact, phone or email is not filled: "),
			errs.ErrBadRequestBody,
		)
		HandleError(c, err)
		return
	}
	cp := models.PostCounterpartyToCounterparty(postCp)

	// создаем
	cp, err := service.CreateCounterparty(cp)
	if err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusCreated, cp)
}

// GetCounterpartyByID godoc
// @Summary Получить контрагента по ID
// @Description Возвращает контрагента по его ID
// @Tags counterparties
// @Security ApiKeyAuth
// @Produce json
// @Param id path int true "ID контрагента"
// @Success 200 {object} models.Counterparty
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /counterparties/{id} [get]
func GetCounterpartyByID(c *gin.Context) {
	// достаем id
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err = errors.Join(err, errors.New("invalid id"), errs.ErrValidationFailed)
		HandleError(c, err)
		return
	}

	// достаем данные с бд
	cp, err := service.GetcounterpartyByID(id)
	if err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusOK, cp)
}

// GetAllCounterparties godoc
// @Summary Получить всех контрагентов
// @Description Возвращает список всех контрагентов в бд
// @Tags counterparties
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {array} models.Counterparty
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /counterparties [get]
func GetAllCounterparties(c *gin.Context) {
	// достаем данные с бд
	cp, err := service.GetAllCounterparties()
	if err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusOK, cp)
}

// UpdateCounterpartyByID godoc
// @Summary Изменить контрагента
// @Description Изменение данных контрагента
// @Tags counterparties
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "ID контрагента"
// @Param operation body models.PostCounterparty true "контрагент"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /counterparties/{id} [patch]
func UpdateCounterpartyByID(c *gin.Context) {
	// достаем id
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err = errors.Join(err, errors.New("invalid id"), errs.ErrValidationFailed)
		HandleError(c, err)
		return
	}

	// парсим в JSON
	var PostCp models.PostCounterparty
	if err = c.ShouldBindJSON(&PostCp); err != nil {
		HandleError(c, err)
		return
	}
	cp := models.PostCounterpartyToCounterparty(PostCp)

	// пытаемся менять данные
	if err = service.UpdateCounterpartyByID(id, cp); err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusOK, gin.H{
		"message": "counterparty data updated successfully",
	})
}
