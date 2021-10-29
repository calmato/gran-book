package request

type UpdateAuthEmailRequest struct {
	Email string `json:"email"`
}

type UpdateAuthPasswordRequest struct {
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

type UpdateAuthProfileRequest struct {
	LastName         string `json:"lastName"`
	FirstName        string `json:"firstName"`
	LastNameKana     string `json:"lastNameKana"`
	FirstNameKana    string `json:"firstNameKana"`
	Username         string `json:"username"`
	PhoneNumber      string `json:"phoneNumber"`
	ThumbnailURL     string `json:"thumbnailUrl"`
	SelfIntroduction string `json:"selfIntroduction"`
}

type RegisterAuthDeviceRequest struct {
	InstanceID string `json:"instanceId"`
}
