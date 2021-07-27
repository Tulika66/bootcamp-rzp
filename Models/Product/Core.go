package Product

import (
	"bootcamp/commerce/Config"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetAllProducts(products *[]Product) (err error) {

	if err:=Config.Db.Preload(clause.Associations).Find(products).Error;err!=nil{
		return err
	}

	return nil

}

func CreateProduct(product *Product)(err error){
    //Cache.CacheLocal.Set(product.Name,product.ID,cache.NoExpiration)
	if err :=Config.Db.Create(product).Error;err!=nil{
		return err
	}
	return nil
}

func GetProductById(product *Product, id string)(err error){
	if err :=Config.Db.Preload(clause.Associations).Where("id = ?",id).First(product).Error;err!=nil{
		return err
	}
	return nil
}


func UpdateProduct(product *Product, id string) (err error) {
	fmt.Println(product)
	//Config.DB.Save(student)


	return Config.Db.Transaction(func(tx *gorm.DB) error {
		if err =tx.Save(product).Error; err!=nil {
			return err
		}
		return nil
	})

}


func DeleteProduct(product *Product, id string) (err error) {


	//Cache.CacheLocal.Delete(product.Name)
	return Config.Db.Transaction(func(tx *gorm.DB) error {
		tx.Where("id = ?", id).Delete(product)

		return nil
	})

	//Config.Db.Where("id = ?", id).Delete(Product)
	//
	//return nil
}