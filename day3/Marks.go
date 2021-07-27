package Models


type Marks struct {
	//gorm:"primaryKey;autoIncrement:false"

	StudentId  uint    `json:"id" `
	Id         uint   ` gorm:"primaryKey"`
	Subject   string   `json:"subject"`
	Marks     uint64   `json:"marks"`

	//Student  Student   `gorm:"ForeignKey:Id;References:Id"`
}


func (b *Marks) TableName() string {
	return "marks"
}
