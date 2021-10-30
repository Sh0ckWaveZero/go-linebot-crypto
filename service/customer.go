package service

type CustomerResponse struct {
	CustomerID int    `json:"customer_id"`
	Name       string `json:"name"`
	Status     string `json:"status"`
}

type CustomerService interface {
	GetCustomers() ([]CustomerResponse, error)
	GetCustomer(customerID int) (*CustomerResponse, error)
}
