package dto

import (
	"github.com/bshadmehr-com/auth/utils"
	"github.com/bshadmehr-com/libs/errs"
)

type ChangePassRequest struct {
	OldPass string `json:"old_password"`
	NewPass string `json:"new_password"`
}

type ChangePassResponse struct {
	Success      bool   `json:"success"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (r ChangePassRequest) Validate() *errs.AppError {
	if err := utils.ValidatePassword(r.NewPass); err != nil {
		return err
	}

	return nil
}
