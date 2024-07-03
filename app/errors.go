package app

import "fmt"

func NewTechnicalError(err error, msg string) error {
	return fmt.Errorf("[Technical] %s: %w", msg, err)
}
