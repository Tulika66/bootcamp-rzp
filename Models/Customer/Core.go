
package Customer

import (
	"bootcamp/commerce/Config"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)



func GetAllCustomer(customer *[]Customer)(err error){

	if err:=Config.Db.Preload(clause.Associations).Find(customer).Error;err!=nil{
		return err
	}

	return nil

}

func CreateCustomer(customer *Customer)(err error){
	if err :=Config.Db.Create(customer).Error;err!=nil{
		return err
	}
	return nil
}

func GetCustomerById(customer *Customer, id string)(err error){
	if err :=Config.Db.Preload(clause.Associations).Where("id = ?",id).First(customer).Error;err!=nil{
		return err
	}
	return nil
}


func UpdateCustomer(customer *Customer, id string) (err error) {
	fmt.Println(customer)


	return Config.Db.Transaction(func(tx *gorm.DB) error {
		if err =tx.Save(customer).Error; err!=nil {
			return err
		}
		return nil
	})

}


func DeleteCustomer(customer *Customer, id string) (err error) {

	return Config.Db.Transaction(func(tx *gorm.DB) error {
		tx.Where("id = ?", id).Delete(customer)

		return nil
	})

	//Config.Db.Where("id = ?", id).Delete(customer)
	//
	//return nil
}