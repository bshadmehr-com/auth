package dto

import (
	"github.com/bshadmehr-com/auth/utils"
	"github.com/bshadmehr-com/libs/errs"
)

type ForgetPassRequest struct {
	Email   string `json:"email"`
	Otp     string `json:"otp"`
	NewPass string `json:"new_password"`
}

type ForgetPassResponse struct {
	Success      bool   `json:"success"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (r ForgetPassRequest) Validate() *errs.AppError {
	if err := utils.ValidateEmail(r.Email); err != nil {
		return err
	}

	if err := utils.ValidatePassword(r.NewPass); err != nil {
		return err
	}

	return nil
}
