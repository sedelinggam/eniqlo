package request

type StaffRegister struct {
	PhoneNumber string `json:"phoneNumber" validate:"required,min=10,max=16"`
	Name        string `json:"name" validate:"required,min=5,max=50"`
	Password    string `json:"password" validate:"required,min=5,max=15"`
}

type StaffLogin struct {
	PhoneNumber string `json:"phoneNumber" validate:"required,min=10,max=16"`
	Password    string `json:"password" validate:"required,min=5,max=15"`
}
