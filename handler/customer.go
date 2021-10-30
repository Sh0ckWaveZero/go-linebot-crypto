package handler

import (
	"encoding/json"
	"fmt"
	"go-linebot-crypto/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type customerHandler struct {
	custSvr service.CustomerService
}

func NewCustomerHandler(custSvr service.CustomerService) customerHandler {
	return customerHandler{custSvr: custSvr}
}

func (h customerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.custSvr.GetCustomers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (h customerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])
	customer, err := h.custSvr.GetCustomer(customerID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(customer)
}
