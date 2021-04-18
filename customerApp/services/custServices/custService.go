package services

import (

	"github.com/narendra121/go-gin-mysql/customerApp/domain/customers"
)

var (
	CustomersService customersInterface = &customersService{}
)

type customersService struct{}
type customersInterface interface {
	CreateCustomer(customers.Customer) (*customers.Customer, error)
	SearchCustomer(int64) (*customers.Customer, error)
	UpdateCustomer(int64,customers.Customer) (*customers.Customer, error)
	DeleteCustomer(int64) (*customers.Customer, error)
}

func (s *customersService) CreateCustomer(customer customers.Customer) (*customers.Customer, error) {
	if err := customer.Save(); err != nil {
		return nil, err
	}
	return &customer, nil
}

func (s *customersService) SearchCustomer(customerId int64) (*customers.Customer, error) {
	result := &customers.Customer{Id: customerId}
	if getErr := result.Get(); getErr != nil {
		return nil, getErr
	}
	return result, nil
}

func (s *customersService)UpdateCustomer(customerId int64,cust customers.Customer) (*customers.Customer, error){
		result:=&customers.Customer{Id: customerId }
		if getErr:=result.Get();getErr!=nil{
			return nil,getErr
		}
		result.FirstName=cust.FirstName
		result.LastName=cust.LastName
		result.Email=cust.Email
		result.Salary=cust.Salary
		if err:=result.Update();err!=nil{
			return nil,err
		}
		return result,nil
}

func (s *customersService) 	DeleteCustomer(customerId int64) (*customers.Customer, error){
	result := &customers.Customer{Id: customerId}
	if getErr := result.Get(); getErr != nil {
		return nil, getErr
	}
	if err:=result.Delete();err!=nil{
		return nil, err

	}
	return result,nil
}