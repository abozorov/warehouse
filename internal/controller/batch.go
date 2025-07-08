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

// CreateBatch godoc
// @Summary Создать партию
// @Description Создается партия поставки
// @Tags batches
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param user body models.PostBatch true "Данные партии"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /batches [post]
func CreateBatch(c *gin.Context) {
	// парсинг из json
	var postB models.PostBatch
	if err := c.ShouldBindJSON(&postB); err != nil {
		HandleError(c, err)
		return
	}

	// проверка валидаций
	if (postB.Type != "in" && postB.Type != "out") ||
		postB.CounterpartyName == "" ||
		postB.Article == "" ||
		postB.Quantity <= 0 {
		logger.Error.
			Printf("[controller] CreateBatch(): one or more fields are filled in incorrectly: %s\n", errs.ErrBadRequestBody.Error())
		err := errors.Join(errors.New("one or more fields are filled in incorrectly"), errs.ErrBadRequestBody)
		HandleError(c, err)
		return
	}
	b := models.PostBatchToBatch(postB)
	b.Username = c.GetString(userUsernameCtx)

	// попытка протолкнуть в сервис
	err := service.CreateBatch(b)
	if err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusCreated, gin.H{
		"message": "batch created successfully",
	})
}

// GetAllBatches godoc
// @Summary Получить все партии
// @Description Возвращает список всех партий в бд
// @Tags batches
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {array} models.Batch
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /batches [get]
func GetAllBatches(c *gin.Context) {
	// достаем данные с бд
	batches, err := service.GetAllBatches()
	if err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusOK, batches)
}

// GetBatchByID godoc
// @Summary Получить партию по ID
// @Description Возвращает данные партии по его ID
// @Tags batches
// @Security ApiKeyAuth
// @Produce json
// @Param id path int true "ID партии"
// @Success 200 {object} models.Batch
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /batches/{id} [get]
func GetBatchByID(c *gin.Context) {
	// доствем if
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err = errors.Join(err, errors.New("invalid id"), errs.ErrValidationFailed)
		HandleError(c, err)
		return
	}

	// достаем данные с бд
	batch, err := service.GetBatchByID(id)
	if err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusOK, batch)
}

func BatchStats(c *gin.Context) {
	filter, err := GetFilterQuery(c)
	if err != nil {
		HandleError(c, err)
		return
	}

	// указать какие поля должны возврашаться и с какой таблицы
	// логика фильтрации и отправка запроса
	stats, err := service.GetStats("b", filter)
	if err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusOK, stats)
}