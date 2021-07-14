package input

type CreateInquiry struct {
	SenderID    string `json:"senderId" validate:"required,max=36"`
	Subject     string `json:"subject" validate:"required,max=64"`
	Description string `json:"description" validate:"required,max=1000"`
	Email       string `json:"email" validate:"required,email,max=256"`
}
