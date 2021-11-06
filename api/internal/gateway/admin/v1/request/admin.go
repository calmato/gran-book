package request

import "github.com/calmato/gran-book/api/internal/gateway/entity"

type CreateAdminRequest struct {
	LastName             string      `json:"lastName"`
	FirstName            string      `json:"firstName"`
	LastNameKana         string      `json:"lastNameKana"`
	FirstNameKana        string      `json:"firstNameKana"`
	Username             string      `json:"username"`
	Role                 entity.Role `json:"role"`
	Email                string      `json:"email"`
	Password             string      `json:"password"`
	PasswordConfirmation string      `json:"passwordConfirmation"`
}

type UpdateAdminContactRequest struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

type UpdateAdminProfileRequest struct {
	LastName      string      `json:"lastName"`
	FirstName     string      `json:"firstName"`
	LastNameKana  string      `json:"lastNameKana"`
	FirstNameKana string      `json:"firstNameKana"`
	Username      string      `json:"username"`
	Role          entity.Role `json:"role"`
	ThumbnailURL  string      `json:"thumbnailUrl"`
}

type UpdateAdminPasswordRequest struct {
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}
