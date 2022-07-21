package utils

import (
	"net/mail"

	"github.com/bshadmehr-com/libs/errs"
)

func ValidateEmail(email string) *errs.AppError {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errs.NewBadRequestError("Invalid email address")
	}
	return nil
}
