package Controllers

import (
	"bootcamp/commerce/Models/Product"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

)

func CreateProduct(c *gin.Context){
	var product Product.Product
	c.BindJSON(&product)

	err := Product.CreateProduct(&product)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func GetProductById(c *gin.Context){
	var product Product.Product
	c.BindJSON(&product)
	id := c.Params.ByName("id")
	err := Product.GetProductById(&product,id)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func GetAllProduct(c *gin.Context){
	var product []Product.Product
	c.BindJSON(&product)

	err := Product.GetAllProducts(&product)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}


func DeleteProduct(c *gin.Context){
	var product Product.Product
	c.BindJSON(&product)
	id := c.Params.ByName("id")
	err := Product.DeleteProduct(&product,id)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}
