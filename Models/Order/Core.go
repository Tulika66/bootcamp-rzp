package Order

import (
	"bootcamp/commerce/Config"
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
	if err :=Config.Db.Create(order).Error;err!=nil{
		return err
	}
	return nil
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