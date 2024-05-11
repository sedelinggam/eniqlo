package request

type CustomerRegister struct {
	PhoneNumber string `json:"phoneNumber" validate:"required,min=10,max=16"`
	Name        string `json:"name" validate:"required,min=5,max=50"`
}

type Customer struct {
	PhoneNumber *string
	Name        *string
}
