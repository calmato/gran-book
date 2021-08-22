package v1

import "github.com/calmato/gran-book/api/gateway/native/internal/entity"

type CreateAuthRequest struct {
	Username             string `json:"username"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

type UpdateAuthEmailRequest struct {
	Email string `json:"email"`
}

type UpdateAuthPasswordRequest struct {
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

type UpdateAuthProfileRequest struct {
	Username         string        `json:"username"`
	Gender           entity.Gender `json:"gender"`
	ThumbnailURL     string        `json:"thumbnailUrl"`
	SelfIntroduction string        `json:"selfIntroduction"`
}

type UpdateAuthAddressRequest struct {
	LastName      string `json:"lastName"`
	FirstName     string `json:"firstName"`
	LastNameKana  string `json:"lastNameKana"`
	FirstNameKana string `json:"firstNameKana"`
	PhoneNumber   string `json:"phoneNumber"`
	PostalCode    string `json:"postalCode"`
	Prefecture    string `json:"prefecture"`
	City          string `json:"city"`
	AddressLine1  string `json:"addressLine1"`
	AddressLine2  string `json:"addressLine2"`
}

type RegisterAuthDeviceRequest struct {
	InstanceID string `json:"instanceId"`
}
