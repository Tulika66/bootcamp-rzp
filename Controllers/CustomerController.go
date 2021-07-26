package Controllers

import (
	"bootcamp/commerce/Models/Customer"

	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCustomer(c *gin.Context){
	var customer Customer.Customer
	c.BindJSON(&customer)

	err := Customer.CreateCustomer(&customer)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, customer)
	}
}

func GetCustomerById(c *gin.Context){
	var customer Customer.Customer
	c.BindJSON(&customer)
	id := c.Params.ByName("id")
	err := Customer.GetCustomerById(&customer,id)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK,customer )
	}
}
func GetAllCustomers(c *gin.Context){
	var customers []Customer.Customer
	c.BindJSON(&customers)

	err := Customer.GetAllCustomer(&customers)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK,customers )
	}
}
func DeleteCustomer(c *gin.Context){
	var customer  Customer.Customer
	c.BindJSON(&customer)
	id := c.Params.ByName("id")
	err := Customer.DeleteCustomer(&customer,id)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK,customer )
	}
}

