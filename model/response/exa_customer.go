package response

import "cms/model"

type ExaCustomerResponse struct {
	Customer model.ExaCustomer `json:"customer"`
}
