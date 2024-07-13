package entity

type User struct {
	ID          int    `gorm:"column:id;primaryKey;type:integer;autoIncrement;not null"`
	FullName    string `gorm:"column:full_name;type:varchar(100);not null"`
	Phone       string `gorm:"column:phone;type:varchar(15);unique;not null"`
	FamilyPhone string `gorm:"column:family_phone;type:varchar(15);not null"`
	Address     string `gorm:"column:address;type:text;not null"`
	Username    string `gorm:"column:username;type:varchar(50);unique;not null"`
	Password    string `gorm:"column:password;type:varchar(255);not null"`
}

func (User) TableName() string {
	return "users"
}
