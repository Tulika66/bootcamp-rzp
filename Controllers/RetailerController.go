package Controllers

import "github.com/gin-gonic/gin"

func ViewOrderOfProducts(c * gin.Context){
	GetAllProcessedOrders(c)
}
