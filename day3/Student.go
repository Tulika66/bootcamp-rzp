package Models

type Student struct {
	ID        uint   `json:"id" gorm:"primaryKey"  `
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	DOB       uint64    `json:"dob"`
	Address   string `json:"address"`
    Marks []Marks

}


func (b *Student) TableName() string {
	return "students"
}



