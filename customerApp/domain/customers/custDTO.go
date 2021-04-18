package customers

import "strings"

type Customer struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Salary    int64 `json:"salary"`
}

func (customer *Customer) Validate() {
	customer.FirstName = strings.TrimSpace(customer.FirstName)
	customer.LastName=strings.TrimSpace(customer.LastName)
	customer.Email=strings.TrimSpace(customer.Email)

}