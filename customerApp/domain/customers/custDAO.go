package customers

import (
	"github.com/narendra121/go-gin-mysql/customerApp/datasource/customerDB"
)

const (
	queryCreateCustomer = "INSERT INTO customer(first_name,last_name,email,salary)VALUES(?,?,?,?);"
	querySearchCustomer = "SELECT id,first_name,last_name,email,salary FROM customer WHERE id=?;"
	queryUpdateCustomer = "UPDATE customer SET first_name=?,last_name=?,email=?,salary=? WHERE id=?;"
	queryDeleteCustomer = "DELETE FROM customer WHERE id=?;"
)

func (customer *Customer) Save() error {
	stmt, sErr := customerDB.CreateDB().Prepare(queryCreateCustomer)
	if sErr != nil {
		return sErr
	}
	defer stmt.Close()
	inResult, exErr := stmt.Exec(customer.FirstName, customer.LastName, customer.Email, customer.Salary)
	if exErr != nil {
		return exErr
	}
	customerId, err := inResult.LastInsertId()
	if err != nil {
		return err
	}
	customer.Id = customerId
	return nil
}

func (customer *Customer) Get() error {
	stmt, err := customerDB.CreateDB().Prepare(querySearchCustomer)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result := stmt.QueryRow(customer.Id)
	if rowsErr := result.Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.Email, &customer.Salary); rowsErr != nil {

		return rowsErr
	}
	return nil
}

func (customer *Customer) Update() error {
	stmt, sErr := customerDB.CreateDB().Prepare(queryUpdateCustomer)
	if sErr != nil {
		return sErr
	}
	_, err := stmt.Exec(customer.FirstName, customer.LastName, customer.Email, customer.Salary, customer.Id)
	if err != nil {
		return err
	}
	return nil
}

func (customer *Customer) Delete() error {
	stmt, sErr := customerDB.CreateDB().Prepare(queryDeleteCustomer)
	if sErr != nil {
		return sErr
	}
	_, err := stmt.Exec(customer.Id)
	if err != nil {
		return err
	}
	return nil
}
