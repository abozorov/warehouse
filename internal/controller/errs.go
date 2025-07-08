package controller

import (
	"errors"
	"fmt"
	"net/http"
	"warehouse/internal/errs"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	if err == nil {
		return
	} else if errors.Is(err, errs.ErrValidationFailed) ||
		errors.Is(err, errs.ErrUserAlreadyExists) ||
		errors.Is(err, errs.ErrBadRequestBody) ||
		errors.Is(err, errs.ErrBadRequestQuery) ||
		errors.Is(err, errs.ErrAlreadyExists) ||
		errors.Is(err, errs.ErrProductAlreadyExists) ||
		errors.Is(err, errs.ErrCounterpartyAlreadyExists) ||
		errors.Is(err, errs.ErrCellAlreadyExists) ||
		errors.Is(err, errs.ErrUserDeactive) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else if errors.Is(err, errs.ErrUserNotFound) ||
		errors.Is(err, errs.ErrNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	} else if errors.Is(err, errs.ErrIncorrectUsernameOrPassword) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
	} else if errors.Is(err, errs.ErrNoPermissionsToEditUser) {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("something went wrong: %s", err.Error()),
		})
	}
}
