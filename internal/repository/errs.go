package repository

import (
	"database/sql"
	"errors"
	"warehouse/internal/errs"

	"github.com/lib/pq"
)

func translateError(err error) error {
	if err == nil {
		return nil
	} else if errors.Is(err, sql.ErrNoRows) {
		return errs.ErrNotFound
	} else if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
		err = errors.Join(errors.New("unique violation"), errs.ErrAlreadyExists)
		return err
	} else if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23503" {
		err = errors.Join(errors.New("foreign key violation"), errs.ErrAlreadyExists)
		return err
	} else if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23502" {
		err = errors.Join(errors.New("not null violation"), errs.ErrBadRequestBody)
		return err
	} else if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "22001" {
		err = errors.Join(errors.New("string data right truncation"), errs.ErrBadRequestBody)
		return err
	} else {
		return err
	}
}
