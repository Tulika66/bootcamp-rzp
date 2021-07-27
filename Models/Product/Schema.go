package Product


type Product struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `gorm:"unique check:not null"`
	Price     uint64
	Quantity  uint64

}

func (b *Product) TableName() string {
	return "product"
}
