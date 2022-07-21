package dto

import (
	"github.com/bshadmehr-com/auth/domain"
	"github.com/bshadmehr-com/auth/utils"
	"github.com/bshadmehr-com/libs/errs"
)

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type SignupResponse struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
}

func (r SignupRequest) ToDao() domain.User {
	return domain.User{
		Email:    r.Email,
		Password: r.Password,
		Name:     r.Name,
	}
}

func (r SignupRequest) Validate() *errs.AppError {
	if err := utils.ValidateEmail(r.Email); err != nil {
		return err
	}

	if err := utils.ValidatePassword(r.Password); err != nil {
		return err
	}

	return nil
}

func UserToSignupResponse(user domain.User) *SignupResponse {
	return &SignupResponse{
		UserID: user.Id,
		Email:  user.Email,
		Name:   user.Name,
	}
}
