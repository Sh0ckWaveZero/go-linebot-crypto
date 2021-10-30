package repository

import "time"

type Customer struct {
	CustomerID  int       `db:"customer_id"`
	Name        string    `db:"name"`
	DateOfBirth time.Time `db:"date_of_birth"`
	City        string    `db:"city"`
	ZipCode     string    `db:"zipcode"`
	Status      string    `db:"status"`
}

type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetById(int) (*Customer, error)
}

func (r customerRepositoryDB) GetAll() ([]Customer, error) {
	customes := []Customer{}
	query := "SELECT customer_id, name, date_of_birth, city, zipcode, status FROM customers"
	err := r.db.Select(&customes, query)
	if err != nil {
		return nil, err
	}
	return customes, nil
}

func (r customerRepositoryDB) GetById(id int) (*Customer, error) {
	customer := Customer{}
	query := "SELECT customer_id, name, date_of_birth, city, zipcode, status FROM customers where customer_id = $1"
	err := r.db.Get(&customer, query, id)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
