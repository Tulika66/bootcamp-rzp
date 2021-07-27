package Order

const (
	Executed="Executed"
	Pending="Processsing"
	Placed ="Placed"
)

type Order struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	ProductName string  `gorm:"check: not null"`
	CustomerId  uint    `json:"cust_id"`
	Price       uint64
	Quantity    uint64
	Status      string `gorm:"check: (status = 'Placed' OR status = 'Processing' OR status = 'Executed' ); not null"`

	//Products []Product.Product


}

func (b *Order) TableName() string {
	return "order"
}