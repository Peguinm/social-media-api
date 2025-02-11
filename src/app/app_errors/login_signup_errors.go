package apperrors

import (
	"errors"
	"strconv"

	"sm.com/m/src/app/constants"
)

var (
	ErrIvalidPassword error = errors.New("invalid password, must be at least 8 length and contains one of each character: special, upper letter, lower letter, number")
	ErrInvalidEmail   error = errors.New("invalid email format, ex: random.email@domain.com")
	ErrInvalidName    error = errors.New("invalid name format, name can only contains letters, numbers and undercores")

	ErrLongPassword  error = errors.New("password too long, must be minor than " + strconv.FormatUint(constants.PASSWORD_MAX_LENGTH, 10) + " characters")
	ErrShortPassword error = errors.New("password too short, must be at least " + strconv.FormatUint(constants.PASSWORD_MIN_LENGTH, 10) + " characters")

	ErrShortName error = errors.New("name too short, must be at least " + strconv.FormatUint(constants.USER_NAME_MIN_LENGTH, 10) + " characters")
	ErrLongName  error = errors.New("name too long, must be minor than " + strconv.FormatUint(constants.USER_NAME_MAX_LENGTH, 10) + " characters")

	ErrEmailAlreadyInUse error = errors.New("email already in use")
	ErrNameAlreadyInUse  error = errors.New("name already in use")

	ErrWrongLogin   error = errors.New("wrong login or password")
	ErrCreatingUser error = errors.New("failed to create user")
	ErrLogin        error = errors.New("failed to login")
)
