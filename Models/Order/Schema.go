package Order

type Order struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	ProductName string
	CustomerId uint
	Price     uint64
	Quantity  uint64
	Status    string

	//Products []Product.Product
	//order

}

func (b *Order) TableName() string {
	return "order"
}
