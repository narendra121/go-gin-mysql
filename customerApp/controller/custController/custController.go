package custController

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/narendra121/go-gin-mysql/customerApp/domain/customers"
	services "github.com/narendra121/go-gin-mysql/customerApp/services/custServices"
)

func Create(c *gin.Context) {
	var customer customers.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid json body")
		return
	}
	result, sErr := services.CustomersService.CreateCustomer(customer)
	if sErr != nil {
		c.JSON(http.StatusInternalServerError, sErr.Error())
		return
	}
	c.JSON(http.StatusOK, result)
}

func Search(c *gin.Context) {
	customerId, idErr := strconv.ParseInt(c.Param("custid"), 10, 64)
	if idErr != nil {
		c.JSON(http.StatusBadRequest, "Invalid Customer Id ")
		return
	}

	customer, getErr := services.CustomersService.SearchCustomer(customerId)
	if getErr != nil {
		c.JSON(http.StatusNotFound, getErr.Error())
		return
	}

	c.JSON(http.StatusOK, customer)
}



func Update(c *gin.Context) {
	customerId, idErr := strconv.ParseInt(c.Param("custid"), 10, 64)
	if idErr != nil {
		c.JSON(http.StatusBadRequest, "Invalid Customer Id ")
		return
	}
	var customer customers.Customer
	if jsonError:=c.ShouldBindJSON(&customer);jsonError!=nil{
		c.JSON(http.StatusBadRequest,"Invalid json body")
		return
	}
	result, updErr := services.CustomersService.UpdateCustomer(customerId,customer)
	if updErr != nil {
		c.JSON(http.StatusNotFound,updErr.Error() )
		return
	}
		c.JSON(http.StatusOK,result)
}


func Delete(c *gin.Context) {
	customerId, idErr := strconv.ParseInt(c.Param("custid"), 10, 64)
	if idErr != nil {
		c.JSON(http.StatusBadRequest, "Invalid Customer Id ")
		return
	}
	result, delErr := services.CustomersService.DeleteCustomer(customerId)
	if delErr != nil {
		fmt.Println(delErr)
		c.JSON(http.StatusNotFound, delErr.Error())
		return
	}
		c.JSON(http.StatusOK,result)
}
