package Customer

import "bootcamp/commerce/Models/Order"



type Customer struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Phone     uint64
	Orders    []Order.Order

//order

}

	func (b *Customer) TableName() string {
	return "customer"
}


