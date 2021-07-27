package main

import (
	"bootcamp/commerce/Cache"
	"bootcamp/commerce/Config"
	"bootcamp/commerce/Models/Customer"
	"bootcamp/commerce/Models/Order"
	"bootcamp/commerce/Models/Product"
	"bootcamp/commerce/Routes"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


func DbInit(){

	var err1 error
	Config.Db, err1  = gorm.Open(mysql.Open(Config.DBUrl(Config.BuildConfig())), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err1 != nil   {
		fmt.Println("Error in db connection of Students.Status :", err1)
	}


	_ = Config.Db.AutoMigrate(&Customer.Customer{})
	_ =Config.Db.AutoMigrate(&Product.Product{})
	_ =Config.Db.AutoMigrate(&Order.Order{})

}

func main()  {

	DbInit()
    Cache.Init()


	r := Routes.SetupRouter()
	//running
	r.Run(":8080")

}


