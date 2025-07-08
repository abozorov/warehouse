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

// CreateProduct godoc
// @Summary Создать товар
// @Description Создается товар в базе
// @Tags products
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param user body models.PostProduct true "товар"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /products [post]
func CreateProduct(c *gin.Context) {
	// парсим в JSON
	var postP models.PostProduct

	if err := c.ShouldBindJSON(&postP); err != nil {
		HandleError(c, err)
		return
	}

	// проверяем валидации
	if postP.Article == "" {
		logger.Error.
			Printf("[controller] CreateProduct(): article is not filled in: %s\n", errs.ErrBadRequestBody.Error())
		err := errors.Join(errors.New("article is not filled in: "), errs.ErrBadRequestBody)
		HandleError(c, err)
		return
	}
	p := models.PostProductToProduct(postP)

	// создаем
	p, err := service.CreateProduct(p)
	if err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusCreated, p)

}

// GetProductByID godoc
// @Summary Получить товар по ID
// @Description Возвращает товар по его ID
// @Tags products
// @Security ApiKeyAuth
// @Produce json
// @Param id path int true "ID товара"
// @Success 200 {object} models.Product
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /products/{id} [get]
func GetProductByID(c *gin.Context) {
	// достаем id
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err = errors.Join(err, errors.New("invalid id"), errs.ErrValidationFailed)
		HandleError(c, err)
		return
	}

	// достаем данные с бд
	product, err := service.GetProductByID(id)
	if err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusOK, product)

}

// GetAllProducts godoc
// @Summary Получить все товары
// @Description Возвращает список всех товаров в бд
// @Tags products
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {array} models.Product
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /products [get]
func GetAllProducts(c *gin.Context) {
	// достаем данные с бд
	products, err := service.GetAllProducts()
	if err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusOK, products)
}

// UpdateProductByID godoc
// @Summary Изменить товар
// @Description Изменение данных товара
// @Tags products
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "ID ячейки"
// @Param operation body models.PostProduct true "товар"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /products/{id} [patch]
func UpdateProductByID(c *gin.Context) {
	// достаем id
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err = errors.Join(err, errors.New("invalid id"), errs.ErrValidationFailed)
		HandleError(c, err)
		return
	}

	// парсим в JSON
	var postP models.PostProduct

	if err := c.ShouldBindJSON(&postP); err != nil {
		HandleError(c, err)
		return
	}
	p := models.PostProductToProduct(postP)

	// пытаемся менять данные
	if err = service.UpdateProductByID(id, p); err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusOK, gin.H{
		"message": "product data updated successfully",
	})
}

func ProductStats(c *gin.Context) {
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
	stats, err := service.GetStats("p", filter)
	if err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusOK, stats)
}
