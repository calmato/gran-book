package input

type CreateInquiry struct {
	SenderId    string `json:"sender_id" validate:"max=36"`
	Subject     string `json:"subject" validate:"required,min=1,max=64"`
	Description string `json:"description" validate:"required,min=1,max=1000"`
	Email       string `json:"email" validate:"required,email,max=256"`
}
