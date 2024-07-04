package app

import (
	"errors"
)

type NotFoundError string

var ErrNotFound = errors.New("[NotFound]")
var ErrTechnical = errors.New("[Technical]")
