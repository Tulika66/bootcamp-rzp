package Controllers

import (

	"bootcamp/commerce/Models/Customer"
	"encoding/json"
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


// features : 1. custom view all products
//            2. view his order history
//            3. Place an Order

func CustomerViewsProducts(c * gin.Context){
	GetAllProduct(c)
}

func ViewOrderHistory(c * gin.Context){
	id := c.Params.ByName("id")
	var customer Customer.Customer
	c.BindJSON(&customer)
	err := Customer.GetCustomerById(&customer,id)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(404,"Invalid Customer")
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		stringprint,_:=json.Marshal("\n Valid Customer. Fetching all orders!\n")
		c.JSON(http.StatusOK,string(stringprint) )
		GetOrdersOfId(c)
	}

}

func PlaceOrder(c *gin.Context){
	id := c.Params.ByName("id")
	var customer Customer.Customer
	c.BindJSON(&customer)
	err := Customer.GetCustomerById(&customer,id)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(404,"Invalid Customer")
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		stringprint,_:=json.Marshal("Valid Customer. Fetching all orders!")
		c.JSON(http.StatusOK,string(stringprint) )
		CreateOrder(c)
	}


}


