package controller

import (
	"errors"
	"log"
	"time"
	"warehouse/internal/errs"
	"warehouse/internal/models"
	"warehouse/logger"

	"github.com/gin-gonic/gin"
)

func GetFilterQuery(c *gin.Context) (query models.Filter, err error) {
	// достать query
	query.DateFrom = c.DefaultQuery("date_from", "1970-01-01T00:00:00.00Z")
	query.DateTo = c.DefaultQuery("date_to", time.Now().Format(time.RFC3339))
	query.BatchType = c.DefaultQuery("type", "all")

	log.Println(query.DateFrom, query.DateTo)

	if (query.BatchType != "in" &&
		query.BatchType != "out" &&
		query.BatchType != "all") ||
		query.DateFrom > query.DateTo {
		logger.Error.
			Printf("[controller] GetFilterQuery(): queries is not filled correctly: %s\n",
				errs.ErrBadRequestQuery.Error(),
			)
		err = errors.Join(errors.New("queries is not filled correctly: "),
			errs.ErrBadRequestQuery,
		)
		return models.Filter{}, err
	}

	return query, nil
}
