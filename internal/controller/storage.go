package controller

import (
	"net/http"
	"warehouse/internal/service"

	"github.com/gin-gonic/gin"
)

// GetStorageByAdressCode godoc
// @Summary Получить место хранения
// @Description Возвращает место хранения по его adressCode
// @Tags storages
// @Security ApiKeyAuth
// @Produce json
// @Param adressCode path string true "AdressCode места хранения"
// @Success 200 {object} models.Storage
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /storages/{adressCode} [get]
func GetStorageByAdressCode(c *gin.Context) {
	// достаем adressCode
	adressCode := c.Param("adressCode")

	// достаем данные с бд
	storage, err := service.GetStorageByAdressCode(adressCode)
	if err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusOK, storage)

}

// GetAllStorages godoc
// @Summary Получить все места хранения
// @Description Возвращает список всех мест с товарами
// @Tags storages
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {array} models.Storage
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /storages [get]
func GetAllStorages(c *gin.Context) {
	// достаем данные с бд
	storages, err := service.GetAllStorages()
	if err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusOK, storages)
}
