package response

import (
	"eniqlo/internal/entity"
)

type CustomerRegister struct {
	UserID      string `json:"userId"`
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
}

type Customer struct {
	UserID      string `json:"userId"`
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
}

func MapCustomerEntityToResponse(e *entity.Customer) *Customer {
	return &Customer{
		e.ID,
		e.PhoneNumber,
		e.Name,
	}
}

func MapCustomerListEntityToListResponse(e []*entity.Customer) []*Customer {
	var resp []*Customer
	for _, v := range e {
		resp = append(resp, MapCustomerEntityToResponse(v))
	}
	return resp
}
