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

// CreateCell godoc
// @Summary Создать ячейку
// @Description Создается ячейка хранения в базе
// @Tags cells
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param cell body models.PostCell true "Данные ячейки"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /cells [post]
func CreateCell(c *gin.Context) {
	// парсим в JSON
	var postC models.PostCell

	if err := c.ShouldBindJSON(&postC); err != nil {
		HandleError(c, err)
		return
	}

	// проверяем валидации
	if postC.Zone == "" ||
		postC.Row == 0 ||
		postC.AdressCode == "" {
		logger.Error.
			Printf("[controller] CreateCell(): zone, row or adress code is not filled in: %s\n", errs.ErrBadRequestBody.Error())
		err := errors.Join(errors.New("zone, row or adress code is not filled in: "), errs.ErrBadRequestBody)
		HandleError(c, err)
		return
	}
	cell := models.PostCellToCell(postC)

	// создаем
	cell, err := service.CreateCell(cell)
	if err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusCreated, cell)

}

// GetCellByID godoc
// @Summary Получить ячейку по ID
// @Description Возвращает ячейку по его ID
// @Tags cells
// @Security ApiKeyAuth
// @Produce json
// @Param id path int true "ID ячейки"
// @Success 200 {object} models.Cell
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /cells/{id} [get]
func GetCellByID(c *gin.Context) {
	// достаем id
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err = errors.Join(err, errors.New("invalid id"), errs.ErrValidationFailed)
		HandleError(c, err)
		return
	}

	// достаем данные с бд
	Cell, err := service.GetCellByID(id)
	if err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusOK, Cell)

}

// GetAllCells godoc
// @Summary Получить все ячейка
// @Description Возвращает список всех ячеек хранения в бд
// @Tags cells
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {array} models.Cell
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /cells [get]
func GetAllCells(c *gin.Context) {
	// достаем данные с бд
	cells, err := service.GetAllCells()
	if err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusOK, cells)
}

// UpdateCellByID godoc
// @Summary Изменить ячейки
// @Description Изменение данных ячейки хранения
// @Tags cells
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "ID ячейки"
// @Param cell body models.PostCell true "ячейка"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /cells/{id} [patch]
func UpdateCellByID(c *gin.Context) {
	// достаем id
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err = errors.Join(err, errors.New("invalid id"), errs.ErrValidationFailed)
		HandleError(c, err)
		return
	}

	// парсим в JSON
	var postC models.PostCell

	if err := c.ShouldBindJSON(&postC); err != nil {
		HandleError(c, err)
		return
	}
	cell := models.PostCellToCell(postC)

	// пытаемся менять данные
	if err = service.UpdateCellByID(id, cell); err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusOK, gin.H{
		"message": "cell data updated successfully",
	})
}
