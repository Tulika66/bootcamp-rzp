package Order

import (
	//"bootcamp/commerce/Cache"
	"bootcamp/commerce/Config"
	"bootcamp/commerce/Models/Product"
	//"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetAllOrders(orders *[]Order) (err error) {

	if err:=Config.Db.Preload(clause.Associations).Find(orders).Error;err!=nil{
		return err
	}

	return nil

}

func CreateOrder(order *Order)(err error){
	fmt.Println("Reached createorder()")
	var product Product.Product
	err2 := Config.Db.Where("Name = ?", order.ProductName).First(&product).Error
    fmt.Println("Reached in createorder()")
	if err2!=nil {
		fmt.Println("Invalid Product. Product does not exist!")
		return err2
	}
	if err :=Config.Db.Create(&order).Error;err!=nil{
				return err
			}

   return nil
    //_,found:=Cache.CacheLocal.Get(order.ProductName)
    //if found{
	//	if err :=Config.Db.Create(&order).Error;err!=nil{
	//		return err
	//	}
	//	return nil
	//}else{
	//	return errors.New("Invalid Product.Product does not exist!")
	//}


}

func GetOrderById(order *Order, id string)(err error){
	if err :=Config.Db.Preload(clause.Associations).Where("id = ?",id).First(order).Error;err!=nil{
		return err
	}
	return nil
}


func UpdateOrder(order *Order, id string) (err error) {
	fmt.Println(order)
	//Config.DB.Save(student)


	return Config.Db.Transaction(func(tx *gorm.DB) error {
		if err =tx.Save(order).Error; err!=nil {
			return err
		}
		return nil
	})

}


func DeleteOrder(order *Order, id string) (err error) {

	return Config.Db.Transaction(func(tx *gorm.DB) error {
		tx.Where("id = ?", id).Delete(order)

		return nil
	})

	//Config.Db.Where("id = ?", id).Delete(order)
	//
	//return nil
}

func GetOrdersOfId(orders []Order,id string)(err error){

	if err :=Config.Db.Preload(clause.Associations).Where("CustomerId = ?",id).Find(orders).Error;err!=nil{
		return err
	}
	return nil

}

func GetOrdersProcessed(orders []Order)(err error){

	if err :=Config.Db.Preload(clause.Associations).Where("status = ?",Executed).Find(orders).Error;err!=nil{
		return err
	}
	return nil

}