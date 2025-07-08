package errs

import "errors"

var (
	// ServerError
	ErrSomethingWentWrong = errors.New("something went wrong")

	// BadRequest
	ErrBadRequestBody            = errors.New("bad request body")
	ErrBadRequestQuery           = errors.New("bad request query")
	ErrValidationFailed          = errors.New("validation failed")
	ErrUserAlreadyExists         = errors.New("user already exists")
	ErrAlreadyExists             = errors.New("already exists")
	ErrProductAlreadyExists      = errors.New("product already exists")
	ErrInvalidCharacter          = errors.New("invalid character")
	ErrCounterpartyAlreadyExists = errors.New("counterparty already exists")
	ErrCellAlreadyExists         = errors.New("cell already exists")
	ErrUserDeactive              = errors.New("user is deactivated")

	// NotFound
	ErrNotFound       = errors.New("not found")
	ErrUserNotFound   = errors.New("user not found")
	ErrUserIDNotFound = errors.New("user ID Not Found")

	// Forbidden
	ErrNoPermissionsToEditUser = errors.New("no permissions to edit user")

	// Unauthorized
	ErrIncorrectUsernameOrPassword = errors.New("incorrect username or password")
)
