package Routes


import (
	"bootcamp/commerce/Controllers"
	"github.com/gin-gonic/gin"
)

func setupCustomerRoutes(grp1 *gin.RouterGroup){

	//customer routes
	grp1.GET("customers",Controllers.GetAllCustomers)
	grp1.POST("customer", Controllers.CreateCustomer)
	grp1.GET("customer/:id", Controllers.GetCustomerById)
	grp1.DELETE("customer/:id", Controllers.DeleteCustomer)
	//grp1.GET("customer/:id",Controllers.CustomerViewsProducts)
	grp1.POST("customer/:id", Controllers.PlaceOrder)
	grp1.GET("customers/:id/history",Controllers.ViewOrderHistory)
	grp1.GET("customers/viewProducts",Controllers.GetAllProduct)

}

func setupProductRoutes(grp1 *gin.RouterGroup) {

	//product routes
	grp1.GET("products",Controllers.GetAllProduct)
	grp1.POST("product", Controllers.CreateProduct)
	grp1.GET("product/:id", Controllers.GetProductById)
	grp1.DELETE("product/:id", Controllers.DeleteProduct)
}

func setupOrderRoutes(grp1 *gin.RouterGroup) {

	grp1.GET("/orders",Controllers.GetAllOrder)
	grp1.POST("/order",Controllers.CreateOrder)
	grp1.GET("/order/:id",Controllers.GetOrderById)
	grp1.DELETE("/orders",Controllers.DeleteOrder)


}


func setupRetailerRoutes(grp1 *gin.RouterGroup) {

	//retailer routes
	grp1.GET("retailer/viewTransactions",Controllers.ViewOrderOfProducts)


}


//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {

	r := gin.Default()
	grp1 := r.Group("/ecommerce")
	{
		setupCustomerRoutes(grp1)
        setupProductRoutes(grp1)
		setupOrderRoutes(grp1)
        setupRetailerRoutes(grp1)
	}
	return r
}

