package Controllers

import (
	"bootcamp/commerce/Models/Order"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)



func CreateOrder(c *gin.Context){
	var order Order.Order
	c.BindJSON(&order)

	err := Order.CreateOrder(&order)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, Order)
	}
}

func GetOrderById(c *gin.Context){
	var order Order.Order
	c.BindJSON(&order)
	id := c.Params.ByName("id")
	err := Order.GetOrderById(&order,id)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}
func GetAllOrder(c *gin.Context){
	var order []Order.Order
	c.BindJSON(&order)

	err := Order.GetAllOrders(&order)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}
func DeleteOrder(c *gin.Context){
	var order Order.Order
	c.BindJSON(&order)
	id := c.Params.ByName("id")
	err := Order.DeleteOrder(&order,id)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

