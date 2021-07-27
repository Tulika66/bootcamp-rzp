package Controllers

import (
	"bootcamp/commerce/Models/Order"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)



func CreateOrder(c *gin.Context){
	var order Order.Order
	fmt.Println("Reached order's creates controller")
	c.BindJSON(&order)
	//fmt.Println("binded object order to json")
	err := Order.CreateOrder(&order)
	fmt.Println("error status check. err=",err)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

// GetOrderById gets the order having  id=id
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

func GetOrdersOfId(c *gin.Context){

	var orders []Order.Order
	c.BindJSON(orders)

	id := c.Params.ByName("id")
	err := Order.GetOrdersOfId(orders,id)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, orders)
	}


}


func OrderUpdate(c *gin.Context){

}

func GetAllProcessedOrders(c * gin.Context){

	var orders []Order.Order
	c.BindJSON(orders)


	err := Order.GetOrdersProcessed(orders)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, orders)
	}

}

