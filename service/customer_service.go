package service

import (
	"database/sql"
	"errors"
	"go-linebot-crypto/repository"
	"log"
)

type customerService struct {
	cusRepo repository.CustomerRepository
}

func NewCustomerService(cusRepo repository.CustomerRepository) CustomerService {
	return customerService{cusRepo: cusRepo}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.cusRepo.GetAll()
	if err != nil {
		log.Panicln(err)
		return nil, err
	}
	cusResponses := []CustomerResponse{}
	for _, customer := range customers {
		cusResponse := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		cusResponses = append(cusResponses, cusResponse)
	}
	return cusResponses, nil
}

func (s customerService) GetCustomer(customerID int) (*CustomerResponse, error) {
	customer, err := s.cusRepo.GetById(customerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("customer not found")
		}
		log.Panicln(err)
		return nil, err
	}
	cusResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &cusResponse, nil
}
